package immkv

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestWriterAbortErrorDoesNotDeleteRetry(t *testing.T) {
	base := filepath.Join(t.TempDir(), "lib")
	w, err := Create(base, CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	// 注入 Close 错误；文件已经关闭，但 Abort 仍须删除本次临时文件并结束。
	if err := w.kf.Close(); err != nil {
		t.Fatal(err)
	}
	if err := w.Abort(); !errors.Is(err, os.ErrClosed) {
		t.Fatalf("Abort 应返回注入的关闭错误，得到 %v", err)
	}

	retry, err := Create(base, CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer retry.Abort()
	if err := retry.Add([]byte("k"), []byte("retry")); err != nil {
		t.Fatal(err)
	}
	if err := retry.Finish(); err != nil {
		t.Fatal(err)
	}
	if err := w.Abort(); err != nil {
		t.Fatal(err)
	}
	db, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatalf("旧 writer 的第二次 Abort 不应删除重建成品：%v", err)
	}
	defer db.Close()
	if value, found, err := db.Get([]byte("k")); err != nil || !found || string(value) != "retry" {
		t.Fatalf("重建结果 = %q, %v, %v", value, found, err)
	}
}

func TestWriterPreservesUnpublishedTarget(t *testing.T) {
	for _, suffix := range []string{".keys", ".vals"} {
		for _, finish := range []bool{false, true} {
			name := suffix + "/Abort"
			if finish {
				name = suffix + "/Finish"
			}
			t.Run(name, func(t *testing.T) {
				base := filepath.Join(t.TempDir(), "lib")
				w, err := Create(base, CreateOptions{})
				if err != nil {
					t.Fatal(err)
				}
				defer w.Abort()
				if err := w.Add([]byte("k"), []byte("value")); err != nil {
					t.Fatal(err)
				}
				// 在发布前放入同名目标：构建必须拒绝覆盖，也不能清理该路径。
				path := base + suffix
				if err := os.WriteFile(path, []byte("existing"), 0o600); err != nil {
					t.Fatal(err)
				}
				if finish {
					if err := w.Finish(); err == nil {
						t.Fatal("目标已存在，Finish 应报错")
					}
				} else if err := w.Abort(); err != nil {
					t.Fatal(err)
				}
				got, err := os.ReadFile(path)
				if err != nil || string(got) != "existing" {
					t.Fatalf("已有目标被改写或删除：%q, %v", got, err)
				}
				for _, ext := range []string{".keys", ".vals", ".keys.tmp", ".vals.tmp"} {
					if ext == suffix {
						continue
					}
					if _, err := os.Stat(base + ext); !os.IsNotExist(err) {
						t.Fatalf("本次失败产物 %s 应已清理：%v", ext, err)
					}
				}
			})
		}
	}
}

func TestWriteBuildFileRejectsExisting(t *testing.T) {
	for _, suffix := range []string{"", ".tmp"} {
		t.Run("existing"+suffix, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "header")
			if err := os.WriteFile(path+suffix, []byte("existing"), 0o600); err != nil {
				t.Fatal(err)
			}
			if err := writeBuildFile(path, []byte("replacement")); err == nil {
				t.Fatal("已有目标或临时文件时应报错")
			}
			got, err := os.ReadFile(path + suffix)
			if err != nil || string(got) != "existing" {
				t.Fatalf("已有内容被改写或删除：%q, %v", got, err)
			}
			other := path + ".tmp"
			if suffix == ".tmp" {
				other = path
			}
			if _, err := os.Stat(other); !os.IsNotExist(err) {
				t.Fatalf("失败的构建不应留下 %s：%v", other, err)
			}
		})
	}
}

func TestColumnarFailureCleansOnlyCreatedFiles(t *testing.T) {
	for _, suffix := range []string{".col", ".col.tmp", ".manifest", ".manifest.tmp"} {
		t.Run(suffix, func(t *testing.T) {
			base := filepath.Join(t.TempDir(), "lib")
			w, err := Create(base, CreateOptions{})
			if err != nil {
				t.Fatal(err)
			}
			defer w.Abort()
			if err := w.Add([]byte("k"), []byte("value")); err != nil {
				t.Fatal(err)
			}
			if err := w.Finish(); err != nil {
				t.Fatal(err)
			}
			path := base + suffix
			if err := os.WriteFile(path, []byte("existing"), 0o600); err != nil {
				t.Fatal(err)
			}
			id, err := newBuildID()
			if err != nil {
				t.Fatal(err)
			}
			if err := writeColumnar(base, simpleKVSchema(), id, SourceInfoOrder{}, 0, 0); err == nil {
				t.Fatal("列存目标已存在时应报错")
			}
			got, err := os.ReadFile(path)
			if err != nil || string(got) != "existing" {
				t.Fatalf("已有文件被改写或删除：%q, %v", got, err)
			}
			for _, ext := range []string{".col", ".col.tmp", ".manifest", ".manifest.tmp"} {
				if ext == suffix {
					continue
				}
				if _, err := os.Stat(base + ext); !os.IsNotExist(err) {
					t.Fatalf("失败后 %s 应已清理：%v", ext, err)
				}
			}
			db, err := OpenSimpleKV(base, 0)
			if err != nil {
				t.Fatalf("列存失败不应自行清理调用方的记录序文件：%v", err)
			}
			defer db.Close()
			if value, found, err := db.Get([]byte("k")); err != nil || !found || string(value) != "value" {
				t.Fatalf("记录序查询 = %q, %v, %v", value, found, err)
			}
		})
	}
}

func TestBuildFromTSVsCleansAfterColumnarFailure(t *testing.T) {
	dir := t.TempDir()
	input := filepath.Join(dir, "input.tsv")
	if err := os.WriteFile(input, []byte("k\tvalue\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	// 记录序可完成，但 manifest 不接受 Base="."，使失败发生在列存写完之后。
	base := dir + string(os.PathSeparator) + "."
	if err := BuildFromTSVs(base, []string{input}, BuildOptions{}); err == nil {
		t.Fatal("无效 manifest base 应使构建失败")
	}
	for _, suffix := range []string{".keys", ".vals", ".col", ".manifest", ".keys.tmp", ".vals.tmp", ".col.tmp", ".manifest.tmp"} {
		if _, err := os.Stat(base + suffix); !os.IsNotExist(err) {
			t.Fatalf("后续步骤失败时应清理 %s：%v", suffix, err)
		}
	}
	if data, err := os.ReadFile(input); err != nil || string(data) != "k\tvalue\n" {
		t.Fatalf("父目录中的输入文件不应受影响：%q, %v", data, err)
	}
}

func TestBuildFilesCleanupIsConsumedOnError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "not-empty")
	if err := os.Mkdir(path, 0o700); err != nil {
		t.Fatal(err)
	}
	child := filepath.Join(path, "child")
	if err := os.WriteFile(child, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	files := buildFiles{paths: []string{path}}
	if err := files.cleanup(); err == nil {
		t.Fatal("非空目录删除应失败")
	}
	if err := os.Remove(child); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("retry"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := files.cleanup(); err != nil {
		t.Fatal(err)
	}
	if got, err := os.ReadFile(path); err != nil || string(got) != "retry" {
		t.Fatalf("重复清理不得删除重试文件：%q, %v", got, err)
	}
}

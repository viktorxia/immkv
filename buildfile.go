package immkv

import (
	"errors"
	"fmt"
	"os"
)

// buildFiles 记录本次成功创建的路径。调用方独占构建目标；清理取走清单，
// 即使删除报错也不重试，避免影响后续重建。
type buildFiles struct {
	paths []string
}

func (b *buildFiles) create(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return nil, err
	}
	b.paths = append(b.paths, path)
	return f, nil
}

// publish 保持临时文件 + rename 的流程；目标已存在就报错。构建期间不得
// 由其他操作删除或替换同名文件。
func (b *buildFiles) publish(tmp, path string) error {
	for i, p := range b.paths {
		if p == tmp {
			if _, err := os.Lstat(path); err == nil {
				return fmt.Errorf("immkv: %s already exists; delete it and rebuild", path)
			} else if !os.IsNotExist(err) {
				return err
			}
			if err := os.Rename(tmp, path); err != nil {
				return err
			}
			b.paths[i] = path
			return nil
		}
	}
	return fmt.Errorf("immkv: 未创建临时文件 %s", tmp)
}

func (b *buildFiles) cleanup() error {
	paths := b.paths
	b.paths = nil
	var err error
	for i := len(paths) - 1; i >= 0; i-- {
		if rerr := os.Remove(paths[i]); rerr != nil && !os.IsNotExist(rerr) {
			err = errors.Join(err, fmt.Errorf("immkv: 清理 %s: %w", paths[i], rerr))
		}
	}
	return err
}

// writeBuildFile 用于构建产物，禁止覆盖；公开的 WriteManifest 仍可原子更新。
func writeBuildFile(path string, data []byte) (err error) {
	var files buildFiles
	f, err := files.create(path + ".tmp")
	if err != nil {
		return err
	}
	defer func() {
		if f != nil {
			err = errors.Join(err, f.Close())
		}
		if err != nil {
			err = errors.Join(err, files.cleanup())
		}
	}()
	if _, err := f.Write(data); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	err = f.Close()
	f = nil
	if err != nil {
		return err
	}
	return files.publish(path+".tmp", path)
}

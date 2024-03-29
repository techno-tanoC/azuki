package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Copier struct {
	dir string
	uid int
	gid int
}

func NewCopier(dir string, uid int, gid int) *Copier {
	return &Copier{
		dir: dir,
		uid: uid,
		gid: gid,
	}
}

func (c *Copier) Copy(r io.Reader, name string, ext string) error {
	f, err := createNewFile(c.dir, sanitize(name), sanitize(ext))
	if err != nil {
		return fmt.Errorf("failed to create file( name: %s, ext: %s ): %w", name, ext, err)
	}
	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		return fmt.Errorf("failed to copy file( name: %s, ext: %s ): %w", name, ext, err)
	}

	err = os.Chown(f.Name(), c.uid, c.gid)
	if err != nil {
		return fmt.Errorf("failed to change owner( name: %s, ext: %s ): %w", name, ext, err)
	}

	return nil
}

func createNewFile(dir, name, ext string) (*os.File, error) {
	count := 0

	for {
		var base string
		if count == 0 {
			base = fmt.Sprintf("%s.%s", name, ext)
		} else {
			base = fmt.Sprintf("%s(%d).%s", name, count, ext)
		}

		path := filepath.Join(dir, base)
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err == nil {
			return f, nil
		}
		if !os.IsExist(err) {
			return nil, fmt.Errorf("failed to create new file: %w", err)
		}

		count += 1
	}
}

func sanitize(path string) string {
	return strings.ReplaceAll(path, "/", "／")
}

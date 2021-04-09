package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/xerrors"
)

type LockCopy struct {
	mux sync.Mutex
	dir string
	uid int
	gid int
}

func NewLockCopy(dir string, uid int, gid int) *LockCopy {
	return &LockCopy{
		dir: dir,
		uid: uid,
		gid: gid,
	}
}

func (lc *LockCopy) Copy(r io.Reader, name string, ext string) error {
	lc.mux.Lock()
	defer lc.mux.Unlock()

	fresh := lc.buildFreshName(name, ext)

	f, err := os.Create(fresh)
	if err != nil {
		return xerrors.Errorf("failed to create file( name: %s, ext: %s ): %w", name, ext, err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		return xerrors.Errorf("failed to copy file( name: %s, ext: %s ): %w", name, ext, err)
	}

	err = os.Chown(f.Name(), lc.uid, lc.gid)
	if err != nil {
		return xerrors.Errorf("failed to change owner( name: %s, ext: %s ): %w", name, ext, err)
	}

	return nil
}

func (lc *LockCopy) buildFreshName(name string, ext string) string {
	var (
		candidate string
		count     int
		base      string
	)

	for {
		if count == 0 {
			base = fmt.Sprintf("%s.%s", name, ext)
		} else {
			base = fmt.Sprintf("%s(%d).%s", name, count, ext)
		}

		candidate = filepath.Join(lc.dir, base)
		if _, err := os.Stat(candidate); err != nil {
			break
		} else {
			count++
		}
	}

	return candidate
}

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
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

func (lc *LockCopy) Copy(r io.ReadSeeker, name string, ext string) error {
	lc.mux.Lock()
	defer lc.mux.Unlock()

	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	fresh := buildFreshName(lc.dir, name, ext)

	f, err := os.Create(fresh)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}

	err = os.Chown(f.Name(), lc.uid, lc.gid)
	return err
}

func buildFreshName(dir string, name string, ext string) string {
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

		candidate = filepath.Join(dir, base)
		if _, err := os.Stat(candidate); err != nil {
			break
		} else {
			count++
		}
	}

	return candidate
}

package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCopyContents(t *testing.T) {
	dir, err := ioutil.TempDir("", "azuki_*")
	if err != nil {
		t.Fatalf("TestCopyContents error creating temporary dir: %v", err)
	}
	defer os.RemoveAll(dir)

	lc := NewLockCopy(dir, 1000, 1000)
	data := "ハローワールド"

	buf := bytes.NewBuffer([]byte(data))
	lc.Copy(buf, "azuki", "txt")

	// Read copied file
	contents, err := ioutil.ReadFile(filepath.Join(dir, "azuki.txt"))
	if err != nil {
		t.Fatalf("TestCopyContents error reading file: %v", err)
	}

	diff := cmp.Diff(string(contents), data)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
}

func TestCopyName(t *testing.T) {
	dir, err := ioutil.TempDir("", "azuki_*")
	if err != nil {
		t.Fatalf("TestCopyName: %v", err)
	}
	defer os.RemoveAll(dir)

	lc := NewLockCopy(dir, 1000, 1000)

	buf := bytes.NewBuffer([]byte(""))
	lc.Copy(buf, "azuki", "txt")
	lc.Copy(buf, "azuki", "txt")
	lc.Copy(buf, "azuki", "txt")

	// List copied files
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatalf("TestCopyName: %v", err)
	}

	var names []string
	for _, name := range fileInfos {
		names = append(names, name.Name())
	}

	expectedNames := []string{"azuki.txt", "azuki(1).txt", "azuki(2).txt"}
	opt := cmpopts.SortSlices(func(i, j string) bool {
		return i < j
	})
	diff := cmp.Diff(names, expectedNames, opt)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
}

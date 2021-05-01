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
	dir, err := os.MkdirTemp("", "azuki_*")
	if err != nil {
		t.Fatalf("TestCopyContents create temporary dir error: %v", err)
	}
	defer os.RemoveAll(dir)

	c := NewCopier(dir, 1000, 1000)
	data := "ハローワールド"

	buf := bytes.NewBuffer([]byte(data))
	c.Copy(buf, "azuki", "txt")

	// Read copied file
	contents, err := ioutil.ReadFile(filepath.Join(dir, "azuki.txt"))
	if err != nil {
		t.Fatalf("TestCopyContents read file error: %v", err)
	}

	diff := cmp.Diff(string(contents), data)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
}

func TestCopyName(t *testing.T) {
	dir, err := os.MkdirTemp("", "azuki_*")
	if err != nil {
		t.Fatalf("TestCopyName create temporary dir error: %v", err)
	}
	defer os.RemoveAll(dir)

	c := NewCopier(dir, 1000, 1000)

	buf := bytes.NewBuffer([]byte(""))
	c.Copy(buf, "azuki", "txt")
	c.Copy(buf, "azuki", "txt")
	c.Copy(buf, "azuki", "txt")

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

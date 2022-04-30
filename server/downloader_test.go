package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type TestClient struct {
	sourcePath string
}

func (t *TestClient) Get(url string) (ResponseLike, error) {
	path := "go.sum"
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	res := &TestResponse{f: f, len: info.Size()}
	return res, nil
}

type TestResponse struct {
	f   *os.File
	len int64
}

func (t *TestResponse) Read(p []byte) (int, error) {
	return t.f.Read(p)
}

func (t *TestResponse) Close() error {
	return t.f.Close()
}

func (t *TestResponse) ContentLength() int64 {
	return t.len
}

func TestDownload(t *testing.T) {
	src := "go.sum"

	temp, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatalf("TestDownload create temporary dir error: %v", err)
	}
	defer os.RemoveAll(temp)

	user, err := user.Current()
	if err != nil {
		t.Fatalf("TestDownload current user error: %v", err)
	}
	uid, err := strconv.Atoi(user.Uid)
	if err != nil {
		t.Fatalf("TestDownload user id error: %v", err)
	}
	gid, err := strconv.Atoi(user.Gid)
	if err != nil {
		t.Fatalf("TestDownload current user error: %v", err)
	}

	client := &TestClient{src}
	downloader := NewDownloader(client, temp, uid, gid)
	err = downloader.Download("", "test", "txt")
	if err != nil {
		t.Fatalf("TestDownload download error: %v", err)
	}

	srcData, err := ioutil.ReadFile(src)
	if err != nil {
		t.Fatalf("TestDownload src read file error: %v", err)
	}

	dest := filepath.Join(temp, "test.txt")
	destData, err := ioutil.ReadFile(dest)
	if err != nil {
		t.Fatalf("TestDownload dest read file error: %v", err)
	}

	diff := cmp.Diff(string(srcData), string(destData))
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
}

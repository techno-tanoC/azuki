package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	temp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("TestDownload create temporary dir error: %v", err)
	}

	client := &TestClient{src}
	downloader := NewDownloader(client, temp, 1000, 1000)
	err = downloader.Download("", "test", "txt")
	if err != nil {
		t.Fatalf("TestDownload download error: %v", err)
	}

	srcData, err := ioutil.ReadFile(src)

	dest := filepath.Join(temp, "test.txt")
	destData, err := ioutil.ReadFile(dest)

	diff := cmp.Diff(string(srcData), string(destData))
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
}

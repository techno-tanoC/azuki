package main

import (
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type Downloader struct {
	client   *http.Client
	lockCopy *LockCopy
	table    *Table
}

func NewDownloader(dir string, uid int, gid int) *Downloader {
	dialer := &net.Dialer{
		Timeout: 3 * time.Second,
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial:                  dialer.Dial,
			TLSHandshakeTimeout:   3 * time.Second,
			ResponseHeaderTimeout: 3 * time.Second,
		},
	}
	lockCopy := NewLockCopy(dir, uid, gid)
	table := NewTable()

	return &Downloader{
		client:   client,
		lockCopy: lockCopy,
		table:    table,
	}
}

func (d *Downloader) Download(url string, name string, ext string) error {
	// Make pg
	pg := NewProgress(name)

	// Register pg( defer unregister )
	v4, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uuid := v4.String()

	d.table.Add(uuid, pg)
	defer d.table.Delete(uuid)

	// Make temp file( defer delete)
	temp, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	defer os.Remove(temp.Name())

	// Open http
	res, err := d.client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Set ContentLength
	pg.SetTotal(res.ContentLength)

	// Make MultiWriter
	tempPg := io.MultiWriter(temp, pg)

	// Copy data from response to temp file
	_, err = io.Copy(tempPg, res.Body)
	if err != nil {
		return err
	}

	// Copy data from temp file to file
	err = d.lockCopy.Copy(temp, name, ext)
	if err != nil {
		return err
	}

	return nil
}

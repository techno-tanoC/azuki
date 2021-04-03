package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type Downloader struct {
	client   ClientLike
	lockCopy *LockCopy
	table    *Table
}

func NewDownloader(client ClientLike, dir string, uid int, gid int) *Downloader {
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

	// Make uuid
	v4, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uuid := v4.String()

	// Register pg( defer unregister )
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
	defer res.Close()

	// Set ContentLength
	pg.SetTotal(res.ContentLength())

	// Make MultiWriter
	tempPg := io.MultiWriter(temp, pg)

	// Copy data from response to temp file
	_, err = io.Copy(tempPg, res)
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

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

type Downloader struct {
	client ClientLike
	copier *Copier
	table  *Table
}

func NewDownloader(client ClientLike, dir string, uid int, gid int) *Downloader {
	copier := NewCopier(dir, uid, gid)
	table := NewTable()

	return &Downloader{
		client: client,
		copier: copier,
		table:  table,
	}
}

func (d *Downloader) Download(url string, name string, ext string) error {
	// Make pg
	pg := NewProgress(name)

	// Make uuid
	v4, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %w", err)
	}
	uuid := v4.String()

	// Register pg( defer unregister )
	d.table.Add(uuid, pg)
	defer d.table.Delete(uuid)

	// Make temp file( defer delete )
	temp, err := os.CreateTemp("", "")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer os.Remove(temp.Name())

	// Open http
	res, err := d.client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get response: %w", err)
	}
	defer res.Close()

	// Set ContentLength
	pg.SetTotal(res.ContentLength())

	// Make MultiWriter
	tempPg := io.MultiWriter(temp, pg)

	// Copy data from response to temp file
	_, err = io.Copy(tempPg, res)
	if err != nil {
		return fmt.Errorf("failed to copy data from response to temporary file: %w", err)
	}

	// Reset tempfile offset
	_, err = temp.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to seek temporary file: %w", err)
	}

	if pg.Canceled {
		return fmt.Errorf("cancelled downloading %s", name)
	}

	// Copy data from temp file to file unless canceled
	err = d.copier.Copy(temp, name, ext)
	if err != nil {
		return fmt.Errorf("failed to copy with lock: %w", err)
	}

	return nil
}

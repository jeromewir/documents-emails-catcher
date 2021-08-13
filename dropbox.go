package main

import (
	"fmt"
	"io"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/jeromewir/invoices-fwder/config"
)

type DropboxClient struct {
	files files.Client
}

func NewDropboxClient() *DropboxClient {
	token := config.GetDropboxToken()

	c := dropbox.Config{
		Token:    token,
		LogLevel: dropbox.LogOff, // if needed, set the desired logging level. Default is off
	}

	f := files.New(c)

	return &DropboxClient{
		files: f,
	}
}

func (d *DropboxClient) getDirectoryPathDatedFilename(filename string) string {
	year, month, _ := time.Now().Date()

	directoryPath := fmt.Sprintf("/%s/Documents/%d/%s/%s", config.GetDropboxDestinationDirectory(), year, month, filename)

	return directoryPath
}

func (d *DropboxClient) uploadFile(directoryPath string, fd io.Reader) error {
	config := files.NewCommitInfo(directoryPath)

	config.Autorename = true

	fmt.Println("Uploading", directoryPath)

	_, err := d.files.Upload(config, fd)

	if err != nil {
		return err
	}

	return err
}

func (d *DropboxClient) UploadDatedFile(fd io.Reader, filename string) error {
	directoryPath := d.getDirectoryPathDatedFilename(filename)

	return d.uploadFile(directoryPath, fd)
}

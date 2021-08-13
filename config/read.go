package config

import (
	"errors"
	"os"
)

var (
	dropboxToken                string
	dropboxDestinationDirectory string
)

func ReadFromEnvironment() error {
	dropboxToken = os.Getenv("DROPBOX_TOKEN")
	dropboxDestinationDirectory = os.Getenv("DROPBOX_DESTINATION_DIRECTORY")

	if dropboxToken == "" {
		return errors.New("expected DROPBOX_TOKEN to be defined")
	}

	if dropboxDestinationDirectory == "" {
		return errors.New("expected DROPBOX_DESTINATION_DIRECTORY to be defined")
	}

	return nil
}

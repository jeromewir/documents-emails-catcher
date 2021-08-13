package config

// Returns the dropbox token read from environment
func GetDropboxToken() string {
	return dropboxToken
}

func GetDropboxDestinationDirectory() string {
	return dropboxDestinationDirectory
}

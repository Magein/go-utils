package oss

import (
	"strings"
)

// FileExtension Get file extension
func FileExtension(filename string) string {
	index := strings.LastIndex(filename, ".")
	var extension string
	if index == -1 {
		extension = ""
	} else {
		extension = filename[index+1:]
	}
	return strings.ToLower(extension)
}

// FileInfo Find information about the file
type FileInfo struct {
	Path       string
	Name       string
	Extension  string
	Size       int
	SizeFormat string
	Mode       string
	ModTime    string
}

package tools

import "os"

// IsDir Determine whether it is a directory
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile Judge whether the given path is a file
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if info.IsDir() {
		return false
	}
	return true
}

// IsInt Judge whether the given parameter is an int
func IsInt(parameter interface{}) bool {
	_, ok := parameter.(int)
	return ok
}

// IsString Judge whether the given parameter is an string
func IsString(parameter interface{}) bool {
	_, ok := parameter.(string)
	return ok
}

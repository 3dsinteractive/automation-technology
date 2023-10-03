package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

// File is the struct to work with file
type File struct {
}

// NewFile return file service
func NewFile() *File {
	return &File{}
}

// TempDir create temp directory with prefix
func (f *File) TempDir(prefix string) (string, error) {
	return ioutil.TempDir("", prefix)
}

// Write write string to file
func (f *File) Write(file string, data string) error {
	err := ioutil.WriteFile(file, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

// WriteTemp create temp file and write string to file
// It return path for the created file
func (f *File) WriteTemp(data string, prefix string) ( /*filepath*/ string, error) {
	file, err := os.CreateTemp(os.TempDir(), prefix)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}

// WriteTempB create temp file and write byte array to file
// It return path for the created file
func (f *File) WriteTempB(data []byte, prefix string) ( /*filepath*/ string, error) {
	file, err := os.CreateTemp(os.TempDir(), prefix)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}

// Read read content of file into string
func (f *File) Read(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Exists check if file is exists
func (f *File) Exists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Delete the file
func (f *File) Delete(file string) error {
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

// Download will download file from URL to destination folder and return the downloaded file
// url is a URL of file you want to download
// destinationDir is a destination folder
// filename is a destination file name
func (f *File) Download(url string, destinationDir string, filename string) ( /*filepath*/ string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//get the final url if original url make redirection
	finalURL := resp.Request.URL.String()

	saveFileName := filename
	if filename == "" {
		_, fileName := path.Split(finalURL)
		saveFileName = fmt.Sprintf("%s", fileName)
	}

	filePath := path.Join(destinationDir, saveFileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	io.Copy(file, resp.Body)

	return file.Name(), nil
}

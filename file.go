package filebucket

import (
	"bytes"
	"context"
	"os"
)

type File struct {
	Filename string
	Data     bytes.Buffer
}

type FileBucket struct {
	Filepath string
}

func NewFilebucket(bucketname string) (*FileBucket, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := homeDir + string(os.PathListSeparator) + bucketname
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return nil, err
		}
	}

	return &FileBucket{
		Filepath: path,
	}, nil
}

// Write writes file to a bucket
func (f *FileBucket) Write(ctx context.Context, file File) error {
	data := file.Data.Bytes()

	err := os.WriteFile(f.Filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Read returns array of byte of a file
func (f *FileBucket) Read(ctx context.Context, filename string) ([]byte, error) {
	file := f.Filepath + string(os.PathListSeparator) + filename
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

package store

import (
	"encoding/json"
)

type FileMock struct {
	Data []byte
	Err  error
}

func (fs *FileMock) Write(data interface{}) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if fs.Err != nil {
		return nil
	}
	fs.Data = fileData
	return nil
}

func (fs *FileMock) Read(data interface{}) error {
	if fs.Err != nil {
		return fs.Err
	}
	return json.Unmarshal(fs.Data, data)
}

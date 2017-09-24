package client

import (
	"os"
)

type ClientInterface interface {
	UploadFile(file *os.File) (map[string]string, error)
	DeleteFile(filename string) (map[string]string, error)
	OverwriteFile(file *os.File, filename string) (map[string]string, error)
}

type Client struct {
	Request GOrequest
}

// Use this method to upload file to GOstorage service.
func (client Client) UploadFile(file *os.File) (map[string]string, error) {
	const UPLOAD_FILE_URL = "/api/v1/storage/file/upload"

	return client.Request.UploadFile(UPLOAD_FILE_URL, file)
}


// Use this method to Delete file from GOstorage service.
func (client Client) DeleteFile(filename string) (map[string]string, error) {
	const DELETE_FILE_URL = "/api/v1/storage/file/upload"

	return client.Request.Delete(DELETE_FILE_URL, map[string]string{"FileName": filename})
}


// Use this method to Delete file from GOstorage service.
func (client Client) OverwriteFile(file *os.File, filename string) (map[string]string, error) {
	data, err := client.DeleteFile(filename)

    if err != nil {
    	return data, err
	}

    data, err = client.UploadFile(file)

	if err != nil {
		return data, err
	}

	return map[string]string{"result": filename + " overwritten."}, nil
}
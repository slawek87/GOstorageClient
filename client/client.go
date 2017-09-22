package client

import (
	"os"
	"mime/multipart"
)

type ClientInterface interface {
	UploadFile(file os.File) (map[string]string, error)
	DeleteFile(filename string) map[string]string
	OverwriteFile(file os.File, filename string) (map[string]string, error)
}

type Client struct {
	request GOrequest
}

// Use this method to upload file to GOstorage service.
func (client *Client) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (map[string]string, error) {
	const UPLOAD_FILE_URL = "/api/v1/storage/file/upload"
	var data map[string]string
	var body map[string]multipart.File

	file, err := fileHeader.Open()

	if err != nil {
		return data, err
	}

	body["upload"] = &file

	return client.request.sendFile(UPLOAD_FILE_URL, &body)
}


// Use this method to delete file from GOstorage service.
func (client *Client) DeleteFile(filename string) (map[string]string, error) {
	const DELETE_FILE_URL = "/api/v1/storage/file/upload"
	var formData map[string]string

	formData["FileName"] = filename

	return client.request.delete(DELETE_FILE_URL, formData)
}


// Use this method to delete file from GOstorage service.
func (client *Client) OverwriteFile(file multipart.File, fileHeader *multipart.FileHeader, filename string) (map[string]string, error) {
	data, err := client.DeleteFile(filename)

    if err != nil {
    	return data, err
	}

    data, err = client.UploadFile(file, fileHeader)

	if err != nil {
		return data, err
	}

	return map[string]string{"result": filename + " overwritten."}, nil
}
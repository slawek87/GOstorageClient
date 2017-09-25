package client

import (
	"gopkg.in/resty.v0"
	"github.com/slawek87/GOstorageClient/conf"
	"errors"
	"io"
	"bytes"
	"mime/multipart"
	"os"
	"strings"
)

type GOrequest struct{}

// method returns resty instance with already set BasicAuth token.
func (goRequest *GOrequest) resty() *resty.Request {
	request := resty.R()
	request.SetBasicAuth(conf.Settings.GetSettings("USERNAME"), conf.Settings.GetSettings("PASSWORD"))

	return request
}

// method handle response. If status code is equal or grater then 300 returns error msg.
func (GOrequest *GOrequest) handleResponse(data map[string]string, response *resty.Response) (map[string]string, error) {
	if response.StatusCode() >= 300 {
		return data, errors.New(response.Status())
	}

	return data, nil
}

// method deletes file in storage. It must be Post method because endpoints in REST don't support Delete request with params.
func (goRequest *GOrequest) Delete(url string, filename string) (map[string]string, error) {
	return goRequest.Post(url, map[string]string{"FileName": filename})
}

func (goRequest *GOrequest) Post(url string, formData map[string]string) (map[string]string, error) {
	var data map[string]string

	url = goRequest.GetURL(url)

	response, _ := goRequest.resty().
		SetFormData(formData).
		SetResult(&data).
		Post(url)

	return goRequest.handleResponse(data, response)
}

// method upload new file to storage.
func (goRequest *GOrequest) UploadFile(url string, file *os.File) (map[string]string, error) {
	var data map[string]string
	var body bytes.Buffer

	url = goRequest.GetURL(url)

	writeBody := multipart.NewWriter(&body)

	filename := file.Name()

	strings.Contains(filename, "/")
	{
		tmp := strings.Split(filename, "/")
		filename = tmp[len(tmp)-1]
	}

	formFile, _ := writeBody.CreateFormFile("upload", filename)
	io.Copy(formFile, file)

	writeBody.Close()

	response, _ := goRequest.resty().
		SetHeader("Content-Type", writeBody.FormDataContentType()).
		SetBody(&body).
		SetContentLength(true).
		SetResult(&data).
		Post(url)

	return goRequest.handleResponse(data, response)
}

// method returns correct url to storage service.
func (goRequest *GOrequest) GetURL(url string) string {
	return conf.Settings.GetSettings("PROTOCOL") + "://" +
		conf.Settings.GetSettings("HOST") + ":" +
		conf.Settings.GetSettings("PORT") +
		url
}

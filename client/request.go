package client

import (
	"gopkg.in/resty.v0"
	"github.com/slawek87/GOstorageClient/conf"
	"mime/multipart"
	"errors"
)

type GOrequest struct{}

func (goRequest *GOrequest) resty() *resty.Request  {
	request := resty.R()
	request.SetHeader("Content-Type", "application/json")
	request.SetBasicAuth(conf.Settings.GetSettings("USERNAME"), conf.Settings.GetSettings("PASSWORD"))

	return request
}

func (goRequest *GOrequest) delete(url string, formData map[string]string) (map[string]string, error) {
	var data map[string]string
	response, _ := goRequest.resty().SetFormData(formData).SetResult(&data).Delete(goRequest.GetURL(url))

	if response.StatusCode() >= 300 {
		errors.New(response.Status())
	}

	return data, nil
}

func (goRequest *GOrequest) post(url string, formData map[string]string) (map[string]string, error) {
	var data map[string]string
	response, _ := goRequest.resty().SetFormData(formData).SetResult(&data).Post(goRequest.GetURL(url))

	if response.StatusCode() >= 300 {
		errors.New(response.Status())
	}

	return data, nil
}

func (goRequest *GOrequest) sendFile(url string, body *map[string]multipart.File) (map[string]string, error) {
	var data map[string]string

	response, _ := goRequest.resty().SetBody(body).
		SetContentLength(true).
		SetResult(&data).
		Post(url)

	if response.StatusCode() >= 300 {
		errors.New(response.Status())
	}

	return data, nil
}

func (goRequest *GOrequest) GetURL(url string) string {
	return conf.Settings.GetSettings("PROTOCOL") + "://" +
		conf.Settings.GetSettings("HOST") + ":" +
		conf.Settings.GetSettings("PORT") +
		url
}

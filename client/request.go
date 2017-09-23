package client

import (
	"gopkg.in/resty.v0"
	"github.com/slawek87/GOstorageClient/conf"
	"errors"
	"os"
	"fmt"
)

type GOrequest struct{}

func (goRequest *GOrequest) resty() *resty.Request  {
	request := resty.R()
	request.SetHeader("Content-Type", "application/json")
	fmt.Println(conf.Settings.GetSettings("USERNAME"), conf.Settings.GetSettings("PASSWORD"))

	request.SetBasicAuth(conf.Settings.GetSettings("USERNAME"), conf.Settings.GetSettings("PASSWORD"))

	return request
}

func (GOrequest *GOrequest) handleResponse(data map[string]string, response *resty.Response) (map[string]string, error) {
	if response.StatusCode() >= 300 {
		return data, errors.New(response.Status())
	}

	return data, nil
}

func (goRequest *GOrequest) Delete(url string, formData map[string]string) (map[string]string, error) {
	var data map[string]string

	url = goRequest.GetURL(url)

	response, _ := goRequest.resty().SetFormData(formData).SetResult(&data).Delete(url)

	return goRequest.handleResponse(data, response)
}

func (goRequest *GOrequest) Post(url string, formData map[string]string) (map[string]string, error) {
	var data map[string]string

	url = goRequest.GetURL(url)

	response, _ := goRequest.resty().SetFormData(formData).SetResult(&data).Post(url)

	return goRequest.handleResponse(data, response)
}

func (goRequest *GOrequest) UploadFile(url string, body map[string]*os.File) (map[string]string, error) {
	var data map[string]string

	url = goRequest.GetURL(url)

	response, _ := goRequest.resty().SetBody(body).
		SetContentLength(true).
		SetResult(&data).
		Post(url)

	return goRequest.handleResponse(data, response)
}

func (goRequest *GOrequest) GetURL(url string) string {
	return conf.Settings.GetSettings("PROTOCOL") + "://" +
		conf.Settings.GetSettings("HOST") + ":" +
		conf.Settings.GetSettings("PORT") +
		url
}

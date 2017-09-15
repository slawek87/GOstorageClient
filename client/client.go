package client

import (
	"os"
	"reflect"
	"github.com/slawek87/GOstorageClient/conf"
	"gopkg.in/resty.v0"
)

type ClientInterface interface {
	UploadFile(file os.File) error
	DeleteFile(filename string) error
	OverwriteFile(file os.File, filename string) error
}

type GOrequest struct {}

func (goRequest *GOrequest) resty() *resty.Request {
	request := resty.R()
	request.SetHeader("Content-Type", "application/json")
	request.SetBasicAuth(conf.Settings.GetSettings("USERNAME"), conf.Settings.GetSettings("PASSWORD"))

	return request
}

func (goRequest *GOrequest) mapStructure(data map[string]interface{}, result interface{}) interface{} {
	elements := reflect.ValueOf(result).Elem()

	for key, value := range data {
		getValue := elements.FieldByName(key)

		if getValue.IsValid() {
			getValue.Set(reflect.ValueOf(value))
		}
	}

	return result
}

func (goRequest *GOrequest) post(url string, formData map[string]string, result interface{}) (map[string]interface{}, interface{}) {
	var data map[string]interface{}

	goRequest.resty().SetFormData(formData).SetResult(&data).Post(goRequest.GetURL(url))
	goRequest.mapStructure(data, result)

	return data, result
}

func (goRequest *GOrequest) GetURL(url string) string {
	return conf.Settings.GetSettings("PROTOCOL") + "://" + conf.Settings.GetSettings("HOST") + ":" + conf.Settings.GetSettings("PORT") + url
}


type Client struct {
	request GOrequest
}

func (client *Client) UploadFile(file os.File) error {

}
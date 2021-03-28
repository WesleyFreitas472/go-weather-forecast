package services

import (
	"fmt"
	"net/http"
)

type HttpServiceInterface interface {
	Get(url string, params map[string]string) (*http.Response, error)
	Post(url string, params map[string]string) (*http.Response, error)
}

type HttpServiceImpl struct {
}

func makeURL(url string, params map[string]string) string {
	url = url + "?&"
	for key, value := range params {
		url = fmt.Sprintf("%s%s=%s&", url, key, value)
	}
	return url
}

func (httpSvc HttpServiceImpl) Get(url string, params map[string]string) (*http.Response, error) {
	url = makeURL(url, params)
	resp, err := http.Get(url)
	return resp, err
}

func (httpSvc HttpServiceImpl) Post(url string, params map[string]string) (*http.Response, error) {
	url = makeURL(url, params)
	resp, err := http.Post(url, "", nil)
	return resp, err
}

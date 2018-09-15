package infrastructure

import (
	"gitcli/infrastructure/errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// GetResponseNetHttp ...
func GetResponseNetHttp(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, &errors.ErrorCmdRequest
	}

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, &errors.ErrorCmdResponse

	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

// PostResponseNetHttp ...
func PostResponseNetHttp(url string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, &errors.ErrorCmdRequest
	}
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, &errors.ErrorCmdResponse
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)

}

// GetResponseGoRequests ...
func GetResponseGoRequests(url string) ([]byte, error) {
	request := gorequest.New()
	response, _, err := request.Get(url).End()
	if err != nil {
		return nil, &errors.ErrorCmdResponse
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

// PostResponseGoRequests ...
func PostResponseGoRequests(url string, body interface{}) ([]byte, error) {
	request := gorequest.New()
	response, body, err := request.Post(url).
		Send(body).End()
	if err != nil {
		return nil, &errors.ErrorCmdResponse
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

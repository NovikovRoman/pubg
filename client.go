package pubg

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	apiPointPath   = "https://api.pubg.com"
	layoutDateTime = "2006-01-02T15:04:05.9Z"
)

type Client struct {
	httpClient http.Client
	apikey     string
}

func NewClient(apikey string) Client {
	return Client{
		httpClient: http.Client{},
		apikey:     apikey,
	}
}

func (c Client) requestGET(platform Platform, u string) (body []byte, err error) {
	if platform.IsEmpty() {
		u = fmt.Sprintf("%s%s", apiPointPath, u)
	} else {
		u = fmt.Sprintf("%s/shards/%s%s", apiPointPath, platform, u)
	}

	req, _ := http.NewRequest("GET", u, nil)
	req.Header.Set("Authorization", "Bearer "+c.apikey)
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Add("Accept-Encoding", "gzip")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}

	body, err = c.getBytes(resp)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = c.handlerErrorResponse(resp.StatusCode, body)
	}
	return
}

func (c Client) getBytes(resp *http.Response) (body []byte, err error) {
	defer func() {
		if derr := resp.Body.Close(); derr != nil {
			err = derr
		}
	}()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer func() {
			if derr := reader.Close(); derr != nil {
				err = derr
			}
		}()

	default:
		reader = resp.Body
	}

	body, err = ioutil.ReadAll(reader)
	return
}

func (c Client) handlerErrorResponse(statusCode int, body []byte) (err error) {
	var data struct {
		Errors []struct {
			Title  string `json:"title"`
			Detail string `json:"detail"`
		} `json:"errors"`
	}

	if len(body) == 0 {
		err = fmt.Errorf("Empty response body. StatusCode: %d ", statusCode)
		return
	}

	if err = json.Unmarshal(body, &data); err != nil {
		err = fmt.Errorf("%s. StatusCode: %d. Body: %s ", err.Error(), statusCode, string(body))
		return
	}

	if len(data.Errors) == 0 {
		err = fmt.Errorf("StatusCode: %d. Body: %s ", statusCode, string(body))
		return
	}

	switch statusCode {
	case 400:
		return ErrBadRequest{
			title:  data.Errors[0].Title,
			detail: data.Errors[0].Detail,
		}

	case 401:
		return ErrUnauthorized{
			title:  data.Errors[0].Title,
			detail: data.Errors[0].Detail,
		}

	case 404:
		return ErrNotFound{
			title:  data.Errors[0].Title,
			detail: data.Errors[0].Detail,
		}

	case 415:
		return ErrUnsupportedMediaType{
			title:  data.Errors[0].Title,
			detail: data.Errors[0].Detail,
		}

	case 429:
		return ErrTooManyRequest{}
	}

	err = fmt.Errorf("StatusCode: %d. Body: %s ", statusCode, string(body))
	return
}

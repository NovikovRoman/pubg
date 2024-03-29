package pubg

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	apiPointPath   = "https://api.pubg.com"
	layoutDateTime = "2006-01-02T15:04:05.9Z"
)

// Client structure.
type Client struct {
	httpClient *http.Client
	apikey     string
}

// NewClient create new PUBG Client structure.
// This structure will be used to acquire make API calls.
func NewClient(apikey string, transport *http.Transport) *Client {
	if transport == nil {
		transport = &http.Transport{}
	}

	return &Client{
		httpClient: &http.Client{
			Transport: transport,
		},
		apikey: apikey,
	}
}

func (c Client) requestGET(platform Platform, u string) (body []byte, statusCode int, err error) {
	if !platform.IsValid() {
		err = errors.New("Unknown platform. ")
		return
	}

	if platform.IsEmpty() {
		u = fmt.Sprintf("%s%s", apiPointPath, u)
	} else {
		u = fmt.Sprintf("%s/shards/%s%s", apiPointPath, platform, u)
	}

	req, _ := http.NewRequest("GET", u, nil)
	req.Header.Set("Authorization", "Bearer "+c.apikey)
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Add("Accept-Encoding", "gzip")
	var resp *http.Response
	if resp, err = c.httpClient.Do(req); err != nil {
		return
	}

	statusCode = resp.StatusCode
	if body, err = c.getBytes(resp); err != nil {
		return
	}

	if statusCode != http.StatusOK {
		err = c.handlerErrorResponse(statusCode, body)
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

	body, err = io.ReadAll(reader)
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
	case http.StatusBadRequest:
		return &ErrBadRequest{
			title:  data.Errors[0].Title,
			detail: data.Errors[0].Detail,
		}

	case http.StatusUnauthorized:
		return &ErrUnauthorized{
			title:  data.Errors[0].Title,
			detail: data.Errors[0].Detail,
		}

	case http.StatusNotFound:
		return &ErrNotFound{
			title:  data.Errors[0].Title,
			detail: data.Errors[0].Detail,
		}

	case http.StatusUnsupportedMediaType:
		return &ErrUnsupportedMediaType{
			title:  data.Errors[0].Title,
			detail: data.Errors[0].Detail,
		}

	case http.StatusTooManyRequests:
		return &ErrTooManyRequest{}
	}

	err = fmt.Errorf("StatusCode: %d. Body: %s ", statusCode, string(body))
	return
}

package pubg

import "net/http"

func (c Client) Status() bool {
	resp, _ := http.Get(apiPointPath + "/status")
	return resp.StatusCode == 200
}

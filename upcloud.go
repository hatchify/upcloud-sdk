package upcloud

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	// Hostname of API
	Hostname = "https://api.upcloud.com"
	// APIVersion Current API version
	APIVersion = "1.3"
)

// New will return a new instance of the UpCloud API SDK
func New(username, password string) (up *UpCloud, err error) {
	var u UpCloud
	if u.host, err = url.Parse(Hostname); err != nil {
		return
	}

	u.username = username
	u.password = password
	up = &u
	return
}

// UpCloud manages requests to the UpCloud API
type UpCloud struct {
	hc   http.Client
	host *url.URL

	username string
	password string
}

func (u *UpCloud) getURL(endpoint string) (url string) {
	reqURL := *u.host
	reqURL.Path = path.Join(APIVersion, endpoint)
	return reqURL.String()
}

func (u *UpCloud) request(method, endpoint string, body io.Reader, resp interface{}) (err error) {
	var req *http.Request
	if req, err = http.NewRequest(method, u.getURL(endpoint), body); err != nil {
		return
	}

	req.SetBasicAuth(u.username, u.password)

	var res *http.Response
	if res, err = u.hc.Do(req); err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return u.handleError(res.Body)
	}

	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return
	}

	return
}

func (u *UpCloud) handleError(body io.Reader) (err error) {
	var errResp errorResponse
	if err = json.NewDecoder(body).Decode(&errResp); err != nil {
		return
	}

	return errResp.Error.Error()
}

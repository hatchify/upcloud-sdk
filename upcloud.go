package upcloud

import (
	"context"
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

const (
	// RouteGetAccount is the route for getting current account
	RouteGetAccount = "account"
)

// New will return a new instance of the UpCloud API SDK
func New(username, password string) (up *UpCloud, err error) {
	var u UpCloud
	if u.host, err = url.Parse(Hostname); err != nil {
		return
	}

	// Set username
	u.username = username
	// Set password
	u.password = password
	// Assign pointer reference
	up = &u
	return
}

// UpCloud manages requests to the UpCloud API
type UpCloud struct {
	hc   http.Client
	host *url.URL

	// Login credentials
	username string
	password string
}

func (u *UpCloud) request(ctx context.Context, method, endpoint string, body io.Reader, resp interface{}) (err error) {
	var req *http.Request
	// Create a new request
	if req, err = u.newHTTPRequest(ctx, method, u.getURL(endpoint), body); err != nil {
		// Error encountered while creating new HTTP request, return
		return
	}

	var res *http.Response
	// Perform request using SDK's underlying HTTP client
	if res, err = u.hc.Do(req); err != nil {
		// Error encountered while performing request, return
		return
	}
	// Defer closing the HTTP response body
	defer res.Body.Close()

	// Process HTTP response from UpCloud API
	return u.processResponse(res, resp)
}

func (u *UpCloud) newHTTPRequest(ctx context.Context, method, url string, body io.Reader) (req *http.Request, err error) {
	// Create a new request using provided method, url, and body
	if req, err = http.NewRequest(method, url, body); err != nil {
		// Error encoutered while creating new HTTP request, return
		return
	}

	// The provided req must be non-nil
	if req == nil {
		panic("nil http.Request")
	}

	// apply context to the the http.Request
	req = req.WithContext(ctx)
	// Set API authentication using the username/password provided at SDK initialization
	req.SetBasicAuth(u.username, u.password)
	return
}

func (u *UpCloud) getURL(endpoint string) (url string) {
	// Create copy of host url.URL by derefencing source pointer
	reqURL := *u.host
	// Set the url path by concatinating the api version and the provided endpoint
	reqURL.Path = path.Join(APIVersion, endpoint)
	// Return the string representation of the built url
	return reqURL.String()
}

func (u *UpCloud) processResponse(res *http.Response, value interface{}) (err error) {
	// Check to see if error was successful
	if res.StatusCode >= 400 {
		// Error status code encountered, process as error
		return u.processError(res.Body)
	}

	// Initialize new JSON decoder and attempt to decode as provided value
	err = json.NewDecoder(res.Body).Decode(&value)
	return
}

func (u *UpCloud) processError(body io.Reader) (err error) {
	var errResp errorResponse
	// Initialize new JSON decoder and attempt to decode as an error response
	if err = json.NewDecoder(body).Decode(&errResp); err != nil {
		// Error encountered while decoding, return
		return
	}

	// Set returning error as the error response's Error value
	err = errResp.Error
	return
}

// GetAccount will get the account of the currently logged in user
func (u *UpCloud) GetAccount(ctx context.Context) (a *Account, err error) {
	var resp getAccountResponse
	// Make request to "Get Account" route
	if err = u.request(ctx, "GET", RouteGetAccount, nil, &resp); err != nil {
		return
	}

	// Set return value from response
	a = resp.Account
	return
}

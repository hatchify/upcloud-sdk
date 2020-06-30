package upcloud

import (
	"encoding/json"
	"github.com/hatchify/requester"
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
	// RouteGetZone gets all the zones
	RouteGetZone = "zone"
	// RouteGetPlan gets all the plans
	RouteGetPlan = "plan"
)

// New will return a new instance of the UpCloud API SDK
func New(username, password string) (up *UpCloud, err error) {
	var u UpCloud

	u.req = requester.New(&http.Client{}, Hostname)

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
	req  *requester.Requester
	host *url.URL

	// Login credentials
	username string
	password string
}

func (u *UpCloud) request(method, endpoint string, body []byte, resp interface{}) (err error) {
	var res *http.Response

	// We authenticate with BasicAuth
	var setBasicAuth requester.Modifier = func(request *http.Request, client *http.Client) (err error) {
		request.SetBasicAuth(u.username, u.password)
		return nil
	}

	if res, err = u.req.Request(method, u.getURL(endpoint), body, requester.Opts{setBasicAuth}); err != nil {
		return
	}
	// Defer closing the HTTP response body
	defer res.Body.Close()

	// Process HTTP response from UpCloud API
	return u.processResponse(res, resp)
}

func (u *UpCloud) getURL(endpoint string) (url string) {
	// Set the url path by concatenating the api version and the provided endpoint
	return path.Join(APIVersion, endpoint)
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
func (u *UpCloud) GetAccount() (a *Account, err error) {
	var resp getAccountResponse
	// Make request to "Get Account" route
	if err = u.request("GET", RouteGetAccount, nil, &resp); err != nil {
		return
	}

	// Set return value from response
	a = resp.Account
	return
}

// GetZones
func (u *UpCloud) GetZones() (z *[]Zone, err error) {
	var resp getZonesResponse
	// Make request to "Get Zones" route
	if err = u.request("GET", RouteGetZone, nil, &resp); err != nil {
		return
	}

	// Set return value from response
	z = resp.ZonesWrapper.Zone
	return
}

// GetPlans
func (u *UpCloud) GetPlans() (p *[]Plan, err error) {
	var resp getPlansResponse
	// Make request to "Get Zones" route
	if err = u.request("GET", RouteGetPlan, nil, &resp); err != nil {
		return
	}

	// Set return value from response
	p = resp.PlansWrapper.Plan
	return
}

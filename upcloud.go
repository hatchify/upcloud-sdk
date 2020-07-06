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
	// RouteGetServerSize gets all the server sizes
	RouteGetServerSize = "server_size"
	// RouteServer manages all the servers
	RouteServer = "server"
)

// RouteGetStorageFilter gets all the storage options for the server
type RouteGetStorageFilter string
const (
	All RouteGetStorageFilter = "storage"
	Public RouteGetStorageFilter = "storage/public"
	Private RouteGetStorageFilter = "storage/private"
	Normal RouteGetStorageFilter = "storage/normal"
	Backup RouteGetStorageFilter = "storage/backup"
	Cdrom RouteGetStorageFilter = "storage/cdrom"
	Template RouteGetStorageFilter = "storage/template"
	Favorite RouteGetStorageFilter = "storage/favorite"
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

func (u *UpCloud) SetRequester(newReq requester.Interface) {
	u.req = newReq
}

// UpCloud manages requests to the UpCloud API
type UpCloud struct {
	req  requester.Interface
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
	z = resp.Zones.Zone
	return
}

// GetPlans
func (u *UpCloud) GetPlans() (p *[]Plan, err error) {
	var resp getPlansResponse
	// Make request to "Get Plans" route
	if err = u.request("GET", RouteGetPlan, nil, &resp); err != nil {
		return
	}

	// Set return value from response
	p = resp.Plans.Plan
	return
}

// GetServerSizes
func (u *UpCloud) GetServerSizes() (p *[]ServerSize, err error) {
	var resp getServerSizesResponse
	// Make request to "Get Server Sizes" route
	if err = u.request("GET", RouteGetServerSize, nil, &resp); err != nil {
		return
	}

	// Set return value from response
	p = resp.ServerSizes.ServerSize
	return
}

// GetServers
func (u *UpCloud) GetServers() (p *[]Server, err error) {
	var resp getServersResponse
	// Make request to "Get Servers" route
	if err = u.request("GET", RouteServer, nil, &resp); err != nil {
		return
	}

	// Set return value from response
	p = resp.Servers.Server
	return
}

// GetServerDetails requires UUID of the server to get the details
func (u *UpCloud) GetServerDetails(uuid string) (p *ServerDetails, err error) {
	var resp serverDetailsWrapper
	// Make request to "Get Servers" route
	if err = u.request("GET", path.Join(RouteServer, uuid), nil, &resp); err != nil {
		return
	}

	// Set return value from response
	p = resp.ServerDetails
	return
}

// GetStorages
func (u *UpCloud) GetStorages(filter RouteGetStorageFilter) (p *[]Storage, err error) {
	var resp getStoragesResponse
	// Make request to "Get Servers" route
	if err = u.request("GET", string(filter), nil, &resp); err != nil {
		return
	}

	// Set return value from response
	p = resp.Storages.Storage
	return
}


// CreateServer
func (u *UpCloud) CreateServer(serverDetails *ServerDetails) (p *ServerDetails, err error) {

	//Dress up our new server in and wrap it
	var req = serverDetailsWrapper{
		ServerDetails: serverDetails,
	}
	var reqJson, _ = json.Marshal(req)

	var resp serverDetailsWrapper
	//Let's go and make us a server
	if err = u.request("POST", RouteServer, reqJson, &resp); err != nil {
		return
	}

	// Set return value from response
	p = resp.ServerDetails
	return
}




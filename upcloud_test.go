package upcloud

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/hatchify/requester/mock"
)

const (
	machineHostname = "sdk-test-machine"
)

func setup(t *testing.T) (u *UpCloud) {
	var err error

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal("Couldn't create UpCloud object")
	}

	var store = mock.NewFileStore("testdata/test-machine-full-run.json")
	store.Load()

	u.SetRequester(mock.NewMock(&http.Client{}, Hostname, store))
	//u.SetRequester(requester.NewSpy(&http.Client{}, Hostname, store))

	return
}

func TestNew(t *testing.T) {
	var err error
	var username string
	var password string
	var ok bool

	// Get username from OS environment
	if username, ok = os.LookupEnv("UPCLOUD_USERNAME"); !ok {
		t.Fatal("$UPCLOUD_USERNAME not set")
	}

	// Get password from OS environment
	if password, ok = os.LookupEnv("UPCLOUD_PASSWORD"); !ok {
		t.Fatal("$UPCLOUD_PASSWORD not set")
	}

	if _, err = New(username, password); err != nil {
		t.Fatal(err)
	}
}

func TestUpcloud_GetAccount(t *testing.T) {

	var err error
	u := setup(t)

	var a *Account
	if a, err = u.GetAccount(); err != nil {
		t.Fatal(err)
	}
	t.Log(a)

}

func ExampleNew() {
	var (
		u   *UpCloud
		err error
	)

	// Initialize new instance of UpCloud SDK
	if u, err = New("username", "password"); err != nil {
		// Error encountered while initializing SDK, return
		log.Fatal(err)
	}

	// UpCloud SDK is now ready to use!
	fmt.Println("UpCloud SDK is now ready to use!", u)
}

func ExampleUpCloud_GetAccount() {
	var (
		u   *UpCloud
		err error
	)

	// Initialize new instance of UpCloud SDK
	if u, err = New("username", "password"); err != nil {
		// Error encountered while initializing SDK, return
		log.Fatal(err)
	}

	var a *Account
	// Get account information of currently logged in user
	if a, err = u.GetAccount(); err != nil {
		// Error encountered while getting account information
		log.Fatal(err)
	}

	// Log account information
	fmt.Printf("My username is %s and I have %s credits", a.Username, a.Credits)
}

func TestUpCloud_GetZones(t *testing.T) {

	var err error
	u := setup(t)

	var zones *[]Zone
	// Get zones information of currently logged in user
	if zones, err = u.GetZones(); err != nil {
		// Error encountered while getting zones information
		t.Fatal(err)
	}

	for _, z := range *zones {
		if z.ID == "de-fra1" && z.Description == "Frankfurt #1" {
			t.Log(z.ID)
			// Output: de-fra1
		}
	}
}

func TestUpCloud_GetPlans(t *testing.T) {

	var err error
	u := setup(t)

	var plans *[]Plan
	// Get plans information of currently logged in user
	if plans, err = u.GetPlans(); err != nil {
		// Error encountered while getting account information
		t.Fatal(err)
	}

	for _, p := range *plans {
		if p.Name == "1xCPU-2GB" && p.CoreNumber == 1 {
			t.Log(p.Name)
			// Output: 1xCPU-2GB
		}
	}
}

func TestUpCloud_GetServerSizes(t *testing.T) {

	var err error
	u := setup(t)

	var serverSizes *[]ServerSize
	// Get plans information of currently logged in user
	if serverSizes, err = u.GetServerSizes(); err != nil {
		// Error encountered while getting account information
		t.Fatal(err)
	}

	for _, ss := range *serverSizes {
		if ss.MemoryAmount == "2048" && ss.CoreNumber == "1" {
			t.Log("We found our 2048mb with 1 core machine!")
			// Output: We found our 2048mb with 1 core machine!
		}
	}
}

func TestUpCloud_GetServers(t *testing.T) {

	var err error
	u := setup(t)

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		t.Fatal(err)
	}

	for _, s := range *servers {
		t.Log(s)
	}
}

func TestUpCloud_GetServerDetails(t *testing.T) {

	var err error
	u := setup(t)

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		t.Fatal(err)
	}

	var oneWeFound = (*servers)[0].UUID

	//Debug
	t.Log((*servers)[0])

	var serverDetails *ServerDetails
	// Get servers of currently logged in user
	if serverDetails, err = u.GetServerDetails(oneWeFound); err != nil {
		// Error encountered while getting account information
		t.Fatal(err)
	}

	if serverDetails.Hostname == "hatch-dev" {
		t.Log("found our machine in server details")
	}
}

func TestUpCloud_GetStorages(t *testing.T) {

	var err error
	u := setup(t)

	var storages *[]Storage
	// Get storages
	if storages, err = u.GetStorages(Public); err != nil {
		// Error encountered while getting storages
		t.Fatal(err)
	}

	t.Log((*storages)[0].Access == "public")
}

func TestUpCloud_CreateServer(t *testing.T) {

	var err error
	u := setup(t)

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		t.Fatal(err)
	}

	for _, server := range *servers {

		if machineHostname == server.Hostname {
			t.Skip("server with this name already exists!")
		}
	}

	var networking = &Networking{
		Interfaces: &Interfaces{
			Interface: &[]Interface{{
				IPAddresses: &IPAddresses{
					IPAddress: &[]IPAddress{{
						Family: "IPv4",
					}}},
				Type: "public",
			}},
		}}

	var storage = &StorageDevices{
		StorageDevice: &[]StorageDevice{{
			Action:  "clone",
			Storage: "01000000-0000-4000-8000-000030200200",
			Title:   "MadFastStripedRaid",
		}}}

	var serverDetails = &ServerDetails{
		Hostname:       machineHostname,
		Networking:     networking,
		StorageDevices: storage,
		Title:          machineHostname,
		Zone:           "us-chi1",
	}

	var result *ServerDetails
	// Get servers of currently logged in user
	if result, err = u.CreateServer(serverDetails); err != nil {
		// Error encountered while getting servers
		t.Fatal(err)
	}

	if result.Hostname == machineHostname {
		t.Log("found our machine in server details")
	}
}

func TestUpCloud_CreateServerWithMocks(t *testing.T) {
	//Implement mock based server creation very similar to regular CreateServer Test
	//We can also implement dynamic parameters by taking requests from FileStore and dumping them into MapStore

	//1. Create server request gets sent and a 200 response comes back for it
	//2. Make sure the server name matches in the response that comes back
}

func TestUpCloud_FullServerCreationCycleWithMocks(t *testing.T) {
	//Implement a mock based full working cycle for create/stop/start/stop/delete
	//This should probably use several JsonFileStores to have different started/stopped states
	//We should use already existing test and just doing t.Run() on them
}

func ExampleUpCloud_CreateServer() {

	var (
		u   *UpCloud
		err error
	)
	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		log.Fatal("Couldn't create UpCloud object")
	}

	var networking = &Networking{
		Interfaces: &Interfaces{
			Interface: &[]Interface{{
				IPAddresses: &IPAddresses{
					IPAddress: &[]IPAddress{{
						Family: "IPv4",
					}}},
				Type: "public",
			}},
		}}

	var storage = &StorageDevices{
		StorageDevice: &[]StorageDevice{{
			Action:  "clone",
			Storage: "01000000-0000-4000-8000-000030200200",
			Title:   "MadFastStripedRaid",
		}}}

	var serverDetails = &ServerDetails{
		Hostname:       machineHostname,
		Networking:     networking,
		StorageDevices: storage,
		Title:          machineHostname,
		Zone:           "us-chi1",
	}

	var result *ServerDetails
	// Get servers of currently logged in user
	if result, err = u.CreateServer(serverDetails); err != nil {
		// Error encountered while getting servers
		log.Fatal(err)
	}

	if result.Hostname == machineHostname {
		fmt.Println("found our machine in server details")
	}
}

func TestUpCloud_StartServer(t *testing.T) {

	var err error
	u := setup(t)

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		t.Fatal(err)
	}

	for _, server := range *servers {

		if machineHostname == server.Hostname {
			//Debug
			t.Log(server)

			if server.State == "stopped" {
				var serverDetails *ServerDetails
				// Get servers details of the server we are about to stop
				serverDetails, err = u.StartServer(server.UUID, StartServer{})
				if err != nil {
					// Error encountered while stopping the server
					t.Fatal(err)
				}

				if serverDetails.UUID == server.UUID {
					t.Log("found our matching stopped server and we started it")
					return
				}
			} else {
				t.Skip("server is in a bad state to start (aka \"maintenance\" or \"started\")")
			}
		}
	}
}

func TestUpCloud_StopServer(t *testing.T) {

	var err error
	u := setup(t)

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		t.Fatal(err)
	}

	for _, server := range *servers {
		if machineHostname == server.Hostname {
			//Debug
			t.Log(server)

			if server.State == "started" {
				var serverDetails *ServerDetails
				// Get servers details of the server we are about to stop
				serverDetails, err = u.StopServer(server.UUID, StopServer{StopType: string(Hard), Timeout: "60"})
				if err != nil {
					// Error encountered while stopping the server
					log.Fatal(err)
				}

				if serverDetails.UUID == server.UUID {
					t.Log("found our matching started server and we stopped it")
					return
				}
			} else {
				t.Skip("server is in a bad state to start (aka \"maintenance\" or \"stopped\")")
			}
		}
	}
}

func TestUpCloud_DeleteServer(t *testing.T) {

	var err error
	u := setup(t)

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		t.Fatal(err)
	}

	for _, server := range *servers {
		if server.Hostname == machineHostname {
			if server.State == "stopped" {
				if err = u.DeleteServer(server.UUID, true); err != nil {
					t.Fatal("Unable to delete the server")
				}
			} else {
				t.Skip("server is in a bad state to delete (aka \"maintenance\" or \"started\")")
			}
		}
	}
}

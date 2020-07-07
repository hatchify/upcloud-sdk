package upcloud

import (
	"fmt"
	"github.com/hatchify/requester"
	"log"
	"net/http"
	"os"
	"testing"
)

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
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	var a *Account
	if a, err = u.GetAccount(); err != nil {
		t.Fatal(err)
	}

	if a.Username != username {
		t.Fatalf("invalid username, expected %s and received %s", username, a.Username)
	}
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
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	//u.SetRequester(requester.NewMock(&http.Client{}, Hostname, requester.NewJsonFileStore("testdata/test.json")))

	var zones *[]Zone
	// Get zones information of currently logged in user
	if zones, err = u.GetZones(); err != nil {
		// Error encountered while getting zones information
		log.Fatal(err)
	}

	for _, z := range *zones {
		if z.ID == "de-fra1" && z.Description == "Frankfurt #1" {
			fmt.Println(z.ID)
			// Output: de-fra1
		}
	}
}

func TestUpCloud_GetPlans(t *testing.T) {
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	var plans *[]Plan
	// Get plans information of currently logged in user
	if plans, err = u.GetPlans(); err != nil {
		// Error encountered while getting account information
		log.Fatal(err)
	}

	for _, p := range *plans {
		if p.Name == "1xCPU-2GB" && p.CoreNumber == 1 {
			fmt.Println(p.Name)
			// Output: 1xCPU-2GB
		}
	}
}

func TestUpCloud_GetServerSizes(t *testing.T) {
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	var serverSizes *[]ServerSize
	// Get plans information of currently logged in user
	if serverSizes, err = u.GetServerSizes(); err != nil {
		// Error encountered while getting account information
		log.Fatal(err)
	}

	for _, ss := range *serverSizes {
		if ss.MemoryAmount == "2048" && ss.CoreNumber == "1" {
			fmt.Println("We found our 2048mb with 1 core machine!")
			// Output: We found our 2048mb with 1 core machine!
		}
	}
}

func TestUpCloud_GetServers(t *testing.T) {
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		log.Fatal(err)
	}

	for _, s := range *servers {
		fmt.Println(s)
	}
}

func TestUpCloud_GetServerDetails(t *testing.T) {
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	//u.SetRequester(requester.NewMock(&http.Client{}, Hostname, requester.NewJsonFileStore("testdata/test.json")))

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		log.Fatal(err)
	}

	var oneWeFound = (*servers)[0].UUID

	//Debug
	fmt.Println((*servers)[0])

	var serverDetails *ServerDetails
	// Get servers of currently logged in user
	if serverDetails, err = u.GetServerDetails(oneWeFound); err != nil {
		// Error encountered while getting account information
		log.Fatal(err)
	}

	if serverDetails.Hostname == "hatch-dev" {
		fmt.Println("found our machine in server details")
	}
}

func TestUpCloud_GetStorages(t *testing.T) {
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	var storages *[]Storage
	// Get storages
	if storages, err = u.GetStorages(Public); err != nil {
		// Error encountered while getting storages
		log.Fatal(err)
	}

	fmt.Println((*storages)[0].Access == "public")
}

func TestUpCloud_CreateServer(t *testing.T) {
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	//u.SetRequester(requester.NewMock(&http.Client{}, Hostname, requester.NewJsonFileStore("testdata/test_post.json")))

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
		Hostname:       "sergey-test",
		Networking:     networking,
		StorageDevices: storage,
		Title:          "SergeyTest",
		Zone:           "us-chi1",
	}

	var result *ServerDetails
	// Get servers of currently logged in user
	if result, err = u.CreateServer(serverDetails); err != nil {
		// Error encountered while getting servers
		log.Fatal(err)
	}

	if result.Hostname == "sergey-test" {
		fmt.Println("found our machine in server details")
	}
}

func TestUpCloud_StopServer(t *testing.T) {
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	//u.SetRequester(requester.NewMock(&http.Client{}, Hostname, requester.NewJsonFileStore("testdata/test_post.json")))

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		log.Fatal(err)
	}

	var oneWeFound = (*servers)[0].UUID

	//Debug
	fmt.Println((*servers)[0])

	var serverDetails *ServerDetails
	// Get servers details of the server we are about to stop
	serverDetails, err = u.StopServer(oneWeFound, StopServer{StopType: string(Soft), Timeout:  "60"})
	if err != nil {
		// Error encountered while stopping the server
		log.Fatal(err)
	}

	if serverDetails.UUID == (*servers)[0].UUID {
		fmt.Println("found our matching stopping server")
	}
}

func TestUpCloud_StartServer(t *testing.T) {
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	//u.SetRequester(requester.NewMock(&http.Client{}, Hostname, requester.NewJsonFileStore("testdata/test_post.json")))

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		log.Fatal(err)
	}

	var oneWeFound = (*servers)[0].UUID

	//Debug
	fmt.Println((*servers)[0])

	var serverDetails *ServerDetails
	// Get servers details of the server we are about to stop
	serverDetails, err = u.StartServer(oneWeFound, StartServer{})
	if err != nil {
		// Error encountered while stopping the server
		log.Fatal(err)
	}

	if serverDetails.UUID == (*servers)[0].UUID {
		fmt.Println("found our matching stopping server")
	}
}


func TestUpCloud_DeleteServer(t *testing.T) {
	var (
		u   *UpCloud
		err error
	)

	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	u.SetRequester(requester.NewSpy(&http.Client{}, Hostname, requester.NewJsonFileStore("testdata/test_post.json")))

	var servers *[]Server
	// Get servers of currently logged in user
	if servers, err = u.GetServers(); err != nil {
		// Error encountered while getting servers
		log.Fatal(err)
	}

	var oneWeFoundHostname = (*servers)[0].Hostname
	var oneWeFoundUUID = (*servers)[0].UUID

	if "sergey-test" == oneWeFoundHostname {
		if err = u.DeleteServer(oneWeFoundUUID, true); err != nil {
			fmt.Println("things didn't go well in deleting")
			log.Fatal(err)
		}

	}
}
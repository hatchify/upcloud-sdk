package upcloud

import (
	"fmt"
	"log"
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

	var zones *[]Zone
	// Get zones information of currently logged in user
	if zones, err = u.GetZones(); err != nil {
		// Error encountered while getting account information
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

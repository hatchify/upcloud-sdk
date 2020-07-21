package upcloud

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	var err error
	// Get username from OS environment
	username := os.Getenv("UPCLOUD_USERNAME")
	// Get password from OS environment
	password := os.Getenv("UPCLOUD_PASSWORD")

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

	ctx := context.Background()
	if u, err = New(username, password); err != nil {
		t.Fatal(err)
	}

	var a *Account
	if a, err = u.GetAccount(ctx); err != nil {
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

	// ctx := context.Background()
	ctx := context.Background()
	var a *Account
	// Get account information of currently logged in user
	if a, err = u.GetAccount(ctx); err != nil {
		// Error encountered while getting account information
		log.Fatal(err)
	}

	// Log account information
	fmt.Printf("My username is %s and I have %s credits", a.Username, a.Credits)
}

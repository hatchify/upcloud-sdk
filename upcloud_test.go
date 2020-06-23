package upcloud

import (
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

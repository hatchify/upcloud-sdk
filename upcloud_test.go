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

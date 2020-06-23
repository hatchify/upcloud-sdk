# UpCloud SDK [![GoDoc](https://godoc.org/github.com/hatchify/upcloud-sdk?status.svg)](https://godoc.org/github.com/hatchify/upcloud-sdk) ![Status](https://img.shields.io/badge/status-beta-yellow.svg)

UpCloud SDK is an SDK wrapper for the UpCloud API

## Usage 
### New
```go
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
```

### UpCloud.GetAccount
```go

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
```
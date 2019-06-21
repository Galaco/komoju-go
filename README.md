[![GoDoc](https://godoc.org/github.com/Galaco/komoju-go?status.svg)](https://godoc.org/github.com/Galaco/komoju-go)
[![Go report card](https://goreportcard.com/badge/github.com/galaco/komoju-go)](https://goreportcard.com/badge/github.com/galaco/komoju-go)
[![GolangCI](https://golangci.com/badges/github.com/galaco/komoju-go.svg)](https://golangci.com)
[![Build Status](https://travis-ci.com/Galaco/komoju-go.svg?branch=master)](https://travis-ci.com/Galaco/komoju-go)
[![codecov](https://codecov.io/gh/Galaco/komoju-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Galaco/komoju-go)
[![CircleCI](https://circleci.com/gh/Galaco/komoju-go.svg?style=svg)](https://circleci.com/gh/Galaco/komoju-go)

# komoju-go
Unoffical Golang wrapper for Komoju api. Based on documentation available here: [https://docs.komoju.com/en/getting_started/overview/#products](https://docs.komoju.com/en/getting_started/overview/#products)

A work-in-progress written mostly for fun. Any bugs found please feel free to open an issue and/or raise a PR.

### Usage

#### Hosted Page
```go
package main

import (
	komoju "github.com/galaco/komoju-go"
)

func main() {
	locale := "en"
	merchantUuid := "degica-mart"
	paymentMethod := komoju.PaymentMethodCreditCard
	baseUri := komoju.HostedPageBaseURI(locale, merchantUuid, paymentMethod)
	
	// your secret key
	secretKey := []byte("ABCD1234567890")
	config := komoju.HostedPageConfig{
    	Transaction: komoju.HostedPageTransaction{
    		Amount: 130,
    		Currency: "JPY",
    		Customer: komoju.HostedPageCustomer{
    			GivenName: "John",
    			FamilyName: "Smith",
    			GivenNameKana: "John",
    			FamilyNameKana: "Smith",
    			Email: "",
    			Phone: "",
    		},
    		ExternalOrderNumber: "M8x6U6Z5HEeXv3",
    		ReturnUrl: "http://example.com/?sucess=true",
    		CancelUrl: "http://example.com/?cancel=true",
    		Tax: 0,
    	},
        Timestamp: 1561022519,
    }
	
	paymentUri,_ := komoju.HostedPageURI(secretKey, baseUri, &config)
	// Now redirect to paymentUri
}
````

#### API
```go
package main

import (
	komoju "github.com/galaco/komoju-go"
	"log"
)	
func main() {
    // ... make a request to the komoju API
    // expect a standard response.Body as io.ReadCloser

    payload := komoju.ApiResponseEvents{}
    // remember to handle error
    _ := komoju.ParseApiResponse(response.Body, &payload)
    
    // Do whatever with response
    log.Println(payload.Data[0].Resource)
}
```


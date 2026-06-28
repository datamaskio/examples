// Example: basic struct redaction with mask tags.
//
//	go run github.com/datamaskio/examples/basic
package main

import (
	"fmt"

	"github.com/datamaskio/datamask"
)

type Customer struct {
	Name  string
	Email string `mask:"email"`
	Phone string `mask:"phone"`
	SSN   string `mask:"secret"`
}

func main() {
	masker := datamask.New()

	customer := Customer{
		Name:  "John Doe",
		Email: "john@example.com",
		Phone: "+1-555-1234",
		SSN:   "123-45-6789",
	}

	masked, err := datamask.Redact(masker, customer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original: %+v\n", customer)
	fmt.Printf("  Masked: %+v\n", masked)
}

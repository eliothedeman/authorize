package cim

import (
	"log"
	"testing"

	"github.com/eliothedeman/authorize"
)

func TestCreateCustomerProfile(t *testing.T) {
	c := authorize.NewTestClient()
	r := &CreateCustomerProfileRequest{}
	r.Profile.PaymentProfiles.CustomerType = "individual"
	r.Profile.Email = "test@gmail.com"
	err := c.Do(r)

	if err != nil {
		t.Fail()
	}

	log.Println(r.Response)

}

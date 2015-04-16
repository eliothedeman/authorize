package cim

import (
	"log"
	"testing"

	"github.com/eliothedeman/authorize"
)

func TestGetProfileTest(t *testing.T) {
	c := authorize.NewTestClient()
	r := &GetCustomerProfileRequest{}
	p := createRandomeProfile()

	r.CustomerProfileId = p.MerchantCustomerId
	log.Println(p.MerchantCustomerId)

	resp := c.Do(r)

	if resp.Err != nil {
		t.Error(resp.Err)
	}
}

package authorize

import (
	"log"
	"testing"
)

func TestGetCustomerProfile(t *testing.T) {
	c := NewTestClient()
	id := createRandomeProfile()
	_, err := c.GetCustomerProfile(id)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCustomerProfileBadId(t *testing.T) {
	c := NewTestClient()
	id := randomNumberString(10)
	r, err := c.GetCustomerProfile(id)
	if err != INVALID_CONTENT {
		t.Error(err)
	}

	log.Println(r)
}

package authorize

import "testing"

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
	_, err := c.GetCustomerProfile(id)
	if err != INVALID_CONTENT {
		t.Error(err)
	}
}

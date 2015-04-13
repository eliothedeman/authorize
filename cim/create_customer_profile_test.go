package cim

import (
	"math/rand"
	"testing"

	"github.com/eliothedeman/authorize"
)

const (
	TEST_AMEX = "370000000000002"
	TEST_VISA = "4007000000027"
)

func randomString(length int) string {
	str := make([]byte, length)
	for i := range str {
		str[i] = uint8(rand.Int()%26) + 66
	}
	return string(str)
}

func randomNumberString(length int) string {
	str := make([]byte, length)
	for i := range str {
		str[i] = uint8(rand.Int()%9) + 48
	}

	return string(str)
}

func TestCreateCustomerProfile(t *testing.T) {
	c := authorize.NewTestClient()
	r := &CreateCustomerProfileRequest{}
	r.Profile.PaymentProfile.CustomerType = "individual"
	r.Profile.Email = randomString(10) + "@gmail.com"
	cred := &CreditCard{}
	cred.CardCode = "134"
	cred.CardNumber = randomString(13)
	cred.ExpirationDate = "2020-01"
	r.Profile.PaymentProfile.Payment.CreditCard = cred
	r.Profile.PaymentProfile.BillTo.Address = randomString(10)
	r.Profile.PaymentProfile.BillTo.FirstName = randomString(10)
	r.Profile.PaymentProfile.BillTo.LastName = randomString(10)
	r.Profile.PaymentProfile.BillTo.Address = randomString(10)
	r.Profile.PaymentProfile.BillTo.Company = randomString(10)

	resp := c.Do(r)

	if resp.Error() != "" {
		t.Error(resp.Error())
	}
}

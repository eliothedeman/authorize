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

func randomPaymenProfile() *PaymentProfile {
	cred := &CreditCard{}
	cred.CardCode = "134"
	cred.CardNumber = randomNumberString(13)
	cred.ExpirationDate = "2020-01"
	paymentProfile := &PaymentProfile{}
	paymentProfile.Payment.CreditCard = cred
	paymentProfile.BillTo.Address = randomString(10)
	paymentProfile.BillTo.FirstName = randomString(10)
	paymentProfile.BillTo.LastName = randomString(10)
	paymentProfile.BillTo.Address = randomString(10)
	paymentProfile.BillTo.Company = randomString(10)
	paymentProfile.CustomerType = "individual"
	return paymentProfile
}

func randomProfile() Profile {
	p := Profile{}
	p.Email = randomString(10) + "@gmail.com"
	p.MerchantCustomerId = randomNumberString(10)
	p.PaymentProfile = randomPaymenProfile()
	return p
}

func TestCreateCustomerProfile(t *testing.T) {
	c := authorize.NewTestClient()
	r := &CreateCustomerProfileRequest{}
	r.Profile = randomProfile()

	resp := c.Do(r)

	if resp.Error() != "" {
		t.Error(resp.Error())
	}
}

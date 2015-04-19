package authorize

import (
	"math/rand"
	"testing"

	"github.com/eliothedeman/authorize/cim"
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

func randomPaymenProfile() *cim.PaymentProfile {
	cred := &cim.CreditCard{}
	cred.CardCode = "134"
	cred.CardNumber = randomNumberString(13)
	cred.ExpirationDate = "2020-01"
	paymentProfile := &cim.PaymentProfile{}
	paymentProfile.Payment.CreditCard = cred
	paymentProfile.BillTo.Address = randomString(10)
	paymentProfile.BillTo.FirstName = randomString(10)
	paymentProfile.BillTo.LastName = randomString(10)
	paymentProfile.BillTo.Address = randomString(10)
	paymentProfile.BillTo.Company = randomString(10)
	paymentProfile.CustomerType = "individual"
	return paymentProfile
}

func randomProfile() cim.Profile {
	p := cim.Profile{}
	p.Email = randomString(10) + "@gmail.com"
	p.MerchantCustomerId = randomNumberString(10)
	p.PaymentProfiles = []*cim.PaymentProfile{randomPaymenProfile()}
	return p
}

func randomCreateProfileRequest() *cim.CreateCustomerProfileRequest {
	r := &cim.CreateCustomerProfileRequest{}
	r.Profile = randomProfile()
	return r
}

// creates a new randome profile, and returns the id of the new profile
func createRandomeProfile() string {
	c := NewTestClient()
	id, _ := c.CreateCustomerProfile(randomProfile())
	return id
}

func TestCreateCustomerProfile(t *testing.T) {
	c := NewTestClient()
	id, err := c.CreateCustomerProfile(randomProfile())
	if err != nil {
		t.Error(err)
	}

	if id == "" {
		t.Fail()
	}
}

func TestCreateCustomerProfileBadCard(t *testing.T) {
	c := NewTestClient()
	p := randomProfile()
	p.PaymentProfiles[0].Payment.CreditCard.CardNumber = randomString(13)

	_, err := c.CreateCustomerProfile(p)

	if err != INVALID_CONTENT {
		t.Error(err)
	}
}

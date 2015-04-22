package authorize

import (
	"log"
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

func fakeAmex() *cim.CreditCard {
	c := &cim.CreditCard{}
	c.CardCode = "1234"
	c.CardNumber = TEST_AMEX
	c.ExpirationDate = "2020-01"
	return c
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

func randomTransaction(c_id, cpp_id, ship_id string) *cim.Transaction {
	t := &cim.Transaction{}
	t.Amount = randomNumberString(4)
	t.CustomerProfileId = c_id
	t.CustomerShippingAddressId = ship_id
	t.CustomerPaymentProfileId = cpp_id
	return t
}

func randomAddress() *cim.Address {
	a := &cim.Address{}
	a.Address = randomString(10)
	a.City = randomString(10)
	a.Company = randomString(10)
	a.Country = randomString(20)
	a.FirstName = randomString(10)
	a.LastName = randomString(10)
	a.PhoneNumber = randomNumberString(10)
	a.FaxNumber = randomNumberString(10)
	a.Zip = randomNumberString(4)
	return a
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
func createRandomProfile() string {
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

func TestGetCustomerProfile(t *testing.T) {
	c := NewTestClient()
	id := createRandomProfile()
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

func TestCreateCustomerPaymentProfile(t *testing.T) {
	c := NewTestClient()
	id := createRandomProfile()
	pp := randomPaymenProfile()
	_, err := c.CreateCustomerPaymentProfile(id, pp)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateCustomerPaymentProfileBadId(t *testing.T) {
	c := NewTestClient()
	pp := randomPaymenProfile()
	_, err := c.CreateCustomerPaymentProfile(randomNumberString(13), pp)

	if err != INVALID_CONTENT {
		t.Error(err)
	}
}

func TestCreateShippingAddress(t *testing.T) {
	c := NewTestClient()
	id := createRandomProfile()
	a := randomAddress()
	id, err := c.CreateCustomerSippingAddress(id, a)
	log.Println(id)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateCustomerPaymanetProfileTransaction(t *testing.T) {
	c := NewTestClient()
	id := createRandomProfile()
	pp := randomPaymenProfile()
	pp.Payment.CreditCard = fakeAmex()
	a := randomAddress()
	ship_id, _ := c.CreateCustomerSippingAddress(id, a)
	cpp_id, _ := c.CreateCustomerPaymentProfile(id, pp)
	trans := randomTransaction(id, cpp_id, ship_id)

	_, err := c.CreateCustomerPaymentProfileTransaction(trans)

	if err != nil {
		t.Error(err)
	}

}

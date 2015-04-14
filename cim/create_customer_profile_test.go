package cim

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/eliothedeman/authorize"
)

const (
	TEST_AMEX = "370000000000002"
	TEST_VISA = "4007000000027"
)

func init() {
	log.SetFlags(log.Llongfile)
	rand.Seed(time.Now().Unix())
}
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

func randomCreateProfileRequest() *CreateCustomerProfileRequest {
	r := &CreateCustomerProfileRequest{}
	r.Profile = randomProfile()
	return r
}

func TestCreateCustomerProfile(t *testing.T) {
	c := authorize.NewTestClient()
	r := randomCreateProfileRequest()
	resp := c.Do(r)

	if resp.Err != nil {
		t.Error(resp.Err)
	}
}

func TestCreateCustomerProfileBadCard(t *testing.T) {
	c := authorize.NewTestClient()
	r := randomCreateProfileRequest()
	r.Profile.PaymentProfile.Payment.CreditCard.CardNumber = randomString(13)
	resp := c.Do(r)

	if resp.Err != authorize.INVALID_CARD_NUMBER {
		t.Error("Expected INVALID_CARD_NUMBER got", resp.Err)
	}

}

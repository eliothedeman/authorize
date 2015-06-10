package authorize

import "testing"

var (
	test_cards = map[string]string{
		"v":  "4012888888881881",
		"m":  "5105105105105100",
		"a":  "371449635398431",
		"d":  "38520000023237",
		"di": "6011000990139424",
		"j":  "3566002020360505",
	}
)

func TestTypeVisa(t *testing.T) {
	x := CardType(test_cards["v"])
	if x != VISA {
		t.Fail()
	}
}

func TestTypeMastercard(t *testing.T) {
	x := CardType(test_cards["m"])
	if x != MASTERCARD {
		t.Fail()
	}
}

func TestTypeAmex(t *testing.T) {
	x := CardType(test_cards["a"])
	if x != AMEX {
		t.Fail()
	}
}

func TestTypeDiscover(t *testing.T) {
	x := CardType(test_cards["di"])
	if x != DISCOVER {
		t.Fail()
	}
}

func TestTypeJCB(t *testing.T) {
	x := CardType(test_cards["j"])
	if x != JCB {
		t.Fail()
	}
}

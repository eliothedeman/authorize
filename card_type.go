package authorize

import (
	"regexp"
	"strings"
)

// regex matches for each card type
var (
	VISA_REGEX        = regexp.MustCompile(`^4[0-9]{6,}`)
	MASTERCARD_REGEX  = regexp.MustCompile(`^5[1-5][0-9]{5,}`)
	AMEX_REGEX        = regexp.MustCompile(`^3[47][0-9]{13}`)
	DINERS_CLUB_REGEX = regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}`)
	DISCOVER_REGEX    = regexp.MustCompile(`^6(?:011|5[0-9]{2})[0-9]{12}`)
	JCB_REGEX         = regexp.MustCompile(`^(?:2131|1800|35\d{3})\d{11}`)
)

// card type codes
const (
	VISA = iota
	MASTERCARD
	AMEX
	DINERS_CLUB
	DISCOVER
	JCB
)

// convert a card number to the associated card type
func CardType(cardNum string) int {

	// remove any dashes from the card
	cardNum = strings.Replace(cardNum, "-", "", -1)

	switch cardNum[0] {
	case '4':
		if VISA_REGEX.MatchString(cardNum) {
			return VISA
		}
	case '5':
		if MASTERCARD_REGEX.MatchString(cardNum) {
			return MASTERCARD
		}
	case '3':
		if AMEX_REGEX.MatchString(cardNum) {
			return AMEX
		}
		if DINERS_CLUB_REGEX.MatchString(cardNum) {
			return DINERS_CLUB
		}

	case '6':
		if DISCOVER_REGEX.MatchString(cardNum) {
			return DISCOVER
		}

	}
	if JCB_REGEX.MatchString(cardNum) {
		return JCB
	}
	return -1
}

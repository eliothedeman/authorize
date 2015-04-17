package cim

// Customer Information Manger
type Profile struct {
	MerchantCustomerId string            `json:"merchantCustomerId,omitempty"`
	Description        string            `json:"description,omitempty"`
	Email              string            `json:"email,omitempty"`
	PaymentProfiles    []*PaymentProfile `json:"paymentProfiles,omitempty"`
}

type PaymentProfile struct {
	CustomerType string      `json:"customerType,omitempty"`
	BillTo       Billing     `json:"billTo,omitempty"`
	Payment      PaymentInfo `json:"payment"`
}

type PaymentInfo struct {
	CreditCard  *CreditCard  `json:"creditCard,omitempty"`
	BankAccount *BankAccount `json:"bankAccount,omitempty"`
}

type CreditCard struct {
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	CardCode       string `json:"cardCode,omitempty"`
}
type BankAccount struct {
	AccountType    string `json:"accountType,omitempty"`
	RouteingNumber string `json:"routingNumber,omitempty"`
	NameOnAccount  string `json:"nameOnAccount,omitempty"`
	ECheckType     string `json:"echeckType,omitempty"`
	BankName       string `json:"bankName,omitempty"`
}

type Billing struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Company     string `json:"company,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	FaxNumber   string `json:"faxNumber,omitempty"`
}

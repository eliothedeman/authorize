package auth

// Authenticate the client for a given merchant
type MerchantAuth struct {
	Name           string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}

// create a new authenticate request
func NewMerchantAuth(name, transactionKey string) *MerchantAuth {
	return &MerchantAuth{
		Name:           name,
		TransactionKey: transactionKey,
	}
}

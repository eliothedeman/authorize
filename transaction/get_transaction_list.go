package transaction

import "github.com/eliothedeman/authorize/auth"

type GetTransactionList struct {
	MerchAuth *auth.MerchantAuth `json:"merchantAuthentication"`
	BatchId   string             `json:"batchId"`
}

func NewGetTransactionList(batchId string) *GetTransactionList {
	return &GetTransactionList{
		BatchId: batchId,
	}
}

func (g *GetTransactionList) SetAuth(m *auth.MerchantAuth) {
	g.MerchAuth = m
}

func (g *GetTransactionList) Method() string {
	return "getTransactionListRequest"
}

func (g *GetTransactionList) EndPoint() string {
	return "/xml/v1/request.api"
}

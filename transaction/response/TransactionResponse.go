package kpay_response

import "time"

type TransactionStatus int

const (
	CREATED TransactionStatus = iota
	SUCCESS
	CANCELED
	FAIL
	TIMEOUT
)

type PAYMENT_TYPE string

const (
	VIET_QR = "VIET_QR"
	TM_CARD = "ATM_CARD"
	BANKING = "BANKING"
)

type CreateTransactionResponse struct {
	TransactionId    string `json:"transactionId"`
	RefTransactionId string `json:"refTransactionId"`
	PayLinkCode      string `json:"payLinkCode"`
	Timeout          int64  `json:"timeout"`
	Url              string `json:"url"`
	VirtualAccount   string `json:"virtualAccount"`
	Description      string `json:"description"`
	Amount           int64  `json:"amount"`
	QrCodeString     string `json:"qrCodeString"`
	Status           string `json:"status"`
	Time             string `json:"time"`
	AccountName      string `json:"accountName"`
}
type QueryTransactionResponse struct {
	Status           string `json:"status"`
	RefTransactionId string `json:"refTransactionId"`
	Amount           int64  `json:"amount"`
}

type CancelTransactionResponse struct {
	Success bool `json:"success"`
}

type CreateVirtualAccountResponse struct {
	Order          int64  `json:"order"`
	VirtualAccount string `json:"virtualAccount"`
	Timeout        int64  `json:"timeout"`
	FixAmount      int64  `json:"fixAmount"`
	FixContent     string `json:"fixContent"`
	QrContent      string `json:"qrContent"`
	BankAccountNo  string `json:"bankAccountNo"`
}

type DisableVirtualAccountResponse struct {
	Success bool `json:"success"`
}

type PayTransactionResponse struct {
	Id               string       `json:"id"`
	Status           string       `json:"status"`
	Amount           string       `json:"amount"`
	RefTransactionId string       `json:"refTransactionId"`
	CreateDateTime   time.Time    `json:"createDateTime"`
	CompleteTime     time.Time    `json:"completeTime"`
	VirtualAccount   string       `json:"virtualAccount"`
	Description      string       `json:"description"`
	PaymentType      PAYMENT_TYPE `json:"paymentType"`
	TxnNumber        string       `json:"txnNumber"`
	AccountName      string       `json:"accountName"`
	AccountNo        string       `json:"accountNo"`
	InterBankTrace   string       `json:"interBankTrace"`
}

type PaginateResponse[T any] struct {
	PageNumber int32 `json:"pageNumber"`
	PageSize   int32 `json:"pageSize"`
	TotalSize  int32 `json:"totalSize"`
	TotalPage  int32 `json:"totalPage"`
	Items      []T   `json:"items"`
}

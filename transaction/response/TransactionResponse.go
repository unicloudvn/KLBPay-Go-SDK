package kpay_response

type TransactionStatus int

const (
	CREATED TransactionStatus = iota
	SUCCESS
	CANCELED
	FAIL
	TIMEOUT
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
}
type QueryTransactionResponse struct {
	Status           string `json:"status"`
	RefTransactionId string `json:"refTransactionId"`
	Amount           int64  `json:"amount"`
}

type CancelTransactionResponse struct {
	Success bool `json:"success"`
}

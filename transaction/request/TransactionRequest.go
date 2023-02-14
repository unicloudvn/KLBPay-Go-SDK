package kpay_request

type BodyEncryptRequest struct {
	Data string `json:"data"`
}

type Customer struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type CreateTransactionRequest struct {
	RefTransactionId string   `json:"refTransactionId"`
	Amount           int64    `json:"amount"`
	Description      string   `json:"description"`
	Timeout          int      `json:"timeout"`
	Title            string   `json:"title"`
	Language         string   `json:"language"`
	SuccessUrl       string   `json:"successUrl"`
	CustomerInfo     Customer `json:"customerInfo"`
	FailUrl          string   `json:"failUrl"`
	RedirectAfter    int      `json:"redirectAfter"`
	BankAccountId    string   `json:"bankAccountId"`
}

type QueryTransactionRequest struct {
	TransactionId string `json:"transactionId"`
}

type CancelTransactionRequest struct {
	TransactionId string `json:"transactionId"`
}

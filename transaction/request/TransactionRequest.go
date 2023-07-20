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
	BankAccountNo    string   `json:"bankAccountNo"`
}

type QueryTransactionRequest struct {
	TransactionId string `json:"transactionId"`
}

type CancelTransactionRequest struct {
	TransactionId string `json:"transactionId"`
}

type CreateVirtualAccountRequest struct {
	Order         int64  `json:"order"`
	Timeout       int64  `json:"timeout"`
	FixAmount     int64  `json:"fixAmount"`
	FixContent    string `json:"fixContent"`
	BankAccountNo string `json:"bankAccountNo"`
}

type DisableVirtualAccountRequest struct {
	Order int64 `json:"order"`
}

type GetTransactionRequest struct {
	Size          int32  `json:"size"`
	Page          int32  `json:"page"`
	Order         *int64 `json:"order"`
	BankAccountNo string `json:"bankAccountNo"`
	FromDate      string `json:"fromDate"`
	ToDate        string `json:"toDate"`
}

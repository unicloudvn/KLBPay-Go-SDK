package kpay_model

type Message struct {
	ClientId     string
	Timestamp    int64
	EncryptData  string
	ValidateData string
}

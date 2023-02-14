package kpay_exception

type kPayResponseCode struct {
	Code    int
	Message string
}

var (
	// error common
	Success           = &kPayResponseCode{0, "Success"}
	Failed            = &kPayResponseCode{1, "Failed"}
	CallApiFailed     = &kPayResponseCode{2, "Call api to service failed"}
	InvalidParam      = &kPayResponseCode{2, "Invalid param"}
	DecodeKeyFailed   = &kPayResponseCode{3, "Decode key failed"}
	CreateCipherError = &kPayResponseCode{4, "Create block cipher error"}

	// error transaction
	SecurityViolation   = &kPayResponseCode{1601, "Security violation"}
	OrderCompleted      = &kPayResponseCode{1602, "Order was completed"}
	AmountInvalid       = &kPayResponseCode{1603, "Invalid amount"}
	TransactionCanceled = &kPayResponseCode{1604, "Canceled transaction"}
	TransactionExpired  = &kPayResponseCode{1605, "Transaction expired"}
	TransactionInvalid  = &kPayResponseCode{1606, "Invalid transaction"}
	TransactionFailed   = &kPayResponseCode{1607, "Transaction failed"}
	ServiceUnavailable  = &kPayResponseCode{1608, "Service unavailable"}
	InvalidClientID     = &kPayResponseCode{1609, "Invalid client id"}
)

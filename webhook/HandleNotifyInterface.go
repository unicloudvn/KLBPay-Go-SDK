package webhook

import (
	kpay_request "github.com/unicloudvn/KLBPay-Go-SDK/transaction/request"
)

type HandleNotify interface {
	Handle(request kpay_request.NotifyRequest)
}

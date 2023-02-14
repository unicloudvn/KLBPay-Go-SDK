package main

import (
	"fmt"
	kpay_service "kpay_sdk/service"
	kpay_request "kpay_sdk/transaction/request"
)

func main() {
	checkReq := kpay_request.QueryTransactionRequest{
		TransactionId: "4e128dea-8fa2-4867-8168-e3ed40c49367",
	}
	res, err := kpay_service.QueryTranasction(checkReq)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	fmt.Println(res)
}

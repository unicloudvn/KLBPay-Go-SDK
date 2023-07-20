package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	kpay_config "github.com/unicloudvn/KLBPay-Go-SDK/config"
	kpay_security "github.com/unicloudvn/KLBPay-Go-SDK/security"
	kpay_model "github.com/unicloudvn/KLBPay-Go-SDK/transaction/model"
	kpay_request "github.com/unicloudvn/KLBPay-Go-SDK/transaction/request"
	kpay_response "github.com/unicloudvn/KLBPay-Go-SDK/transaction/response"
	"io"
	"net/http"
	"strconv"
)

var (
	x_api_client   = "x-api-client"
	x_api_time     = "x-api-time"
	x_api_validate = "x-api-validate"
)

func execute(url string, message *kpay_model.Message) (kpay_model.Message, error) {
	bodyEncrypt := kpay_request.BodyEncryptRequest{
		Data: message.EncryptData,
	}
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(bodyEncrypt)
	if err != nil {
		return kpay_model.Message{}, err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return kpay_model.Message{}, err
	}
	req.Header.Add(x_api_client, message.ClientId)
	req.Header.Add(x_api_validate, message.ValidateData)
	req.Header.Add(x_api_time, fmt.Sprintf("%d", message.Timestamp))
	req.Header.Add("Content-type", "application/json")

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return kpay_model.Message{}, err
	}
	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		return kpay_model.Message{}, errors.New(res.Status)
	}

	defer res.Body.Close()
	bodyRes, _ := io.ReadAll(res.Body)

	var bodyEncryptRes kpay_request.BodyEncryptRequest
	err = json.Unmarshal(bodyRes, &bodyEncryptRes)
	if err != nil {
		return kpay_model.Message{}, err
	}

	timestampRes, _ := strconv.ParseInt(res.Header.Get(x_api_time), 10, 64)
	return kpay_model.Message{
		ClientId:     res.Header.Get(x_api_client),
		Timestamp:    timestampRes,
		ValidateData: res.Header.Get(x_api_validate),
		EncryptData:  bodyEncryptRes.Data,
	}, nil
}

func callApi[S any, T any](kPayConfig *kpay_config.KPayConfig, url string, req S, res T) (any, error) {
	messageRequest, err := kpay_security.Encode(kPayConfig, req)
	if err != nil {
		return nil, err
	}

	messageResponse, err := execute(url, messageRequest)
	if err != nil {
		return nil, err
	}

	result, err := kpay_security.Decode(kPayConfig, &messageResponse, &res)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateTransaction(kPayConfig *kpay_config.KPayConfig, request kpay_request.CreateTransactionRequest) (kpay_response.CreateTransactionResponse, error) {
	url := kPayConfig.KPayHost + "/api/payment/v1/create"

	var createRes kpay_response.CreateTransactionResponse
	_, err := callApi(kPayConfig, url, &request, &createRes)
	if err != nil {
		return kpay_response.CreateTransactionResponse{}, err
	}

	return createRes, nil
}

func CancelTransaction(kPayConfig *kpay_config.KPayConfig, request kpay_request.CancelTransactionRequest) (kpay_response.CancelTransactionResponse, error) {
	url := kPayConfig.KPayHost + "/api/payment/v1/cancel"

	var cancelRes kpay_response.CancelTransactionResponse
	_, err := callApi(kPayConfig, url, &request, &cancelRes)
	if err != nil {
		return kpay_response.CancelTransactionResponse{}, err
	}

	return cancelRes, nil
}

func QueryTransaction(kPayConfig *kpay_config.KPayConfig, request kpay_request.QueryTransactionRequest) (kpay_response.QueryTransactionResponse, error) {
	url := kPayConfig.KPayHost + "/api/payment/v1/check"

	var queryRes kpay_response.QueryTransactionResponse
	_, err := callApi(kPayConfig, url, &request, &queryRes)
	if err != nil {
		return kpay_response.QueryTransactionResponse{}, err
	}

	return queryRes, nil
}

func CreateVirtualAccount(kPayConfig *kpay_config.KPayConfig, request kpay_request.CreateVirtualAccountRequest) (kpay_response.CreateVirtualAccountResponse, error) {
	url := kPayConfig.KPayHost + "/api/payment/v1/virtualAccount/enable"

	var queryRes kpay_response.CreateVirtualAccountResponse
	_, err := callApi(kPayConfig, url, &request, &queryRes)
	if err != nil {
		return kpay_response.CreateVirtualAccountResponse{}, err
	}

	return queryRes, nil
}

func DisableVirtualAccount(kPayConfig *kpay_config.KPayConfig, request kpay_request.DisableVirtualAccountRequest) (kpay_response.DisableVirtualAccountResponse, error) {
	url := kPayConfig.KPayHost + "/api/payment/v1/virtualAccount/disable"

	var queryRes kpay_response.DisableVirtualAccountResponse
	_, err := callApi(kPayConfig, url, &request, &queryRes)
	if err != nil {
		return kpay_response.DisableVirtualAccountResponse{}, err
	}

	return queryRes, nil
}

func GetTransaction(kPayConfig *kpay_config.KPayConfig, request kpay_request.GetTransactionRequest) (kpay_response.PaginateResponse[kpay_response.PayTransactionResponse], error) {
	url := kPayConfig.KPayHost + "/api/payment/v1/getTransaction"

	var queryRes kpay_response.PaginateResponse[kpay_response.PayTransactionResponse]
	_, err := callApi(kPayConfig, url, &request, &queryRes)
	if err != nil {
		return kpay_response.PaginateResponse[kpay_response.PayTransactionResponse]{}, err
	}

	return queryRes, nil
}

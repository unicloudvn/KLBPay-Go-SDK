package webhook

import (
	"encoding/json"
	"errors"
	kpay_config "github.com/unicloudvn/KLBPay-Go-SDK/config"
	kpay_security "github.com/unicloudvn/KLBPay-Go-SDK/security"
	kpay_model "github.com/unicloudvn/KLBPay-Go-SDK/transaction/model"
	kpay_request "github.com/unicloudvn/KLBPay-Go-SDK/transaction/request"
	kpay_response "github.com/unicloudvn/KLBPay-Go-SDK/transaction/response"
	"net/http"
	"strconv"
)

type NotifyTransactionController struct {
	HandleNotify HandleNotify
	KPayConfig   *kpay_config.KPayConfig
}

func (notifyController *NotifyTransactionController) NotifyTransactionAPI(w http.ResponseWriter, r *http.Request) {
	// Get headers
	clientID := r.Header.Get("x-api-client")
	signature := r.Header.Get("x-api-validate")
	timestampStr := r.Header.Get("x-api-time")
	timestamp, err := parseTimestamp(timestampStr)
	if err != nil {
		// Handle timestamp parsing error
		http.Error(w, "Invalid timestamp", http.StatusBadRequest)
		return
	}
	// Parse request body
	var request kpay_request.BodyEncryptRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// Handle request body parsing error
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Process the request and generate response
	response, err := notifyController.processNotifyTransactionRequest(clientID, signature, timestamp, request)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize response to JSON
	respJSON, err := json.Marshal(response)
	if err != nil {
		// Handle JSON serialization error
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	// Set response headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func (notifyController *NotifyTransactionController) processNotifyTransactionRequest(clientID string, signature string, timestamp int64, request kpay_request.BodyEncryptRequest) (*kpay_response.KPayBaseResponse[kpay_response.NotifyResponse], error) {
	message := kpay_model.Message{
		ClientId:     clientID,
		ValidateData: signature,
		Timestamp:    timestamp,
		EncryptData:  request.Data,
	}
	var requestRaw kpay_request.NotifyRequest
	_, err := kpay_security.Decode(notifyController.KPayConfig, &message, &requestRaw)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	// Implement your business logic here
	// ...
	notifyController.HandleNotify.Handle(requestRaw)

	// Return a NotifyTransactionResponse wrapped in ResponseBase
	response := kpay_response.NotifyResponse{
		Success: true,
	}
	return &kpay_response.KPayBaseResponse[kpay_response.NotifyResponse]{Data: response}, nil
}

func parseTimestamp(timestampStr string) (int64, error) {
	// Implement parsing logic for timestamp string
	// For simplicity, let's assume timestampStr is a Unix timestamp in seconds
	timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return timestampInt, nil
}

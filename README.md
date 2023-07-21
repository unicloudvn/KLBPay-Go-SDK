# KPAY SDK GOLANG

Thư viện người dùng tích hợp KLBPay vào hệ thống thanh toán của Merchant

## Yêu cầu: GOLANG version >= 1.18

# Install

```
go get github.com/unicloudvn/KLBPay-Go-SDK@latest
```

# Usage

<h3 id="get-key">Lấy thông tin tích hợp từ: <a href = "https://mc.kienlongbank.com/">Klb Pay Portal</a> </h3>

```
Thông tin tích hợp bao gồm: Host, ClientId, SecretKey, EncryptKey
```

## Import sử dụng kpay-sdk

```
import (
	kpay_config "github.com/unicloudvn/KLBPay-Go-SDK/config"
	kpay_service "github.com/unicloudvn/KLBPay-Go-SDK/service"
	kpay_request "github.com/unicloudvn/KLBPay-Go-SDK/transaction/request"
)
```

## Ví dụ cơ bản

```
// Config với các options bên dưới
config := kpay_config.NewKPayConfig(
    kpay_config.WithEncryptKey(EncryptKey),
    kpay_config.WithSecretKey(SecretKey),
    kpay_config.WithClientId(ClientId),
    kpay_config.WithKPayHost(Host),
    kpay_config.WithMaxTimeStampDiff(1800),
)
```

### Tạo giao dịch

#### Request data

```
request := kpay_request.CreateTransactionRequest{
    RefTransactionId: "123456",
    Amount:           123,
    Description:      "Mo ta thanh toan",
    Timeout:          10000,
    Title:            "Thanh Toan",
    Language:         "Viet Nam",
    CustomerInfo: kpay_request.Customer{
        Fullname: "Nguyen Van A",
        Email:    "email@gmail.com",
        Phone:    "0123456789",
        Address:  "Ho Chi Minh",
    },
    SuccessUrl:    "https://success.example.com.vn",
    FailUrl:       "https://fail.example.com.vn",
    RedirectAfter: 5,
    BankAccountNo: "",
}
response, error := kpay_service.CreateTransaction(config, request)
```

#### Response data

```
{
  "accountName": "Cong ty 1 thanh vien ne",
  "amount": 123,
  "description": "Mo ta thanh toan",
  "payLinkCode": "8XBPd4Gp",
  "qrCodeString": "00020101021238610010A000000727013100069704520117106423071940959230208QRIBFTTA530370454031235802VN62220818TT Don hang 12345663044D30",
  "refTransactionId": "123456",
  "status": "CREATED",
  "time": "2023-07-19T16:17:24.650314",
  "timeout": 10000,
  "transactionId": "fbc87f79-92ec-4fdf-abf4-f78d8513c3b3",
  "url": "https://mc-staging.kienlongbank.co/public/paylink/8XBPd4Gp",
  "virtualAccount": "10642307194095923"
}
```

### Kiểm tra trạng thái giao dịch

#### Request data

```
request := kpay_request.QueryTransactionRequest{
    TransactionId: "fbc87f79-92ec-4fdf-abf4-f78d8513c3b3",
}
response, error := kpay_service.QueryTransaction(config, request)
```

#### Response data

```
{
  "amount": 123,
  "refTransactionId": "123456",
  "status": "CREATED"
}
```

### Huỷ giao dịch

#### Request data

```
request := kpay_request.CancelTransactionRequest{
    TransactionId: "fbc87f79-92ec-4fdf-abf4-f78d8513c3b3",
}
response, error := kpay_service.CancelTransaction(config, request)
```

#### Response data

```
{
  "success": true
}
```

### Tạo tài khoản ảo

#### Request data

```
request := kpay_request.CreateVirtualAccountRequest{
  Order:         888,
  Timeout:       300,
  FixAmount:     10000,
  FixContent:    "Content",
  BankAccountNo: "",
}
response, err := kpay_service.CreateVirtualAccount(config, request)
```

#### Response data

```
{
  "bankAccountNo": "4570834602",
  "fixAmount": 10000,
  "fixContent": "Content",
  "order": 888,
  "qrContent": "00020101021238620010A0000007270132000697045201181093992827920088820208QRIBFTTA53037045405100005802VN62110807Content6304D312",
  "timeout": 300,
  "virtualAccount": "109399282792008882"
}
```

### Hủy tài khoản ảo

#### Request data

```
request := kpay_request.DisableVirtualAccountRequest{
  Order: 888,
}
response, err := kpay_service.DisableVirtualAccount(config, request)
```

#### Response data

```
{
  "success": true
}
```

### Lấy lịch sử giao dịch

#### Request data

```
order := int64(2)
request := kpay_request.GetTransactionRequest{
  Page:     0,
  Size:     10,
  Order:    &order,                // order = nil -> lấy toàn bộ giao dịch
  FromDate: "2023-07-19 00:00:00", // Format: "yyyy-MM-dd hh:mm:ss"
  ToDate:   "2023-07-20 23:00:00", // Format: "yyyy-MM-dd hh:mm:ss"
}
response, err := kpay_service.GetTransaction(config, request)
```

#### Response data

```
{
  "items": [
    {
      "id": "f22fcd1a-46fe-4f42-bc5f-8dc02915121e",
      "status": "SUCCESS",
      "amount": 100000,
      "refTransactionId": "",
      "createDateTime": "2023-07-19 17:30:20",
      "completeTime": "2023-07-19 17:30:20",
      "virtualAccount": "109399282792000020",
      "description": "[109399282792000020.4570834602] Payme",
      "paymentType": "VIET_QR",
      "txnNumber": "P00000000353",
      "accountName": "TRAN NGOC THANG",
      "accountNo": "4570834602",
      "interBankTrace": "057ZEXA2313500IW"
    },
    {
      "id": "1b15b159-e8a2-4c84-8c20-4e620377f171",
      "status": "SUCCESS",
      "amount": 1000000,
      "refTransactionId": "",
      "createDateTime": "2023-07-19 17:12:16",
      "completeTime": "2023-07-19 17:12:16",
      "virtualAccount": "109399282792000020",
      "description": "[109399282792000020.4570834602] Payme",
      "paymentType": "VIET_QR",
      "txnNumber": "P00000000351",
      "accountName": "TRAN NGOC THANG",
      "accountNo": "4570834602",
      "interBankTrace": "057ZEXA2313500IV"
    }
  ],
  "pageNumber": 0,
  "pageSize": 10,
  "totalPage": 1,
  "totalSize": 2
}
```

## Call Webhook

### Request Curl

```
curl --location --request GET '${YOUR_HOST}/notifyTransaction' \
--header 'x-api-client: 5a404192-045b-4f6b-863f-192a0b9a88b4' \
--header 'x-api-time: 1689925407435' \
--header 'x-api-validate: 9757bbdbe1dc61e74cd8269126af130de9448e379d6bd12aed4443fea46cc5a0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data" :"qGckAOEvH9IqzQwxODueOmsMtLNUpVXdu/hbg0oPcV6K8ORAN2U/bGPC0OfrlyvN8h4mlCAIW3E94hFdiXGVv4bWbb4siQ5JSKuxQFw9lR/iiixcJ//RCFiGUY1cvkVx4zbpfIGchKGqEeay5wJZwy6V8YC/tvYqBN2zM/upfPQAlapxbvSF2ytW9AufUOe+TjPwe4xjVXx9x6/Ji8aB5sUCOeTzJ+j5D2YPtM6n8f25iUNgYN6eUhRm620aXDt1CmIhJYBUwa07kQN0UsxzC4Nq6FJVlgOz9auvyJrgcz9ym6p9l8KlS/pSKaMyERKP7FB72EpHZBUgbPGtkC8nv7oiEpiMszEw9JSwukOBBCdtBnjT6Xvk4CJUd8lX+Eg5rVhqkEqSvaPuDpPEVHFB6GqilF2o4avm7OGaphGP7fWmN714rA0rZ8rNxzTPdiTk88Wu1DYeeWd1w8RkNf8GGH30hbkRpsZ22UqKGDb00VxDgdArEeslkQOdFUGP9vtsqu4p79ps2e8Q0SA0ng8w9A=="
}'
```

### 1. Implement interface HandleNotify:

```
type NotifyImpl struct {
	webhook.HandleNotify
}

func (notifyImpl *NotifyImpl) Handle(request kpay_request.NotifyRequest) error {
    // request đã được decode từ từ data encrypt
    // Implement your business logic here

    return nil
}
```

### 2. Override method NotifyTransactionAPI trong NotifyController:

```
type NotifyTransactionImpl struct {
	kpay_api.NotifyController
}

func (override *NotifyTransactionImpl) NotifyTransactionAPI(w http.ResponseWriter, r *http.Request){
    // xử lí business logic cho...
}
```

### 3. Ví dụ:

- ### Implement interface HandleNotify

```
func main() {
	clientId := "5a404192-045b-4f6b-863f-192a0b9a88b4"
	encryptKey := "022E0BF1076030B6E64127107CC1BA576F643E08D2617E93DEB1D13AB1443896"
	secretKey := "TZ4IYbmI6d/wvqMMCGV5u4CcE5+KDFJ+lTewp/LoRGDj94p6yFQ1k+/gMY1EXt9Tm80qRfu9BH+KCBoPwoIKs5WXKtsXpCa9vzmxcEyNW1qWwRpSWrNanOzpWRKzyZazWaziWu1A1TleeYYEUhhWZeqR42w2MX68Uu30n5tNGAwY5+Ctl1axCUzoNBC9czTXpzUNr/zIByW9RWW+BKJ6LNZv5V4HxzYesFmTTbcO9fgG5zTkTZ6s/H9n6NJ8UJFItOB5Rska7uYWmEDdkOx/5UGLrddjdHreYDIxUgTI5tHEwJM/1sXeM3N3rlqsJNBFqngA78PYypv6EUtNLFa9wQ=="

	config := kpay_config.NewKPayConfig(
		kpay_config.WithEncryptKey(encryptKey),
		kpay_config.WithSecretKey(secretKey),
		kpay_config.WithClientId(clientId),
	)

	handleNotify := NotifyImpl{}
	controller := webhook.NotifyController{
		HandleNotify: &handleNotify,
		KPayConfig:   config,
	}
	r := mux.NewRouter()
	r.HandleFunc("/test/notifyTransaction", controller.NotifyTransactionAPI)
	http.ListenAndServe(":5557", r)
}

type NotifyImpl struct {
	webhook.HandleNotify
}

func (notifyImpl *NotifyImpl) Handle(request kpay_request.NotifyRequest) error {
	// request đã được decode từ từ data encrypt
	// Implement your business logic here

	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	} else {
		fmt.Println(string(jsonData))
	}

	return nil
}

```

- ### Override method NotifyTransactionAPI trong NotifyController

```

func main() {
	clientId := "5a404192-045b-4f6b-863f-192a0b9a88b4"
	encryptKey := "022E0BF1076030B6E64127107CC1BA576F643E08D2617E93DEB1D13AB1443896"
	secretKey := "TZ4IYbmI6d/wvqMMCGV5u4CcE5+KDFJ+lTewp/LoRGDj94p6yFQ1k+/gMY1EXt9Tm80qRfu9BH+KCBoPwoIKs5WXKtsXpCa9vzmxcEyNW1qWwRpSWrNanOzpWRKzyZazWaziWu1A1TleeYYEUhhWZeqR42w2MX68Uu30n5tNGAwY5+Ctl1axCUzoNBC9czTXpzUNr/zIByW9RWW+BKJ6LNZv5V4HxzYesFmTTbcO9fgG5zTkTZ6s/H9n6NJ8UJFItOB5Rska7uYWmEDdkOx/5UGLrddjdHreYDIxUgTI5tHEwJM/1sXeM3N3rlqsJNBFqngA78PYypv6EUtNLFa9wQ=="

	config := kpay_config.NewKPayConfig(
		kpay_config.WithEncryptKey(encryptKey),
		kpay_config.WithSecretKey(secretKey),
		kpay_config.WithClientId(clientId),
	)

	var name = NewNotifyTransactionImpl(config)
	r := mux.NewRouter()
	r.HandleFunc("/test/notifyTransaction", name.NotifyTransactionAPI)
	http.ListenAndServe(":5557", r)
}

type NotifyTransactionImpl struct {
	kpay_api.NotifyController
}

func (override *NotifyTransactionImpl) NotifyTransactionAPI(w http.ResponseWriter, r *http.Request) {
	// Get headers
	clientID := r.Header.Get("x-api-client")
	signature := r.Header.Get("x-api-validate")
	timestampStr := r.Header.Get("x-api-time")
	timestamp, err := parseTimestamp(timestampStr)
	if err != nil {
		// Handle timestamp parsing error
		http.Error(w, kpay_exception.InvalidParam.Message, http.StatusBadRequest)
		return
	}
	// Decode data encrypt from request
	var request kpay_request.BodyEncryptRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// Handle request body parsing error
		http.Error(w, kpay_exception.InvalidParam.Message, http.StatusBadRequest)
		return
	}

	message := kpay_model.Message{
		ClientId:     clientID,
		ValidateData: signature,
		Timestamp:    timestamp,
		EncryptData:  request.Data,
	}
	var requestRaw kpay_request.NotifyRequest
	_, err = kpay_security.Decode(override.KPayConfig, &message, &requestRaw)

	if err != nil {
		http.Error(w, kpay_exception.InvalidParam.Message, http.StatusBadRequest)
		return
	}

	// xử lí business logic tại đây với dữ liệu đã được mã hoá
	//...
	//...
	fmt.Println(requestRaw)
	fmt.Println(clientID, signature, timestamp)

	// handle response
	// Process the request and generate response
	response := kpay_response.NotifyResponse{
		Success: true,
	}
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

func parseTimestamp(timestampStr string) (int64, error) {
	// Implement parsing logic for timestamp string
	// For simplicity, let's assume timestampStr is a Unix timestamp in seconds
	timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return timestampInt, nil
}
func NewNotifyTransactionImpl(config *kpay_config.KPayConfig) *NotifyTransactionImpl {
	return &NotifyTransactionImpl{
		kpay_api.NotifyController{
			KPayConfig: config,
		},
	}
}

```

## Response Code

<table>
    <thead>
    <tr>
        <th style={{width: '100px', textAlign: 'center'}}>Code</th>
        <th style={{width: '300px', textAlign: 'center'}}>Message</th>
        <th style={{width: '300px', textAlign: 'center'}}>Description</th>
    </tr>
    </thead>
    <tbody>
    <tr>
        <td style={{textAlign: 'center'}}>0</td>
        <td style={{textAlign: 'center'}}>SUCCESS</td>
        <td style={{textAlign: 'center'}}>Thành công</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1</td>
        <td style={{textAlign: 'center'}}>FAILED</td>
        <td style={{textAlign: 'center'}}>Thất bại</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>2</td>
        <td style={{textAlign: 'center'}}>INVALID_PARAM</td>
        <td style={{textAlign: 'center'}}>Tham số không hợp lệ</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1601</td>
        <td style={{textAlign: 'center'}}>PAYMENT_SECURITY_VIOLATION</td>
        <td style={{textAlign: 'center'}}>Vi phạm bảo mật</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1602</td>
        <td style={{textAlign: 'center'}}>PAYMENT_ORDER_COMPLETED</td>
        <td style={{textAlign: 'center'}}>Giao dịch đã được thanh toán</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1603</td>
        <td style={{textAlign: 'center'}}>PAYMENT_AMOUNT_INVALID</td>
        <td style={{textAlign: 'center'}}>Số tiền không hợp lệ</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1604</td>
        <td style={{textAlign: 'center'}}>PAYMENT_TRANSACTION_CANCELED</td>
        <td style={{textAlign: 'center'}}>Giao dịch đã bị huỷ</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1605</td>
        <td style={{textAlign: 'center'}}>PAYMENT_TRANSACTION_EXPIRED</td>
        <td style={{textAlign: 'center'}}>Giao dịch đã hết hạn</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1606</td>
        <td style={{textAlign: 'center'}}>PAYMENT_TRANSACTION_INVALID</td>
        <td style={{textAlign: 'center'}}>Giao dịch không hợp lệ</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1607</td>
        <td style={{textAlign: 'center'}}>PAYMENT_TRANSACTION_FAILED</td>
        <td style={{textAlign: 'center'}}>Giao dịch thất bại</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1608</td>
        <td style={{textAlign: 'center'}}>PAYMENT_SERVICE_UNAVAILABLE</td>
        <td style={{textAlign: 'center'}}>Dịch vụ không khả dụng</td>
    </tr>
    <tr>
        <td style={{textAlign: 'center'}}>1609</td>
        <td style={{textAlign: 'center'}}>PAYMENT_INVALID_CLIENT_ID</td>
        <td style={{textAlign: 'center'}}>Mã khách hàng không hợp lệ</td>
    </tr>
    </tbody>
</table>

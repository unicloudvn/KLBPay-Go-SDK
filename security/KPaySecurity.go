package kpay_security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	kpay_config "kpay_sdk/config"
	kpay_exception "kpay_sdk/exception"
	kpay_model "kpay_sdk/transaction/model"
)

func SignData(data string, clientId string, timestamp int64, secretKey string) string {
	// format data
	message := fmt.Sprintf("%s|%d|%s", clientId, timestamp, data)

	var messageByte = []byte(message)
	var key = []byte(secretKey)

	// create hmac-sha-256
	var hash = hmac.New(sha256.New, key)
	hash.Write(messageByte)
	sign := hash.Sum(nil)
	return hex.EncodeToString(sign)
}

func AesEncrypt(data string, encryptKey string) (string, error) {
	// convert encrypt key(hex string) to byte array
	decodedKey, err := hex.DecodeString(encryptKey)
	if err != nil {
		return "", errors.New(kpay_exception.DecodeKeyFailed.Message)
	}

	// create cipher block
	block, err := aes.NewCipher(decodedKey)
	if err != nil {
		return "", errors.New(kpay_exception.CreateCipherError.Message)
	}

	// init IV using 16 bytes(range 0 to 16) of key
	iv := decodedKey[:16]
	mode := cipher.NewCBCEncrypter(block, iv)

	// using padding pkcs5
	plaintext := []byte(data)
	plaintext = pkcs5Padding(plaintext, aes.BlockSize)

	// create buffer to save cipher text
	cipherText := make([]byte, len(plaintext))

	mode.CryptBlocks(cipherText, plaintext)
	result := base64.StdEncoding.EncodeToString(cipherText)
	return result, nil
}

func AesDecrypt(ciphertext string, key string) (string, error) {
	decodedKey, err := hex.DecodeString(key)
	if err != nil {
		return "", errors.New(kpay_exception.DecodeKeyFailed.Message)
	}

	block, err := aes.NewCipher(decodedKey)
	if err != nil {
		return "", errors.New(kpay_exception.CreateCipherError.Message)
	}

	iv := decodedKey[:16]
	mode := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))

	cipherDecode, _ := base64.StdEncoding.DecodeString(ciphertext)

	mode.CryptBlocks(plaintext, cipherDecode)

	plaintext = pkcs5UnPadding(plaintext)

	return string(plaintext), nil
}

func Encode(kPayConfig *kpay_config.KPayConfig, data any) (*kpay_model.Message, error) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	if data == nil {
		return nil, errors.New(kpay_exception.InvalidParam.Message)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(kpay_exception.InvalidParam.Message)
	}

	encryptData, err := AesEncrypt(string(jsonData), kPayConfig.EncryptKey)
	if err != nil {
		return &kpay_model.Message{}, err
	}
	validateData := SignData(encryptData, kPayConfig.ClientId, timestamp, kPayConfig.SecretKey)

	return &kpay_model.Message{
		ClientId:     kPayConfig.ClientId,
		Timestamp:    timestamp,
		EncryptData:  encryptData,
		ValidateData: validateData,
	}, nil
}

func Decode[T any](kPayConfig *kpay_config.KPayConfig, message *kpay_model.Message, s *T) (*T, error) {
	if message.ClientId != "" && message.ClientId == kPayConfig.ClientId {
		timestamp := time.Now().Unix()
		if timestamp-message.Timestamp > kPayConfig.MaxTimeStampDiff {
			return nil, errors.New(kpay_exception.TransactionExpired.Message)
		}
		if message.EncryptData == "" {
			return nil, errors.New(kpay_exception.InvalidParam.Message)
		}
		validateData := SignData(message.EncryptData, message.ClientId, message.Timestamp, kPayConfig.SecretKey)
		if !strings.EqualFold(strings.ToLower(validateData), strings.ToLower(message.ValidateData)) {
			return nil, errors.New(kpay_exception.SecurityViolation.Message)
		}
		decryptData, err := AesDecrypt(message.EncryptData, kPayConfig.EncryptKey)
		if err != nil {
			return nil, err
		}
		decryptData = removeNullChars(decryptData)
		err = json.Unmarshal([]byte(decryptData), &s)
		if err != nil {
			return nil, err
		}
		return s, nil
	}
	return nil, errors.New(kpay_exception.InvalidClientID.Message)
}
func pkcs5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

func pkcs5UnPadding(plaintext []byte) []byte {
	plaintext = []byte(removeNullChars(string(plaintext)))
	length := len(plaintext)
	unPadding := int(plaintext[length-1])
	return plaintext[:(length - unPadding)]
}

func removeNullChars(jsonString string) string {
	re := regexp.MustCompile(string([]byte{0}))
	return re.ReplaceAllString(jsonString, "")
}

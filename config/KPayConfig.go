package kpay_config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type KPayConfig struct {
	ClientId         string
	SecretKey        string
	EncryptKey       string
	MaxTimeStampDiff int64
	KPayHost         string
}

func (config KPayConfig) Init() KPayConfig {
	godotenv.Load()

	config.ClientId = os.Getenv("Client_Id")
	config.SecretKey = os.Getenv("Secret_Key")
	config.EncryptKey = os.Getenv("Encrypt_Key")
	config.MaxTimeStampDiff, _ = strconv.ParseInt(os.Getenv("Max_Timestamp_Diff"), 10, 64)
	config.KPayHost = os.Getenv("Kpay_Host")

	return config
}

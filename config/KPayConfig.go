package kpay_config

type KPayConfig struct {
	ClientId         string
	SecretKey        string
	EncryptKey       string
	MaxTimeStampDiff int64
	KPayHost         string
}

type KPayConfigOption func(*KPayConfig)

func NewKPayConfig(opts ...KPayConfigOption) *KPayConfig {
	// Here we can do any initialization for all options, then the provided parameters can overwrite them.
	uCtx := KPayConfig{
		MaxTimeStampDiff: 1800,
		KPayHost:         "https://api-staging.kienlongbank.co/pay",
	}
	for _, o := range opts {
		o(&uCtx)
	}
	return &uCtx
}

// With ClientId
func WithClientId(clientId string) KPayConfigOption { // HL
	return func(c *KPayConfig) {
		c.ClientId = clientId
	}
}

// With SecretKey
func WithSecretKey(secretKey string) KPayConfigOption { // HL
	return func(c *KPayConfig) {
		c.SecretKey = secretKey
	}
}

// With EncryptKey
func WithEncryptKey(encryptKey string) KPayConfigOption { // HL
	return func(c *KPayConfig) {
		c.EncryptKey = encryptKey
	}
}

// With MaxTimeStampDiff
func WithMaxTimeStampDiff(maxTimeStampDiff int64) KPayConfigOption { // HL
	return func(c *KPayConfig) {
		c.MaxTimeStampDiff = maxTimeStampDiff
	}
}

// With KPayHost
func WithKPayHost(kPayHost string) KPayConfigOption { // HL
	return func(c *KPayConfig) {
		c.KPayHost = kPayHost
	}
}

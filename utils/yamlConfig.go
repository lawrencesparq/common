package utils

import (
	"log"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Config represents configuration for the app.
type Config struct {

	// Port to run the http server.
	AppPort string `yaml:"appPort"`

	// server ip address.
	ServerAddress string `yaml:"serverAddress"`

	// address to consul address.
	ConsulAddress string `yaml:"consulAddress"`

	// address to jaeger.
	JaegerAddress string `yaml:"jaegerAddress"`

	// broker username.
	BrokerUser string `yaml:"brokerUser"`

	// broker password.
	BrokerPassword string `yaml:"brokerPassword"`

	// broker host address.
	BrokerHost string `yaml:"brokerHost"`

	// broker port number.
	BrokerPort string `yaml:"brokerPort"`

	// database user name.
	DBUser string `yaml:"dbUser"`

	// database password.
	DBPassword string `yaml:"dbPassword"`

	// database host address.
	DBHost string `yaml:"dbHost"`

	// database port number.
	DBPort string `yaml:"dbPort"`

	// database name.
	DBName string `yaml:"dbName"`

	// Jwt Secret.
	JwtSecret string `yaml:"jwtSecret"`

	// Jwt Exp.
	JwtExp string `yaml:"jwtExp"`

	// grant type.
	GrantType string `yaml:"grantType"`

	// client secret.
	ClientSecret string `yaml:"clientSecret"`

	// client ID.
	ClientID string `yaml:"clientID"`

	// Airtel Auth Url.
	AirtelAuthUrl string `yaml:"airtelAuthUrl"`

	// Airtel collection Url.
	AirtelCollectionUrl string `yaml:"airtelCollectionUrl"`

	// Airtel disbursement Url.
	AirtelDisbursementUrl string `yaml:"airtelDisbursementUrl"`

	// Airtel collection Status Url.
	AirtelCollectionStatusUrl string `yaml:"airtelCollectionStatusUrl"`

	// MTN Collection Auth Url.
	MtnCollectionAuthUrl string `yaml:"mtnCollectionAuthUrl"`

	// MTN Disbursement Auth Url.
	MtnDisbursementAuthUrl string `yaml:"mtnDisbursementAuthUrl"`

	// MTN collection Url.
	MtnCollectionUrl string `yaml:"mtnCollectionUrl"`

	// MTN disbursement Url.
	MtnDisbursementUrl string `yaml:"mtnDisbursementUrl"`

	// MTN collection Status Url.
	MtnCollectionStatusUrl string `yaml:"mtnCollectionStatusUrl"`

	// MTN disbursement status Url.
	MtnDisbursementStatusUrl string `yaml:"mtnDisbursementStatusUrl"`

	// coll Sub Key.
	CollSubKey string `yaml:"collSubKey"`

	// coll user.
	CollUser string `yaml:"collUser"`

	// coll password.
	CollPassword string `yaml:"collPassword"`

	// disb Sub Key.
	DisbSubKey string `yaml:"disbSubKey"`

	// disb user.
	DisbUser string `yaml:"disbUser"`

	// disb password.
	DisbPassword string `yaml:"disbPassword"`

	//sender email
	SenderEmail string `yaml:"senderEmail"`

	//smtp key
	SmtpKey string `yaml:"smtpKey"`

	//cipher key
	CipherKey string `yaml:"cipherKey"`

	//cipher iv
	CipherIv string `yaml:"cipherIv"`
}

// loadConfig loads app config from YAML file.
func LoadConfig(path []byte) *Config {
	// Create config structure
	config := &Config{}
	// Read the config file from the disk.

	// Convert the YAML config into a Go struct.
	err := yaml.Unmarshal(path, config)
	if err != nil {
		log.Printf("yaml.Unmarshal failed: %v", err)
		return &Config{}
	}

	return config
}

func YamlString(key, fallback string) string {
	if key != "" {
		return key
	}

	return fallback
}

func YamlInt(key string, fallback int64) int64 {
	if key != "" {
		i, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}

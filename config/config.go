package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort             string `mapstructure:"APP_PORT"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	FirebaseCred        string `mapstructure:"FIREBASE_CREDENTIALS_FILE"`
	ImageKitPublicKey   string `mapstructure:"IMAGEKIT_PUBLIC_KEY"`
	ImageKitPrivateKey  string `mapstructure:"IMAGEKIT_PRIVATE_KEY"`
	ImageKitUrlEndpoint string `mapstructure:"IMAGEKIT_URL_ENDPOINT"`
	MinioEndpoint       string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKey      string `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretKey      string `mapstructure:"MINIO_SECRET_KEY"`
	MinioBucketName     string `mapstructure:"MINIO_BUCKET_NAME"`
	MinioUseSSL         bool   `mapstructure:"MINIO_USE_SSL"`
	PaymentServerKey    string `mapstructure:"PAYMENT_SERVER_KEY"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, reading from environment variables")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

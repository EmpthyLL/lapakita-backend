package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	// App
	AppName        string `mapstructure:"APP_NAME"`
	AppEnv         string `mapstructure:"APP_ENV"`
	AppPort        string `mapstructure:"APP_PORT"`
	FrontendOrigin string `mapstructure:"FRONTEND_ORIGIN"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`

	// Database
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSslMode  string `mapstructure:"DB_SSLMODE"`
	DBTimeZone string `mapstructure:"DB_TIMEZONE"`

	// JWT
	JWTSecret string `mapstructure:"JWT_SECRET"`
	JWTExpiry string `mapstructure:"JWT_EXPIRY"`

	// Firebase
	FirebaseCred string `mapstructure:"FIREBASE_CREDENTIALS_FILE"`

	// ImageKit
	ImageKitPrivateKey  string `mapstructure:"IMAGEKIT_PRIVATE_KEY"`
	ImageKitUrlEndpoint string `mapstructure:"IMAGEKIT_URL_ENDPOINT"`

	// MinIO
	MinioEndpoint   string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKey  string `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretKey  string `mapstructure:"MINIO_SECRET_KEY"`
	MinioBucketName string `mapstructure:"MINIO_BUCKET_NAME"`
	MinioUseSSL     bool   `mapstructure:"MINIO_USE_SSL"`

	// Payment
	PaymentServerKey    string `mapstructure:"PAYMENT_SERVER_KEY"`
	PaymentClientKey    string `mapstructure:"PAYMENT_CLIENT_KEY"`
	PaymentIsProduction bool   `mapstructure:"PAYMENT_IS_PRODUCTION"`

	// Redis
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDB       int    `mapstructure:"REDIS_DB"`

	// Geoapify
	GeoapifyAPIKey string `mapstructure:"GEOAPIFY_API_KEY"`

	// SMTP Mailer
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPUser     string `mapstructure:"SMTP_USER"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
	SMTPSender   string `mapstructure:"SMTP_SENDER"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Default values
	viper.SetDefault("FRONTEND_ORIGIN", "*")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("SMTP_PORT", 587)
	viper.SetDefault("SMTP_HOST", "smtp.gmail.com")
	viper.SetDefault("SMTP_SENDER", "Lapakita Platform <no-reply@lapakita.id>")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, reading from environment variables")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

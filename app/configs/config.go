package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	APPName           string
	APPPort           string
	APPLogLevel       string
	APPBaseURL        string
	DBHost            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBPort            string
	DBSSLMode         string
	DBTimeZone        string
	DBLogLevel        string
	DBAutoMigrate     bool
	JWTSecret         string
	JWTRefreshSecret  string
	S3Endpoint        string
	S3AccessKeyID     string
	S3SecretAccessKey string
	S3BucketName      string
	S3UseSSL          string
	NotifBaseURL      string
	NotifSMTPHost     string
	NotifSMTPPort     int
	NotifSMTPUserName string
	NotifSMTPPassword string
	NotifSMTPFrom     string
	NotifSMTPFromName string
}

var AppConfig *Config

func InitConfig() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	notifSMTPPort, _ := strconv.Atoi(os.Getenv("NOTIF_SMTP_PORT"))
	// Initialize configuration
	AppConfig = &Config{
		APPName:           os.Getenv("APP_NAME"),
		APPPort:           os.Getenv("APP_PORT"),
		APPLogLevel:       os.Getenv("APP_LOG_LEVEL"),
		APPBaseURL:        os.Getenv("APP_BASE_URL"),
		DBHost:            os.Getenv("DB_HOST"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		DBPort:            os.Getenv("DB_PORT"),
		DBSSLMode:         os.Getenv("DB_SSLMODE"),
		DBTimeZone:        os.Getenv("DB_TIMEZONE"),
		DBLogLevel:        os.Getenv("DB_LOG_LEVEL"),
		DBAutoMigrate:     os.Getenv("DB_AUTO_MIGRATE") == "true",
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTRefreshSecret:  os.Getenv("JWT_REFRESH_SECRET"),
		S3Endpoint:        os.Getenv("S3_ENDPOINT"),
		S3AccessKeyID:     os.Getenv("S3_ACCESS_KEY"),
		S3SecretAccessKey: os.Getenv("S3_SECRET_KEY"),
		S3BucketName:      os.Getenv("S3_BUCKET_NAME"),
		S3UseSSL:          os.Getenv("S3_USE_SSL"),
		NotifBaseURL:      os.Getenv("NOTIF_BASE_URL"),
		NotifSMTPHost:     os.Getenv("NOTIF_SMTP_HOST"),
		NotifSMTPPort:     notifSMTPPort,
		NotifSMTPUserName: os.Getenv("NOTIF_SMTP_USERNAME"),
		NotifSMTPPassword: os.Getenv("NOTIF_SMTP_PASSWORD"),
		NotifSMTPFrom:     os.Getenv("NOTIF_SMTP_FROM"),
		NotifSMTPFromName: os.Getenv("NOTIF_SMTP_FROM_NAME"),
	}
}

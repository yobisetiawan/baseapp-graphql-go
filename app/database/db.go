package database

import (
	"baseapp/app/configs"
	"baseapp/app/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		configs.AppConfig.DBHost,
		configs.AppConfig.DBUser,
		configs.AppConfig.DBPassword,
		configs.AppConfig.DBName,
		configs.AppConfig.DBPort,
		configs.AppConfig.DBSSLMode,
		configs.AppConfig.DBTimeZone,
	)

	var err error

	logLevel := logger.Silent
	if configs.AppConfig.DBLogLevel == "INFO" {
		logLevel = logger.Info
	}
	if configs.AppConfig.DBLogLevel == "WARNING" {
		logLevel = logger.Warn
	}
	if configs.AppConfig.DBLogLevel == "ERROR" {
		logLevel = logger.Error
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel), // Enable detailed logs
	})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	err = DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		log.Fatal(err)
	}

	initMigrate()

}

func initMigrate() {
	if configs.AppConfig.DBAutoMigrate {
		DB.AutoMigrate(
			&models.Admin{},
			&models.User{},
			&models.Product{},
			&models.Session{},
			&models.Otp{})
	}
}

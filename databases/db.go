package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBDatabase string
	DBPort     string
	DBDriver   string
}

func Connect() {

	dbConfig := DBConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBDatabase: getEnv("DB_DATABASE", "ohas_store"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBDriver:   getEnv("DB_DRIVER", "mysql"),
	}

	if dbConfig.DBDriver == "mysql" {

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBDatabase)
		Instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	} else {

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBDatabase, dbConfig.DBPort)
		Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	}

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

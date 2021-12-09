package config

import (
	"fmt"
	"log"
	"os"

	"product-api/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Address		string
	User		string
	Password	string
	Name		string
}

func getEnv(key, fallback string) string {
	godotenv.Load()

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewConfig() DatabaseConfig {
	return DatabaseConfig{
		Address: getEnv("DB_ADDRESS", "localhost:3306"),
		User: getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "root1234"),
		Name: getEnv("DB_NAME", "db_product_api"),
	}
}

func InitDB() *gorm.DB {
	config := NewConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", 
		config.User, config.Password, config.Address, config.Name)
	
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	initMigration(db)

	return db
}

func initMigration(db *gorm.DB) {
	db.AutoMigrate(&domain.Product{})
}
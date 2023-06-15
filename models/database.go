package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

type dbParams struct {
	DbUser     string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
}

var DB *gorm.DB

func ConnectDB() {
	exec, _ := os.Executable()
	configFilePath := filepath.Join(filepath.Dir(exec), ".env")

	if err := godotenv.Load(configFilePath); err != nil {
		fmt.Println(".env file not found")
	}

	var DbParams = dbParams{
		DbName:     os.Getenv("DB_NAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbPort:     os.Getenv("DB_PORT"),
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
	}

	dbParamsString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		DbParams.DbHost,
		DbParams.DbPort,
		DbParams.DbUser,
		DbParams.DbName,
		DbParams.DbPassword)

	db, err := gorm.Open("postgres", dbParamsString)
	if err != nil {
		panic("Не удалось подключиться к базе данных: " + err.Error())
	}
	db.AutoMigrate(&Elm{})
	if err != nil {
		return
	}

	DB = db
}

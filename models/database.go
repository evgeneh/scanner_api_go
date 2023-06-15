package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type dbParams struct {
	DbUser     string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
}

func ConnectDB() *gorm.DB {

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
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Elm{})

	return db
}

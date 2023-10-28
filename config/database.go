package config

import (
	"fmt"

	helper "github.com/marcelomoresco/go-todo-crud/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host="localhost"
	port=5432
	user = "postgres"
	password ="postgres"
	dbName = "todo-go"
)

func DatabaseConfig() *gorm.DB{
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db

}
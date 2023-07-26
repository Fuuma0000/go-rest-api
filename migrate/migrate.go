package main

import (
	"fmt"

	"github.com/Fuuma0000/go-rest-api/db"
	"github.com/Fuuma0000/go-rest-api/model"
)

// GO_ENV=dev go run migrate/migrate.go
func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}

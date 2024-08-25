package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/tahahamdii/basic-api/cmd/api"
	"github.com/tahahamdii/basic-api/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:  "root",
		Passwd: "root",
		Addr: "localhost:3306",
		DBName: "basic_api",
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	server := api.NewApiServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/tahahamdii/basic-api/cmd/api"
	"github.com/tahahamdii/basic-api/config"
	"github.com/tahahamdii/basic-api/db"
)

func main() {
	cfg := mysql.Config{
	User:                 configs.Envs.DBUser,
	Passwd:               configs.Envs.DBPassword,
	Addr:                 configs.Envs.DBAddress,
	DBName:               configs.Envs.DBName,
	Net:                  "tcp",
	AllowNativePasswords: true,
	ParseTime:            true,
}

db, err := db.NewMySQLStorage(cfg)
if err != nil {
	log.Fatal(err)
}

initStorage(db)

server := api.NewApiServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
if err := server.Run(); err != nil {
	log.Fatal(err)
}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}

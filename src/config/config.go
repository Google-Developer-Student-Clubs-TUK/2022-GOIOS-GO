package config

import (
	"GOIOS/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(data.DnsVal), &gorm.Config{})
	if err != nil {
		panic("Failed to connect Postgres database")
	}
	log.Println("DB connection success!")
	return db
}

func DisconnectDB(db *gorm.DB) {
	dbVar, err := db.DB()
	if err != nil {
		panic("Fail to kill connection from databases")
	}
	dbVar.Close()
}

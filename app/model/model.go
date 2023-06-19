package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Scissor struct {
	ID			uint64	`json:"id" gorm:"primaryKey"`
	Redirect	string	`json:"redirect" gorm:"not null"`
	Scissor		string	`json:"scissor" gorm:"unique;not null"`
	Clicked		uint64	`json:"clicked"`
	Random		bool	`json:"random"`
}


func Setup() {

	dsn := "host=172.17.0.2 user=admin password=test dbname=admin port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Scissor{})
	if err != nil {
		fmt.Println(err)
	}
}

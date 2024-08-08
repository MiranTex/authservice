package main

import (
	"authservice/src/Database"
	"authservice/src/Router"

	"gorm.io/gorm"
)

// import "authservice/src/Router"

type User struct {
	gorm.Model
	username string
	password string
}

func main() {

	db := Database.Connect()

	_ = db.AutoMigrate(&User{})

	//test connection

	Router.StartRoute().Run(":8080")
}

package main

import (
	"exercise1/Config"
	"exercise1/Models"
	"exercise1/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	gorm2 "gorm.io/gorm"
)

var err error

func main(){
	Config.DB, err= gorm2.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())),&gorm2.Config{})
	if err!= nil{
		fmt.Println("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Customer{})
	Config.DB.AutoMigrate(&Models.Order{})

	r:= Routes.SetupRouter()

	r.Run()

}

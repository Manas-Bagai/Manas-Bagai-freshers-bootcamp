package main

import (
	"exercise2/Config"
	"exercise2/Models"
	"exercise2/Routes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main(){
	Config.DB, err= gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))
	if err!= nil{
		fmt.Println("Status:", err)

	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	r:= Routes.SetupRouter()

	r.Run()

}

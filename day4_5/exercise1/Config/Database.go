package Config

import (
	"fmt"
	gorm2 "gorm.io/gorm"
)

var DB *gorm2.DB

type DBConfig struct {
	Host string
	Port int
	User string
	DBName string
	Password string
}

func BuildDBConfig() *DBConfig{
	dbConfig:=DBConfig{
		Host: "localhost",
		Port: 3306,
		User: "root",
		DBName: "shop5",
		Password: "password",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig)string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
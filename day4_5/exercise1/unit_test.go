package main

import (
	"bytes"
	"exercise1/Config"
	"exercise1/Models"
	"exercise1/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	gorm2 "gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestGetProductsByID(t *testing.T){
	Config.DB, err= gorm2.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())),&gorm2.Config{})
	if err!= nil{
		fmt.Println("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Customer{})
	Config.DB.AutoMigrate(&Models.Order{})


	req, err:= http.NewRequest("GET","http://localhost:8080/shop-api/product/4",nil)
	if err!=nil{
		t.Fatal(err)
	}
	rr:=httptest.NewRecorder()
	Router:=Routes.SetupRouter()
	handler:= http.Handler(Router)
	handler.ServeHTTP(rr,req)
	if status:=rr.Code; status!= http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",status,http.StatusOK)
	}

	expected:= `{"id":4,"product_name":"CDE","price":30,"quantity":6,"message":"Product added successfully"}`
	if rr.Body.String()!=expected{
		t.Errorf("handler returned unexpected body:got %v want %v ", rr.Body.String(), expected)
	}
}

func TestCreateProducts(t *testing.T){
	Config.DB, err= gorm2.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())),&gorm2.Config{})
	if err!= nil{
		fmt.Println("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Customer{})
	Config.DB.AutoMigrate(&Models.Order{})

	var newentry=[]byte(`{"product_name":"ABC","price":20,"quantity":5}`)

	req,err:= http.NewRequest("POST","http://localhost:8080/shop-api/product",bytes.NewBuffer(newentry))
	if err!=nil{
		t.Errorf("Status: %v", err)
	}
	rr:=httptest.NewRecorder()
	Router:=Routes.SetupRouter()
	handler:=http.Handler(Router)
	handler.ServeHTTP(rr,req)

	if status:=rr.Code; status!= http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",status,http.StatusOK)
	}

	expected:= `{"id":7,"product_name":"ABC","price":20,"quantity":5,message:"Product added successfully"}`
	if rr.Body.String()!=expected{
		t.Errorf("handler returned unexpected body:got %v want %v ", rr.Body.String(), expected)
	}
}
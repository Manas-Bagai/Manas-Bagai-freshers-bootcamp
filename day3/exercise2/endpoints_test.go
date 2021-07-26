package main

import (
	"bytes"
	"exercise2/Config"
	"exercise2/Models"
	"exercise2/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestGetEntries(t *testing.T){
	Config.DB, err= gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))
	if err!= nil{
		fmt.Println("Status:", err)

	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	req, err:= http.NewRequest("GET","http://localhost:8080/user-api/user",nil)
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

	expected:= `[{"id":1,"name":"Manas","lastName":"manas@gmail.com","dob":"16011998","address":"New Road","subject":"Maths","marks":99}]`
	if rr.Body.String()!=expected{
		t.Errorf("handler returned unexpected body:got %v want %v ", rr.Body.String(), expected)
	}
}

func TestGetEntriesByID(t *testing.T){
	Config.DB, err= gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err!=nil{
		fmt.Println("Status: ",err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	req,err:= http.NewRequest("GET","http://localhost:8080/user-api/user/1",nil)
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

	expected:= `{"id":1,"name":"Manas","lastName":"manas@gmail.com","dob":"16011998","address":"New Road","subject":"Maths","marks":99}`
	if rr.Body.String()!=expected{
		t.Errorf("handler returned unexpected body:got %v want %v ", rr.Body.String(), expected)
	}

}

func TestCreateEntries(t *testing.T){
	Config.DB, err= gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err!=nil{
		fmt.Println("Status: ",err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	var newentry=[]byte(`{"id":12,"name":"ABC","lastName":"DEF","dob":"11112000","address":"New Road","subject":"Science","marks":99}`)

	req,err:= http.NewRequest("POST","http://localhost:8080/user-api/user",bytes.NewBuffer(newentry))
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

	expected:= `{"id":12,"name":"ABC","lastName":"DEF","dob":"11112000","address":"New Road","subject":"Science","marks":99}`
	if rr.Body.String()!=expected{
		t.Errorf("handler returned unexpected body:got %v want %v ", rr.Body.String(), expected)
	}
}

func TestUpdateEntries(t *testing.T){
	Config.DB, err= gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err!=nil{
		fmt.Println("Status: ",err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	var updatedentry=[]byte(`{"id":12,"name":"ABC","lastName":"DEF","dob":"11112000","address":"New Colony","subject":"Science","marks":78}`)

	req,err:= http.NewRequest("PUT","http://localhost:8080/user-api/user/12",bytes.NewBuffer(updatedentry))
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

	expected:= `{"id":12,"name":"ABC","lastName":"DEF","dob":"11112000","address":"New Colony","subject":"Science","marks":78}`
	if rr.Body.String()!=expected{
		t.Errorf("handler returned unexpected body:got %v want %v ", rr.Body.String(), expected)
	}
}

func TestDeleteEntries(t *testing.T){
	Config.DB, err= gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err!=nil{
		fmt.Println("Status: ",err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	req,err:= http.NewRequest("DELETE","http://localhost:8080/user-api/user/12",nil)
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

	expected:= `"id12":"is deleted"`
	if rr.Body.String()!=expected{
		t.Errorf("handler returned unexpected body:got %v want %v ", rr.Body.String(), expected)
	}
}
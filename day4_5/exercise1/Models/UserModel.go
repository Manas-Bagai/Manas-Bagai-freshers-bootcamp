package Models

import "time"

type Product struct {
	//gorm2.Model
	Id int `json:"id" gorm:"primary_key"`
	ProductName string `json:"product_name"`
	Price int `json:"price"`
	Quantity int `json:"quantity"`
	Message string `json:"message" gorm:"default:'Product added successfully'"`
}

type Customer struct{
	//gorm2.Model
	Id int `json:"id" gorm:"primary_key"`
	CustomerName string `json:"customer_name"`
	CustomerAddress string `json:"customer_address"`
	CustomerNumber string `json:"customer_number"`
	Time time.Time `gorm:"default:current_timestamp"`
}

type Order struct {
	//gorm2.Model
	Id int `json:"id" gorm:"primary_key" `
	CustomerId int `json:"customer_id"`
	Customer Customer `json:"customer" gorm:"foreignKey:CustomerId ; constraint:OnUpdate:CASCADE,OnDELETE:CASCADE;"`
	ProductId int `json:"product_id"`
	Product Product `json:"product" gorm:"foreignKey:ProductId; constraint:OnUpdate:CASCADE,OnDELETE:CASCADE;"`
	Quantity int `json:"quantity"`
	Status string `json:"status" gorm:"default:'Order Successful'"`
}

func (product Product) TableName1() string{
	return "products"
}

func (customer Customer) TableName2() string{
	return "customers"
}

func (order Order) TableName3() string{
	return "orders"
}
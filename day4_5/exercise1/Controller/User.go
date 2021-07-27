package Controller

import (
	"exercise1/Models"
	"exercise1/Service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProducts(c *gin.Context){
	var products []Models.Product
	err:= Service.GetProducts(&products)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,products)
	}
}

func CreateProduct(c *gin.Context){
	var product Models.Product
	err:= Service.CreateProduct(c,&product)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,product)
	}
}


func GetProductByID(c *gin.Context){
	var product Models.Product
	err:= Service.GetProductByID(c,&product)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,product)
	}
}

func PatchProduct(c *gin.Context){
	var product Models.Product
	err:= Service.PatchProduct(c,&product)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,product)
	}
}

func DeleteProduct(c *gin.Context){
	var product Models.Product
	err:= Service.DeleteProduct(c,&product)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,product)
	}
}


func CreateCustomer(c *gin.Context){
	var customer Models.Customer
	err:= Service.CreateCustomer(c,&customer)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,customer)
	}
}

func GetCustomerByID(c *gin.Context){
	var customer Models.Customer
	err:= Service.GetCustomerByID(c,&customer)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,customer)
	}
}

func PlaceOrder(c *gin.Context){
	var order Models.Order
	err:=Service.PlaceOrder(c,&order)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,order)
	}
}

func GetOrders(c *gin.Context){
	var orders []Models.Order
	err:= Service.GetOrders(&orders)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,orders)
	}
}

func GetOrderByID(c *gin.Context){
	var order Models.Order
	err:= Service.GetOrderByID(c,&order)
	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,order)
	}
}
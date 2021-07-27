package Service

import (
	"errors"
	"exercise1/Config"
	"exercise1/Models"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func PlaceOrder(c *gin.Context, order *Models.Order)error{
	var product1 Models.Product
	var customer Models.Customer

	c.BindJSON(order)

	quant:=order.Quantity
	if err := Config.DB.Where("id = ?", order.ProductId).First(&product1).Error; err != nil {
		return err
	}
	stock := product1.Quantity
	if quant>stock{
		order.Status="Out of stock"
		return errors.New("out of stock")
	}

	product1.Quantity-=quant

	Config.DB.Save(product1)

	customerId :=order.CustomerId
	if err := Config.DB.Where("id = ?", customerId).Last(&customer).Error; err != nil {
		order.Status="Customer Not Found"
		return err
	}
	t1:=customer.Time.Add(200*time.Second)
	if time.Now().Before(t1){
		order.Status="Cool Down"
		return errors.New("Cool Down")
	}else{
		customer.Time=time.Now()
	}
	mu:=sync.Mutex{}
	mu.Lock()
	if err := Config.DB.Create(order).Error; err != nil {
		return err
	}
	mu.Unlock()
	return nil


}

func GetOrders(orders *[]Models.Order)error{
	mu:=sync.Mutex{}
	mu.Lock()
	if err := Config.DB.Find(orders).Error;err!=nil{
		return err
	}
	mu.Unlock()
	return nil
}

func GetOrderByID(c *gin.Context, order *Models.Order)error{
	mu:=sync.Mutex{}
	mu.Lock()
	id:=c.Params.ByName("id")
	if err := Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	mu.Unlock()
	return nil
}

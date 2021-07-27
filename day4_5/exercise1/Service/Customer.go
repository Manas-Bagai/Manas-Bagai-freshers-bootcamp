package Service

import (
	"exercise1/Config"
	"exercise1/Models"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context, customer *Models.Customer)error{
	c.BindJSON(customer)
	if err := Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerByID(c *gin.Context, customer *Models.Customer)error{
	id:=c.Params.ByName("id")
	if err := Config.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}
package Routes

import (
	"exercise1/Controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r:= gin.Default()
	grp1:=r.Group("/shop-api")
	{
		grp1.POST("product",Controller.CreateProduct)
		grp1.PATCH("product/:id", Controller.PatchProduct)
		grp1.GET("product/:id", Controller.GetProductByID)
		grp1.GET("product", Controller.GetProducts)
		grp1.POST("order", Controller.PlaceOrder)
		grp1.GET("order/:id", Controller.GetOrderByID)
		grp1.GET("order",Controller.GetOrders)
		grp1.DELETE("product/:id", Controller.DeleteProduct)
		grp1.POST("customer", Controller.CreateCustomer)
		grp1.GET("customer/:id", Controller.GetCustomerByID)
	}
	return r
}

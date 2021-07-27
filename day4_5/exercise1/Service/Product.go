package Service

import (
	"exercise1/Config"
	"exercise1/Models"
	"github.com/gin-gonic/gin"
	"sync"
)

func CreateProduct(c *gin.Context,product *Models.Product) (err error) {
	c.BindJSON(product)
	mu:=sync.Mutex{}
	mu.Lock()
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	mu.Unlock()
	return nil
}

func GetProducts(products *[]Models.Product)(err error){
	if err = Config.DB.Find(products).Error;err!=nil{
		return err
	}
	return nil
}

func GetProductByID(c *gin.Context,product *Models.Product) (err error) {
	id:=c.Params.ByName("id")
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func PatchProduct(c *gin.Context,product *Models.Product) (err error) {
	err1:=GetProductByID(c,product)
	if err1!=nil{
		return err
	}
	c.BindJSON(product)
	if err = Config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(c *gin.Context, product *Models.Product)error{
	id:=c.Params.ByName("products.id")
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}
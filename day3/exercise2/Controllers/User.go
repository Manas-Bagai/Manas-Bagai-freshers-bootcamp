package Controllers

import (
	"exercise2/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudents(c *gin.Context){
	var user []Models.User
	err:= Models.GetAllStudents(&user)
	if err!=nil{
		fmt.Println(err.Error())
	} else{
		c.JSON(http.StatusOK,user)
	}
}

func CreateStudent(c *gin.Context){
	var user Models.User
	c.BindJSON(&user)
	err:= Models.CreateStudent(&user)
	if err!=nil{
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK,user)
	}
}

func GetStudentByID(c *gin.Context){
	id:= c.Params.ByName("id")
	var user Models.User
	err:= Models.GetStudentByID(&user,id)
	if err!= nil{
		fmt.Println(err.Error())
	}else {
		c.JSON(http.StatusOK,user)
	}
}

func UpdateStudent(c *gin.Context){
	id:= c.Params.ByName("id")
	var user Models.User
	err:=Models.GetStudentByID(&user,id)
	if err!=nil{
		fmt.Println(err.Error())
	}
	c.BindJSON(&user)
	err= Models.UpdateStudent(&user,id)
	if err!= nil{
		fmt.Println(err.Error())
	}else {
		c.JSON(http.StatusOK,user)
	}
}

func DeleteStudent(c *gin.Context){
	id:= c.Params.ByName("id")
	var user Models.User
	err:= Models.DeleteStudent(&user,id)
	if err!= nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK,gin.H{"id"+id:"is deleted"})
	}
}
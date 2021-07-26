package Models

import (
	"exercise2/Config"
	"fmt"
)

func GetAllStudents(user *[]User)(err error){
	if err = Config.DB.Find(user).Error;err!=nil{
		return err
	}
	return nil
}

func CreateStudent(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetStudentByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStudent(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

func DeleteStudent(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
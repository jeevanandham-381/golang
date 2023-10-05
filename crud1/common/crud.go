package common

import (
	"gorm.io/gorm"
	"mymodule/models"

	
)

func CreateProduct(db *gorm.DB,product *models.Employee)error{
	if err:=db.Create(product).Error;err!=nil{
		return err
	}
	return nil
}

func GetProduct(db *gorm.DB,productID uint)(*models.Employee,error){
	var product models.Employee
	if err:=db.First(&product,productID).Error;err!=nil{
		return nil,err

	}
	return &product,nil
}


func GetAll(db *gorm.DB)([]models.Employee,error){
	var products []models.Employee
	if err:=db.Find(&products).Error; err!=nil{
		return nil,err
	}
	return products,nil

	}

func UpdateData(db *gorm.DB,product *models.Employee)error{
	if err:=db.Save(product).Error;err!=nil{
		return err
	}
	return nil
}

func DeleteData(db *gorm.DB,product *models.Employee)error{
	if err:=db.Delete(product).Error;err!=nil{
		return err

	}
	return nil

}


package common

import (
	"gorm.io/gorm"
    "gorm.io/driver/postgres"
)

var DB *gorm.DB

func InitDB()(*gorm.DB,error){
	aut:="user=postgres password=root dbname=demo sslmode=disable"
	db,err:=gorm.Open(postgres.Open(aut),&gorm.Config{})
	if err!=nil{
		return nil,err
	}
	DB=db
	return db,nil
	

}
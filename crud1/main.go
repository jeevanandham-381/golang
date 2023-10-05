package main

import (
	"fmt"
	"mymodule/common"
	"mymodule/models"
)

func main(){
	fmt.Println("1.create 2.read one 3.getall 4.update 5.delete")
	var choice int
	fmt.Scan(&choice)

	db,err:=common.InitDB()
	if err!=nil{
		panic("not connect"+ err.Error())
	}
	db=db.Table("jerry")
	
	if choice==1{
	if err:=db.AutoMigrate(&models.Employee{});err!=nil{
		panic("fail to migrate")
	}
	Emp1:=&models.Employee{Name: "vinoth",City: "salem"}
	err=common.CreateProduct(db,Emp1)
	if err!=nil{
		fmt.Printf("%v",err)
	}

}else if choice==2{
	var id uint
	fmt.Println("enter the id you want see the details")
	fmt.Scan(&id)
	prouctID:=uint(id)
	retriev,err:=common.GetProduct(db,prouctID)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(retriev.ID,retriev.Name,retriev.City)
	}

}else if choice==3{
	products,err:=common.GetAll(db)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("all employe")
		for _,emp:=range products{
			fmt.Println(emp.ID,emp.Name,emp.City)
		}
	}
}else if choice==4{
	productID:=uint(1)
	updateProduct:=&models.Employee{
		ID: productID,
		Name: "kalai",
		City: "salem",
	}
	err:=common.UpdateData(db,updateProduct)
	if err!=nil{
		panic("not update"+err.Error())

	}else{
		fmt.Println("updated")
	}
}else if choice==5{
	var id uint
	fmt.Println("enter the id you want to delete")
	fmt.Scan(&id)
	prouctID:=models.Employee{ID:id}
	err:=common.DeleteData(db,&prouctID)
	if err!=nil{
		panic("not delete")
	}

}



}
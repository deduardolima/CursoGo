package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct{
	ID int `gorm:"primaryKey"`
	Name string
	
}

type Product struct{
	ID    int`gorm:"primaryKey"`
	Name 	string
	Price float64
	CategoryID int
	Category Category
	gorm.Model

}

func main(){
	dsn :="root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)

	}
	db.AutoMigrate(&Product{}, &Category{})

	// create batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Keyboard", Price: 100.00},
	// }
	// db.Create(products)
	
	// select one
		// var product Product
		// db.First(&product, 3)
		// fmt.Println(product);
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product);

	// Select All
	
	// var products []Product
	// db.Find(&products)
	// for _, product := range products{
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("name LIKE ?", "%k%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name ="New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2,1)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)

	// category := Category{Name:"Eletronicos"}
	// db.Create(&category)

	// // create product

	// db.Create(&Product{
	// 	Name: "Notebook",
	// 	Price: 1000.00,
	// 	CategoryID: category.ID,
	// })

	var products []Product
	db.Preload("Category").Find(&products)
	 for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	 }
	
}
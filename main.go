package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// for api
// func main() {
// 	r := gin.Default()
// 	v1 := r.Group("/v1")
// 	{
// 		v1.GET("/ping", api.Pong)
// 		v1.GET("/user/:name", api.Name)
// 		v1.GET("/users", api.Names)
// 	}
// 	v2 := r.Group("/v2")
// 	v2.POST("/users", api.AddUser)
// 	r.Run()
// }

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migration the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ? ", "D42")

	// Update
	db.Model(&product).Update("Price", 200)
	// Update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//Delete
	// db.Delete(&product,1)

}

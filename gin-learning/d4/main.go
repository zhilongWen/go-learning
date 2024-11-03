package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type HelloWorld struct {
	gorm.Model
	Name     string
	Age      int
	Sex      string
	Birthday string
}

// https://jasperxu.github.io/gorm-zh/database.html#m
func main() {

	db, err := gorm.Open("mysql", "root:root123@tcp(10.211.55.102:3306)/go_mall_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Create
	//for i := 0; i < 3; i++ {
	//	db.Create(&HelloWorld{
	//		Name:     fmt.Sprintf("user%d", i),
	//		Age:      i,
	//		Sex:      "male",
	//		Birthday: time.Now().Format("2006-01-02"),
	//	})
	//}

	// Read
	//var helloWorlds []HelloWorld
	////var helloWorlds HelloWorld
	////db.Find(&helloWorlds).Where("name = ?", "user2")
	//db.Where("name = ?", "user2").Find(&helloWorlds)
	//for idx, v := range helloWorlds {
	//	println(fmt.Sprintf("idx:%d, name:%s, age:%d, sex:%s, birthday:%s", idx, v.Name, v.Age, v.Sex, v.Birthday))
	//}
	////fmt.Println(helloWorlds)

	// Update
	//var helloWorld HelloWorld
	////db.Where("name = ?", "user2").First(&helloWorld).Update("age", 20)
	//db.Where("name = ?", "user2").First(&helloWorld).Updates(HelloWorld{Age: 11, Sex: "female"})

	// Delete
	var helloWorld HelloWorld
	// Dialect() 硬删除
	db.Where("id in (?)", []int{1, 2}).First(&helloWorld).Delete(&HelloWorld{})

	db.AutoMigrate(&HelloWorld{})

}

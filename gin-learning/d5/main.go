package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:user_name;type:varchar(255)" `
	Age  int
}

func (u *User) TableName() string {
	return "test_user"
}

// https://jasperxu.github.io/gorm-zh/associations.html
func main() {

	db, err := gorm.Open("mysql", "root:root123@tcp(10.211.55.102:3306)/go_mall_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//db.AutoMigrate(&User{})
	db.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})

	//idCard := IDCard{
	//	Number: 123456,
	//}
	//student := Student{
	//	StudentName: "student1",
	//	IDCard:      idCard,
	//}
	//
	//teacher := Teacher{
	//	TeacherName: "teacher1",
	//	Students:    []Student{student},
	//}
	//
	//class := Class{
	//	ClassName: "class1",
	//	Students:  []Student{student},
	//}
	//
	//db.Create(&class)
	//db.Create(&student)
	//db.Create(&teacher)
	//db.Create(&idCard)

	r := gin.Default()

	r.POST("/create_class", func(c *gin.Context) {

		var student Student
		_ = c.BindJSON(&student)

		db.Create(&student)

	})

	r.Run(":8080")
}

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model

	StudentName string
	CLassID     uint

	IDCard IDCard

	Teachers []Teacher `gorm:"many2many:student_teachers;"`
}

type IDCard struct {
	gorm.Model
	StudentID uint

	Number int
}

type Teacher struct {
	gorm.Model
	TeacherName string
	Classes     []Class `gorm:"many2many:teacher_classes;"`
	Students    []Student
}

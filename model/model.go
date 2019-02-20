package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"html/template"
	"log"
)

var DB *gorm.DB

func init() {

	temp, err := gorm.Open("mysql", "root:Gostanford1@/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal()
	}

	DB = temp.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	SetTables()
	fmt.Println("DONE SETTING TABLES")
	SetRelationship()
	fmt.Println("DONE SETTING RELATIONSHIPS")
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Post{})
	DB.AutoMigrate(&Project{})

	//Populate Tables
	//Populate(DB)
}

type User struct {
	ID       uint `gorm:"primary_key"`
	Password string
	Name     string
	Email    string
}

//Blog post
type Post struct {
	gorm.Model
	Title   string
	Body    template.HTML `sql:"type:longtext"`
	Summary template.HTML
	User    User //`gorm:"foreignkey:UserID"`
	UserID  uint
}

type Project struct {
	gorm.Model
	Title   string
	Body    template.HTML `sql:"type:longtext"`
	Summary template.HTML
	User    User
	UserID  uint
}

type Note struct {
	gorm.Model
	Body   template.HTML
	User   User
	UserID uint
}

/*
type Portfolio struct {


}
*/

func SetTables() {

	DB.CreateTable(&User{})
	DB.CreateTable(&Post{})
	DB.CreateTable(&Project{})
	DB.CreateTable(&Note{})

}

//Set Relationships
func SetRelationship() {

	var user User
	var post Post
	var project Project
	var note Note

	DB.Model(&user).Related(&post)
	DB.Model(&user).Related(&project)
	DB.Model(&user).Related(&note)

}

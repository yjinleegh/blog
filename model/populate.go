package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Function to populate the database when the database is empty
func Populate(db *gorm.DB) {
	Users(db)
	Posts(db)
	Projects(db)
	Notes(db)
	Tags(db)
	Categories(db)
}

//Populating users table with sample users
func Users(db *gorm.DB) {
	fmt.Println("SETTING USERS")
	examples := []User{
		User{Name: "Andy", Email: "Andy@test.com", Password: "1234"},
		User{Name: "Brian", Email: "Brian@test.com", Password: "1234"},
		User{Name: "Joe", Email: "Joe@test.com", Password: "1234"},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}

//Populating posts table with sample posts
func Posts(db *gorm.DB) {

	fmt.Println("SETTING POSTS")
	examples := []Post{
		Post{Title: "LOL", Body: "LOLBA", Summary: "Hi", Tags: []Tag{Tag{Name: "First Tag"}, Tag{Name: "Second Tag"}}, UserID: 1},

		Post{Title: "LOL", Body: "LILBA", Summary: "Hi", Tags: []Tag{Tag{Name: "Second Tag"}, Tag{Name: "Third Tag"}}, UserID: 1},

		Post{Title: "LOL", Body: "LULBA", Summary: "Hi", Tags: []Tag{Tag{Name: "Third Tag"}, Tag{Name: "Fourth Tag"}}, UserID: 1},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}

//Populating projects table with sample projects
func Projects(db *gorm.DB) {

	fmt.Println("SETTING PROJECTS")
	examples := []Project{
		Project{Title: "Jeff", Body: "LOLBA", UserID: 1},
		Project{Title: "Daniel", Body: "LULBA", UserID: 2},
		Project{Title: "Wanggui", Body: "LILBA", UserID: 3},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}

//Populating tags table with sample tags
func Tags(db *gorm.DB) {
	fmt.Println("SETTING TAGS")
	examples := []Tag{
		Tag{Name: "Fifth Tag"},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}

//Populating notes table with sample notes
func Notes(db *gorm.DB) {
	fmt.Println("SETTING NOTES")
	examples := []Note{
		Note{Body: "LOLBA", VisitorID: 1},
		Note{Body: "LULBA", VisitorID: 2},
		Note{Body: "LILBA", VisitorID: 3},
	}

	for _, u := range examples {
		db.Create(&u)
	}
}

//Populating categories table with sample categories
func Categories(db *gorm.DB) {
	fmt.Println("SETTING CATEGORIES")
	examples := []Category{
		Category{Name: "Quant Trading"},
		Category{Name: "Tech"},
		Category{Name: "Value Investment"},
	}
	for _, u := range examples {
		db.Create(&u)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// the model
type Language struct {
	Name  		string 
	Projects []*Project `gorm:"many2many:language_projects;"`
	gorm.Model
} 

type Project struct {
	Name        string 
	Url 		string 
	Description string 
	Video 		string
	Languages []*Language `gorm:"many2many:language_projects;"`
	gorm.Model 
}

type Article struct {
	Title 		string 
	Description 	string 
	Publisher 	string 
	Url 		string
	Claps	 	int 
	Date 		string
}

func InitialMigration() {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	db.AutoMigrate(&Language{}, &Project{}, &Article{})
}

func AllProjects(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	db.Create(&Article{Title: "Replace Everything with Pikachu Chrome Extension", Date: "2019/09/18", Description: "Learn the fundamentals of browser extensions: what they are, how they work, why they're useful. Then build a small DOM-manipulator extension.", Claps: 272, Publisher: "Better Programming", Url: "https://medium.com/better-programming/catch-em-all-chrome-extension-d51a8b6813fd"})
	db.Create(&Article{Title: "Catch Em All Chrome Extension", Date: "2019/09/18", Description: "Step by step tutorial on how to build a Chrome Extension with background script, popup, fetches and localStorage.", Claps: 222, Publisher: "Better Programming", Url: "https://medium.com/better-programming/replace-everything-with-pikachu-chrome-extension-de40497c7f5a"})
	db.Create(&Article{Title: "Movie Comparison Website in Ruby on Rails", Date: "2019/09/2", Description: "Enjoy the fruits of your triple-join labors by displaying them on a webpage. Building Self-Joins and Triple-Joins in Ruby on Rails... Part 3 of 3.", Claps: 226, Publisher: "Better Programming", Url: "https://medium.com/better-programming/catch-em-all-chrome-extension-d51a8b6813fd"})
	db.Create(&Article{Date: "2019/09/2", Title: "Building Self-Joins and Triple Joins in Ruby on Rails", Description: "Step-by-step instructional for creating a Ruby on Rails app that implements a self-join, then triple-join relationship. Part 2 of 3...", Publisher: "Better Programming", Claps: 245, Url: "https://medium.com/better-programming/building-self-joins-and-triple-joins-in-ruby-on-rails-455701bf3fa7"})
	db.Create(&Article{Date: "2019/09/2", Title: "The Coddfather: Relational Database Fundamentals", Description: "Exploring relational databases and why they’re ubiquitous. Part 1 of 3 part series...", Publisher: "Better Programming", Claps: 269, Url: "https://medium.com/better-programming/the-coddfather-relational-database-fundamentals-533b96f87651"})
	
	// var languages []Language
	var articles []Project
	allArticles := db.Find(&articles).Value

	var datatrust Project
	db.Where("name = ?", "DataTrust").Find(&datatrust)
	db.Model(&datatrust).Association("Languages").Find(&datatrust.Languages)
	var primary Project
	db.Where("name = ?", "PrimarySource").Find(&primary)
	db.Model(&primary).Association("Languages").Find(&primary.Languages)
	var mymdb Project
	db.Where("name = ?", "MyMDb").Find(&mymdb)
	db.Model(&mymdb).Association("Languages").Find(&mymdb.Languages)
	var extension Project
	db.Where("name = ?", "DataTrust Extension").Find(&extension)
	db.Model(&extension).Association("Languages").Find(&extension.Languages)


	// for _, project := range allProjects { // can't range over an "instance"
		// db.Where("name = ?", project.Name).Find(&project)
		// db.Model(&project).Association("Languages").Find(&project.Languages)
		// allProjects = append(allProjects, project)
	// }

	allInfo := map[string]Project{"datatrust": datatrust, 
									"primarysource": primary,
									"mymdb": mymdb, 
									"extension": extension}

	

	// & == 'all'
	// var projects []Project
	json.NewEncoder(w).Encode(allInfo)
	json.NewEncoder(w).Encode(allArticles)
}

func NewProject(w http.ResponseWriter, r *http.Request) {
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	
	vars := mux.Vars(r)
	name := vars["name"]
	
	newProject := map[string]string{"name": name}
	
	db.Create(&Project{Name: name})
	json.NewEncoder(w).Encode(newProject)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	
	vars := mux.Vars(r)
	name := vars["name"]
	
	var project Project
	db.Where("name = ?", name).Find(&project)
	
	db.Delete(&project)
	json.NewEncoder(w).Encode(name)
}

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// 	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	// 	// db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// 	if err != nil {
		// 		panic("Could not connect to the database")
		// 	}
		// 	defer db.Close()
		
		// 	vars := mux.Vars(r)
		// 	name := vars["name"]
		
		// 	var user Language
		// 	db.Where("name = ?", name).Find(&user)
		
		// 	user.Email = email
		
		// 	db.Save(&user)
		// 	json.NewEncoder(w).Encode(&user)
		// 	// fmt.Fprintf(w, "Updated Language")
		// }
		
// To create a new instance in the database:
// db.Create(&Project{Name: "DataTrust"})

// Seeding looked liked this: 
	// var datatrust Project
	// db.Where("name = ?", "DataTrust").Find(&datatrust)

	// var react Language
	// db.Where("name = ?", "React.js").Find(&react)

	// db.Model(datatrust).Association("Languages").Append(react)		
/////////////////////////

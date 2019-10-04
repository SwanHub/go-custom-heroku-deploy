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
	
	// var languages []Language
	// var articles []Project
	// allArticles := db.Find(&articles).Value

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
	// json.NewEncoder(w).Encode(allArticles)
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

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
	gorm.Model
	Name  string 
	Projects []*Project `gorm:"many2many:language_projects;"`
} 

type Project struct {
	gorm.Model 
	Name string 
	Url string 
	Description string 
	Video string
	Languages []*Language `gorm:"many2many:language_projects;"`
}

func InitialMigration() {
	// db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	db.AutoMigrate(&Language{}, &Project{})
}

func AllLanguages(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	db.Create(&Project{Name: "DataTrust"})
	db.Create(&Project{Name: "PrimarySource"})
	db.Create(&Project{Name: "MyMDb"})

	db.Create(&Language{Name: "Javascript"})
	db.Create(&Language{Name: "React.js"})
	db.Create(&Language{Name: "Ruby"})
	db.Create(&Language{Name: "Ruby on Rails"})
	db.Create(&Language{Name: "Go"})
		
	// db.Model(language).Find(language)

	// var language Language
	// db.Where("name = ?", "javascript").Find(&language)

	// var project Project
	// db.Where("name = ?", "DataTrust").Find(&project)

	// db.Model(project).Association("Languages").Append(language)
	// db.Model(project).Association("Languages").Append(language)
	// db.Model(project).Association("Languages").Append(language)

	// find all projects associated with a language. 
	// db.Model(&language).Related(&projects,  "Projects")

	// & == 'all'
	var languages []Language
	db.Find(&languages)
	json.NewEncoder(w).Encode(languages)
}

func NewLanguage(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	// db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	newLanguage := map[string]string{"name": name}

	db.Create(&Language{Name: name})
	json.NewEncoder(w).Encode(newLanguage)
}

func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var language Language
	db.Where("name = ?", name).Find(&language)

	db.Delete(&language)
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

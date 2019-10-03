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
type User struct {
	gorm.Model
	Name  string
	Email string
} 

func InitialMigration() {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	// & == 'all'
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	newUser := map[string]string{"name": name, "email": email}
	newUserJson, _ := json.Marshal(newUser)

	db.Create(&User{Name: name, Email: email})
	json.NewEncoder(w).Encode(newUserJson)
	// fmt.Fprintf(w, "New User Successfully created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	json.NewEncoder(w).Encode(name)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "Updated User")
}

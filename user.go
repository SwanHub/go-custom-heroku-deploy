package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// the model
type Language struct {
	Name     string
	Projects []*Project `gorm:"many2many:language_projects;"`
	gorm.Model
}

type Project struct {
	Name        string
	Url         string
	Description string
	Video       string
	Languages   []*Language `gorm:"many2many:language_projects;"`
	gorm.Model
}

type Article struct {
	Title       string
	Description string
	Publisher   string
	Url         string
	Claps       int
	Date        string
}

type Quote struct {
	Quote  string
	Person string
}

func InitialMigration() {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	db.AutoMigrate(&Language{}, &Project{}, &Article{}, &Quote{})
}

func AllQuotes(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	db.Create(&Quote{Quote: "Life is to be lived not controlled; and humanity is won by continuing to play in the face of certain defeat.", Person: "Ralph Ellison"})
	db.Create(&Quote{Quote: "I have no special talent. I am only passionately curious.", Person: "Albert Einstein"})
	db.Create(&Quote{Quote: "Life is like riding a bicycle. To keep your balance you must keep moving.", Person: "Albert Einstein"})
	db.Create(&Quote{Quote: "The adverb is not your friend.", Person: "Stephen King"})
	db.Create(&Quote{Quote: "The instant does not have time; and time is made from the movement of the instant.", Person: "Leonardo Da Vinci"})
	db.Create(&Quote{Quote: "I was always out the door at dawn.", Person: "Phil Knight"})
	db.Create(&Quote{Quote: "No matter the sport–no matter the human endeavor, really–total effort will win people's hearts.", Person: "Phil Knight"})
	db.Create(&Quote{Quote: "Fear of failure, I thought, will never be our downfall as a company. Not that any of us thought we wouldn't fail; in fact we had every expectation that we would. But when we did fail, we had faith that we'd do it fast, learn from it, and be better for it.", Person: "Phil Knight"})
	db.Create(&Quote{Quote: "But that's the nature of money. Whether you have it or not, whether you want it or not, whether you like it or not, it will try to define your days. Our task as human beings is not to let it.", Person: "Phil Knight"})
	db.Create(&Quote{Quote: "Remembering that you are going to die is the best way I know to avoid the trap of thinking you have something to lose. You are already naked. There is no reason not to follow your heart.", Person: "Steve Jobs"})
	db.Create(&Quote{Quote: "If we made something that we wanted to see, others would want to see it, too.", Person: "Ed Catmull"})
	db.Create(&Quote{Quote: "Getting the team right is the necessary precursor to getting the ideas right. ", Person: "Ed Catmull"})
	db.Create(&Quote{Quote: "Mistakes aren't a necessary evil. They aren't evil at all. They are an inevitable consequence of doing something new (and, as such, should be seen as valuable; without them, we'd have no originality).", Person: "Ed Catmull"})
	db.Create(&Quote{Quote: "Craft is what we are expected to know; art is the unexpected use of our craft.", Person: "Ed Catmull"})
	db.Create(&Quote{Quote: "The best way to predict the future is to invent it.", Person: "Alan Kay"})
	db.Create(&Quote{Quote: "Having an ability to figure things out is more important than having specific knowledge of how to do something.", Person: "Ray Dalio"})
	db.Create(&Quote{Quote: "Try, fail, diagnose, redesign, and try again.", Person: "Ray Dalio"})
	db.Create(&Quote{Quote: "One learns from books and reels only that certain things can be done. Actual learning requires that you do those things. ", Person: "Frank Herbert"})
	db.Create(&Quote{Quote: "In general, having more data helps, but having the right data is the more important requirement.", Person: "John D. Kelleher"})
	db.Create(&Quote{Quote: "Life shrinks or expands in proportion to one's courage.", Person: "Anais Nin"})
	db.Create(&Quote{Quote: "When a great adventure is offered, you don't refuse it.", Person: "Amelia Earhart"})

	var quotes []Quote
	allQuotes := db.Find(&quotes).Value
	json.NewEncoder(w).Encode(allQuotes)
}

func AllArticles(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var articles []Article
	allArticles := db.Find(&articles).Value
	json.NewEncoder(w).Encode(allArticles)
}

func AllProjects(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

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
		"mymdb":         mymdb,
		"extension":     extension}

	// & == 'all'
	json.NewEncoder(w).Encode(allInfo)
}

// func NewProject(w http.ResponseWriter, r *http.Request) {
// 	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
// 	// db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		panic("Could not connect to the database")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	name := vars["name"]

// 	newProject := map[string]string{"name": name}

// 	db.Create(&Project{Name: name})
// 	json.NewEncoder(w).Encode(newProject)
// }

// func DeleteProject(w http.ResponseWriter, r *http.Request) {
// 	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password=JaQuEz11! dbname=test3 sslmode=disable")
// 	// db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		panic("Could not connect to the database")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	name := vars["name"]

// 	var project Project
// 	db.Where("name = ?", name).Find(&project)

// 	db.Delete(&project)
// 	json.NewEncoder(w).Encode(name)
// }

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

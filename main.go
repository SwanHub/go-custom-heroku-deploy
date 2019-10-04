package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/languages", AllLanguages).Methods("GET")
	myRouter.HandleFunc("/language/{name}", NewLanguage).Methods("POST")
	myRouter.HandleFunc("/language/{name}", DeleteLanguage).Methods("DELETE")
	// myRouter.HandleFunc("/user/{name}", UpdateLanguage).Methods("PUT")
	// log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myRouter)))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myRouter)))
}


func main() {
	fmt.Println("Go ORM")

	InitialMigration()

	handleRequests()
}
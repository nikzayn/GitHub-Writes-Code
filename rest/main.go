package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Struct for the articles
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

//Function for retrieving the get endpoint requests
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Somebody", Content: "I love the way you lie"},
		Article{Title: "Bye", Desc: "Nobody", Content: "Channa Meraya"},
	}
	fmt.Println("Endpoint Hit: allArticles")
	//Encoding our articles into a JSON string and writing as a part
	//of our response
	json.NewEncoder(w).Encode(articles)
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Endpoint Hit")
}

//Function for homepage req and res
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: homepage")
}

//Function for handling incoming requests
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/postarticles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	handleRequests()
}

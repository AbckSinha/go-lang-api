package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Title"`
	Content string `json:"Title"`
}

type Articles []Article

func allArticle(w http.ResponseWriter, r *http.Request){
	articles := Articles {
      Article{Title: "Hello Application", Desc: "This is sample app", Content:"Lorem Ipsum"},
	}
	fmt.Println(w, "All Articles has been hit")
	json.NewEncoder(w).Encode(articles)
 }

func homePage(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w, "Home page has been hit")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticle)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	handleRequest()
}

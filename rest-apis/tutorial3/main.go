package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]
	// fmt.Fprintf(w, "key: "+key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	fmt.Println("Endpoint hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: deleteArticle")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: updateArticle")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			var article Article
			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &article)
			Articles[index] = article
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8877", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "title1", Desc: "desc1", Content: "content1"},
		Article{Id: "2", Title: "title2", Desc: "desc2", Content: "content2"},
		Article{Id: "3", Title: "title2", Desc: "desc2", Content: "content2"},
		Article{Id: "4", Title: "title2", Desc: "desc2", Content: "content2"},
		Article{Id: "5", Title: "title2", Desc: "desc2", Content: "content2"},
	}
	handleRequests()
}

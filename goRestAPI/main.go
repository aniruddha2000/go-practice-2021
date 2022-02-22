package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aniruddha2000/gorestapi/controllers"
	"github.com/aniruddha2000/gorestapi/database"
	"github.com/gorilla/mux"
)

// type Article struct {
// 	ID      string `json:"id"`
// 	Title   string `json:"title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// var Articles []Article

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the homepage!")
// 	fmt.Println("Endpoint hit: homePage")
// }

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint hit: returnAllArticles")
// 	json.NewEncoder(w).Encode(Articles)
// }

// func returnSingleArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint hit: returnSingleArticles")

// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	for _, article := range Articles {
// 		if article.ID == key {
// 			json.NewEncoder(w).Encode(article)
// 		}
// 	}
// }

// func createNewArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint hit: createNewArticle")
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var newArticle Article
// 	json.Unmarshal(reqBody, &newArticle)
// 	Articles = append(Articles, newArticle)
// 	json.NewEncoder(w).Encode(newArticle)
// }

// func deleteArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint hit: deleteArticle")

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	for index, article := range Articles {
// 		if article.ID == id {
// 			Articles = append(Articles[:index], Articles[index+1:]...)
// 		}
// 	}
// }

// func updateArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint hit: updateArticle")

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var newArticle Article
// 	json.Unmarshal(reqBody, &newArticle)

// 	for index, article := range Articles {
// 		if article.ID == id {
// 			article.Title = newArticle.Title
// 			article.Desc = newArticle.Desc
// 			article.Content = newArticle.Content
// 			// Articles = append(Articles[:index], article)
// 			Articles = append(append(Articles[:index], article), Articles[index+1:]...)
// 			json.NewEncoder(w).Encode(article)
// 		}
// 	}
// }

// func handleRequests() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", homePage)
// 	router.HandleFunc("/get/articles", returnAllArticles)
// 	router.HandleFunc("/get/article/{id}", returnSingleArticles)
// 	router.HandleFunc("/post/article", createNewArticle).Methods("POST")
// 	router.HandleFunc("/delete/article/{id}", deleteArticle).Methods("DELETE")
// 	router.HandleFunc("/update/article/{id}", updateArticle).Methods("PATCH")

// 	log.Fatal(http.ListenAndServe(":10000", router))
// }

func main() {
	db := initDB()
	defer db.Close()

	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/get/articles", controllers.ReturnAllArticleHandler(db))
	router.HandleFunc("/get/article/{id}", controllers.ReturnSingleArticleHandler(db))
	router.HandleFunc("/post/article", controllers.CreateNewArticleHandler(db)).Methods("POST")
	router.HandleFunc("/delete/article/{id}", controllers.DeleteArticle(db)).Methods("DELETE")
	router.HandleFunc("/update/article/{id}", controllers.UpdateArticle(db)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8090", router))
}

func initDB() *sql.DB {
	config := database.Config{
		ServerName: "127.0.0.1:3306",
		User:       "root",
		Password:   "123",
		DB:         "ARTICLE",
	}

	connectionString := database.GetConnectionString(config)
	db, err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	return db
}

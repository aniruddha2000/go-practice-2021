package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aniruddha2000/gorestapi/controllers"
	"github.com/aniruddha2000/gorestapi/database"
	"github.com/gorilla/mux"
)

type Server struct {
	serverPort string
	router     *mux.Router
	db         *sql.DB
}

func main() {
	mysql := initDB()
	defer mysql.Close()

	server := Server{
		serverPort: ":8090",
		router:     mux.NewRouter().StrictSlash(true),
		db:         mysql,
	}

	routeHandler(&server)
}

func routeHandler(s *Server) {
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/get/articles", controllers.ReturnAllArticleHandler(s.db))
	router.HandleFunc("/get/article/{id}", controllers.ReturnSingleArticleHandler(s.db))
	router.HandleFunc("/post/article", controllers.CreateNewArticleHandler(s.db)).Methods("POST")
	router.HandleFunc("/delete/article/{id}", controllers.DeleteArticle(s.db)).Methods("DELETE")
	router.HandleFunc("/update/article/{id}", controllers.UpdateArticle(s.db)).Methods("PUT")

	log.Fatal(http.ListenAndServe(s.serverPort, router))
}

func initDB() *sql.DB {
	config := database.Config{
		ServerName: "127.0.0.1:3306",
		User:       "root",
		Password:   "123",
		DB:         "ARTICLE",
	}

	connectionString := database.GetConnectionString(&config)
	db, err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	return db
}

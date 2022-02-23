package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/aniruddha2000/gorestapi/controllers"
	"github.com/aniruddha2000/gorestapi/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Server struct {
	serverPort string
	router     *mux.Router
	db         *sql.DB
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	mysql := initDB()
	defer mysql.Close()

	server := Server{
		serverPort: os.Getenv("SERVER_PORT"),
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
		ServerName: os.Getenv("DB_SERVER"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		DB:         os.Getenv("DB_NAME"),
	}

	connectionString := database.GetConnectionString(&config)
	db, err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	return db
}

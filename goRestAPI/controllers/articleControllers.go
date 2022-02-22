package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/aniruddha2000/gorestapi/entity"
	"github.com/gorilla/mux"
)

func CreateNewArticleHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("CreateNewArticleHandler --> %s %s", r.Method, r.URL.Path)
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var article entity.Article

		json.Unmarshal(reqBody, &article)

		query := "INSERT INTO article VALUES (" + "'" + article.ID + "'" + ", " +
			"'" + article.Title + "'" + ", " + "'" + article.Description + "'" + ", " +
			"'" + article.Content + "'" + ")"
		log.Print(query)

		insert, err := db.Query(query)
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(article)
	}
}

func ReturnAllArticleHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("ReturnAllArticleHandler --> %s %s", r.Method, r.URL.Path)

		articles := queryAllRows(db)

		for _, article := range articles {
			json.NewEncoder(rw).Encode(article)
		}
	}
}

func ReturnSingleArticleHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		log.Printf("ReturnSingleArticleHandler --> %s %s", r.Method, r.URL.Path)

		var article entity.Article

		query := "SELECT * FROM article WHERE " + "ID=" + "'" + id + "'"
		log.Print(query)
		err := db.QueryRow(query).Scan(&article.ID, &article.Title, &article.Description, &article.Content)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(rw).Encode(article)
	}
}

func DeleteArticle(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		log.Printf("DeleteArticle --> %s %s", r.Method, r.URL.Path)

		query := "DELETE FROM article WHERE " + "ID=" + "'" + id + "'"
		log.Print(query)
		del, err := db.Exec(query)
		if err != nil {
			panic(err)
		}

		rowsAffected, err := del.RowsAffected()
		if err != nil {
			panic(err)
		}

		jsonData := map[string]string{"Rows affected": strconv.Itoa(int(rowsAffected)), "status": "deleted successfully"}
		json.NewEncoder(rw).Encode(jsonData)
	}
}

func UpdateArticle(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		log.Printf("UpdateArticle --> %s %s", r.Method, r.URL.Path)

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var article entity.Article

		json.Unmarshal(reqBody, &article)

		query := "UPDATE article SET " + "Title=" + "'" + article.Title + "', " +
			"Description=" + "'" + article.Description + "', " + "Content=" + "'" +
			article.Content + "' " + "WHERE ID=" + "'" + id + "'"
		log.Print(query)

		res, err := db.Exec(query)
		if err != nil {
			panic(err)
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}

		jsonData := map[string]string{"Rows affected": strconv.Itoa(int(rowsAffected)), "status": "updated successfully"}
		json.NewEncoder(rw).Encode(jsonData)
	}
}

func queryAllRows(db *sql.DB) []entity.Article {
	var articles []entity.Article

	query := "SELECT * FROM article"
	log.Print(query)

	res, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer res.Close()

	for res.Next() {
		var article entity.Article
		err := res.Scan(&article.ID, &article.Title, &article.Description, &article.Content)
		if err != nil {
			panic(err)
		}
		articles = append(articles, article)
	}

	return articles
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

/*
	Setup mysql with docker
	$ docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123 -d mysql:8.0
*/

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL Tutorial" + string(3))

	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// InsertIntoDB(db)

	// ScanFromDB(db)

	GetRowFromTable(db, 3)
}

func InsertIntoDB(db *sql.DB) {
	datas := []struct {
		id   int
		name string
	}{
		{3, "TEST3"},
		{4, "TEST4"},
		{5, "TEST5"},
		{6, "TEST6"},
		{7, "TEST7"},
	}

	for _, data := range datas {
		// insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
		query := "INSERT INTO test VALUES (" + strconv.Itoa(data.id) + ", " + "'" + data.name + "'" + ")"
		fmt.Println(query)
		insert, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	}
}

func ScanFromDB(db *sql.DB) {
	res, err := db.Query("SELECT id, name FROM test")
	if err != nil {
		panic(err.Error())
	}

	for res.Next() {
		var tag Tag

		err = res.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}

		log.Printf(tag.Name)
	}
}

func GetRowFromTable(db *sql.DB, id int) {
	var tag Tag
	err := db.QueryRow("SELECT id, name FROM test where id = ?", id).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error())
	}

	log.Println(tag.ID)
	log.Println(tag.Name)

}

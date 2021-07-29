package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/tes")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var yolo []Data
		rows, err := db.Query("SELECT * FROM tes_a")
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var data Data
			err = rows.Scan(&data.ID, &data.Name)

			yolo = append(yolo, data)
		}

		ssds, err := json.Marshal(yolo)

		w.Write([]byte(ssds))
	})

	server := &http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()
	fmt.Println("server started")
}

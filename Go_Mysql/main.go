package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Tutorial")

	// Create Connection
	database, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tes_go")

	if err != nil {
		panic(err.Error())
	} else {
		defer database.Close()

		// Creating tables
		create, err := database.Query("CREATE TABLE IF NOT EXISTS mahasiswa (id INTEGER PRIMARY KEY,nama VARCHAR(30) NOT NULL)")

		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println("Create table mahasiswa sukses!")
			defer create.Close()

			// Insert data
			// insert, err := database.Query("INSERT INTO mahasiswa (id,nama) VALUES (2,'Alex Sander Hutapea')")

			// if err != nil {
			// 	panic(err.Error())
			// } else {
			// 	fmt.Println("Insert Berhasil!")
			// 	defer insert.Close()
			// }

			// Select data
			table := "mahasiswa"
			query := "SELECT * FROM " + table
			qselect, err := database.Query(query)

			if err != nil {
				panic(err.Error())
			} else {
				fmt.Println("Select Berhasil!")
				defer qselect.Close()

				for qselect.Next() {
					var id int
					var nama string

					err = qselect.Scan(&id, &nama)

					if err != nil {
						panic(err.Error())
					} else {
						fmt.Println("Id : ", id)
						fmt.Println("Nama : ", nama)
					}
				}
			}
		}

	}
}

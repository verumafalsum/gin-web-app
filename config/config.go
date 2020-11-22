package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3" // SQLite3 driver
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("File .env is not found")
	}
}

// SetConfig ....
func SetConfig() {
	loadEnv()

	dbName, exists := os.LookupEnv("SQLLITE_DB_NAME")

	// I'm learning how to working with db. I will delete it
	if exists {
		db, err := sql.Open("sqlite3", dbName)

		if err != nil {
			log.Fatal(err, "Cannot connect to database")
		}

		result, err := db.Query("SELECT * FROM articles")

		if err != nil {
			log.Fatal("result error")
		}

		type art struct {
			id      int
			title   string
			content string
		}

		arts := []art{}

		for result.Next() {
			a := art{}
			err := result.Scan(&a.id, &a.title, &a.content)
			if err != nil {
				log.Fatal(err)
				continue
			}
			arts = append(arts, a)
		}

		for _, a := range arts {
			fmt.Println(a.id, a.title, a.content)
		}
	}

}

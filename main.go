package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	envs := map[string]string{
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"PORT":        os.Getenv("PORT"),
	}

	for k, v := range envs {
		if envs[k] == "" {
			log.Fatalf("%v doesn't exist", envs[v])
		}
	}

	var port string
	if envs["PORT"] == "" {
		port = "8000" // Default port if PORT environment variable is not set
	} else {
		port = envs["PORT"]
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable", envs["DB_HOST"], envs["DB_USER"], envs["DB_PASSWORD"], envs["DB_NAME"])
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("database is connected successfully")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pid := fmt.Sprintf("Hello world! Process ID : %d, Port: %s, Database: %v", os.Getpid(), port, db.Dialector.Name())
		w.Write([]byte(pid))
	})

	http.ListenAndServe(":"+port, nil)
}

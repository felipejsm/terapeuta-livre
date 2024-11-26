package main

import (
	"errors"
	"felipejsm/tp-admin/internal/config"
	database "felipejsm/tp-admin/internal/db"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	config.LoadEnv()
	database.InitDB()
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		database := config.GetEnv("DATABASE_URL", "default_value")
		fmt.Print(w, database)
	})
	sqlDB, _ := database.DB.DB()
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
		defer sqlDB.Close()
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		defer sqlDB.Close()
		os.Exit(1)
	}
	defer sqlDB.Close()
	fmt.Println("Server starte @ localhost:8080/hello")

}

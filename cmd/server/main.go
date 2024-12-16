package main

import (
	"errors"
	database "felipejsm/tp-admin/internal/db"
	"felipejsm/tp-admin/internal/handlers"
	repository "felipejsm/tp-admin/internal/repositories"
	"felipejsm/tp-admin/internal/services"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	sqlDB := database.InitDB()
	// repository
	repo := repository.NewPatientRepository(sqlDB)

	//services
	service := services.NewPatientService(repo)

	//template
	templates := template.Must(template.ParseGlob("internal/templates/*.html"))

	//handles
	handle := handlers.NewPatientHandler(service, templates)

	http.HandleFunc("/patients", handle.HandleGetPatient)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Server start listening @ port 8080")
}

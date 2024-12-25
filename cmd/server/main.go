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
	// Inicializa o banco de dados
	sqlDB := database.InitDB()

	// Repositórios
	repo := repository.NewPatientRepository(sqlDB)
	therapistRepo := repository.NewTherapistRepository(sqlDB)

	// Serviços
	patientService := services.NewPatientService(repo)
	therapistService := services.NewTherapistService(therapistRepo)

	// Carrega os templates
	templates := template.Must(template.ParseGlob("internal/templates/*.html"))

	// Handlers
	patientHandler := handlers.NewPatientHandler(patientService, templates)
	therapistHandler := handlers.NewTherapistHandler(therapistService, templates)
	layoutHandler := handlers.NewLayoutHandler(templates)

	// Roteamento
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			err := templates.ExecuteTemplate(w, "layout.html", nil)
			if err != nil {
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		
			}
			layoutHandler.HandleLayout(w, r)
			return
		}

		// Se não for "/", retorna erro 404
		http.NotFound(w, r)
	})

	http.HandleFunc("/patients", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			patientHandler.HandleGetPatient(w, r)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/therapist", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			therapistHandler.HandleGetTherapist(w, r)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Inicia o servidor
	fmt.Println("Server start listening @ port 8080")
	err := http.ListenAndServe(":8080", nil)

	// Tratamento de erros do servidor
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

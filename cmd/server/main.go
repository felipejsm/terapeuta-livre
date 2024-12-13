package main

import (
	"errors"
	"felipejsm/tp-admin/internal/config"
	database "felipejsm/tp-admin/internal/db"
	repository "felipejsm/tp-admin/internal/repositories"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	config.LoadEnv()
	database.InitDB()
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		datab := config.GetEnv("DATABASE_URL", "default_value")
		patientRep := repository.PatientRepository{DB: database.DB}
		patients, err := patientRep.FindAllByTherapistId("1")
		if err != nil {
			fmt.Println("Erro ao buscar pacientes:", err)
		} else {
			fmt.Printf("Pacientes encontrados: %v\n", patients)
		}
		fmt.Print(w, datab)
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

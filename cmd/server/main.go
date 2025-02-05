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
    "strings"
    "path/filepath"
)

func main() {
	// Inicializa o banco de dados
	sqlDB := database.InitDB()

	// Repositórios
	repo := repository.NewPatientRepository(sqlDB)
	therapistRepo := repository.NewTherapistRepository(sqlDB)
	fileMetadataRepo := repository.NewFileMetadataRepository(sqlDB)
    fileRepo := repository.NewFileRepository(sqlDB)

	// Serviços
	fileMetadataService := services.NewFileMetadataService(fileMetadataRepo)
	patientService := services.NewPatientService(repo)
	therapistService := services.NewTherapistService(therapistRepo)
    fileService := services.NewFileService(fileRepo)
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal("Erro ao obter o diretório de execução:", err)
    }
    templatePath := filepath.Join(dir, "internal", "templates", "*.html")
	// Carrega os templates
	templates := template.Must(template.ParseGlob(templatePath))
	for _, tmplName := range templates.Templates() {
		fmt.Printf("Template carregado: %s\n", tmplName.Name())
	}
	// Configuração de arquivos estáticos
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handlers
	fileMetadataHandler := handlers.NewFileMetadataHandler(fileMetadataService, templates, fileService)
	patientHandler := handlers.NewPatientHandler(patientService, templates)
	therapistHandler := handlers.NewTherapistHandler(therapistService, templates)
	layoutHandler := handlers.NewLayoutHandler(templates)

	// Roteamento
	http.HandleFunc("/", layoutHandler.HandleLayout)

	http.HandleFunc("/patient", patientHandler.HandleGetPatient)

	http.HandleFunc("/therapist", therapistHandler.HandleGetTherapist)

	http.HandleFunc("/file_metadata", fileMetadataHandler.HandleGetFileMetadata)

    http.HandleFunc("/file/", func(w http.ResponseWriter, r *http.Request) {
    // Extrai o `id` da rota
    id := strings.TrimPrefix(r.URL.Path, "/file/")
    if id == "" {
        http.Error(w, "ID not provided", http.StatusBadRequest)
        return
    }

    fileMetadataHandler.HandleFileDownload(w, r, id)
})
	// Inicia o servidor
	fmt.Println("Server start listening @ port 8081")
	err = http.ListenAndServe(":8081", nil)

	// Tratamento de erros do servidor
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

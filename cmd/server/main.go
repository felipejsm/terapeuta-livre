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
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
    "context"
    "time"
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
    var templatePath string
    isProd := os.Getenv("PROD")
    if isProd != "" {
        templatePath = filepath.Join(dir, "..","..", "internal", "templates", "*.html")
    } else {
        templatePath = filepath.Join(dir, "internal", "templates", "*.html")
    }
    log.Println("Caminho dos templates:", templatePath)
	// Carrega os templates
	templates := template.Must(template.ParseGlob(templatePath))
	for _, tmplName := range templates.Templates() {
		fmt.Printf("Template carregado: %s\n", tmplName.Name())
	}
    var staticPath string
    if isProd != "" {
        staticPath = filepath.Join(dir, "..","..", "web", "static")
    } else {
        staticPath = filepath.Join(dir, "web", "static")
    }
    // Configuração de arquivos estáticos
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handlers
	fileMetadataHandler := handlers.NewFileMetadataHandler(fileMetadataService, templates, fileService)
	patientHandler := handlers.NewPatientHandler(patientService, templates)
	therapistHandler := handlers.NewTherapistHandler(therapistService, templates)
	layoutHandler := handlers.NewLayoutHandler(templates)
    loginHandler := handlers.NewLoginHandler(templates)

	// Roteamento
	http.HandleFunc("/", SessionMiddleware(layoutHandler.HandleLayout))

	http.HandleFunc("/patient", SessionMiddleware(patientHandler.HandleGetPatient))

	http.HandleFunc("/therapist", SessionMiddleware(therapistHandler.HandleGetTherapist))

	http.HandleFunc("/file_metadata", SessionMiddleware(fileMetadataHandler.HandleGetFileMetadata))

    http.HandleFunc("/login", loginHandler.HandleLogin)

    http.HandleFunc("/file/", SessionMiddleware(func(w http.ResponseWriter, r *http.Request) {
    // Extrai o `id` da rota
    id := strings.TrimPrefix(r.URL.Path, "/file/")
    if id == "" {
        http.Error(w, "ID not provided", http.StatusBadRequest)
        return
    }
    fileMetadataHandler.HandleFileDownload(w, r, id)
}))
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

func SessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. Recuperar o cookie "hanko"
        cookie, err := r.Cookie("hanko")
        if err != nil {
            if err == http.ErrNoCookie {
                http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
                return
            }
            http.Error(w, "Erro ao recuperar cookie", http.StatusInternalServerError)
            return
        }
        hankoApiURL := os.Getenv("HANKO_URL")
        // 2. Buscar as chaves públicas (JWKS)
        set, err := jwk.Fetch(
            context.Background(),
            fmt.Sprintf("%v/.well-known/jwks.json", hankoApiURL),
        )
        if err != nil {
            log.Println("Erro ao buscar JWKS:", err)
            http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
            return
        }

        // 3. Validar o token JWT
        token, err := jwt.Parse([]byte(cookie.Value), jwt.WithKeySet(set))
        if err != nil {
            log.Println("Erro ao validar token:", err)
            http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
            return
        }

        // 4. Verificar expiração do token
        if token.Expiration().Before(time.Now()) {
            log.Println("Token expirado para usuário:", token.Subject())
            http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
            return
        }

        // 5. Adicionar informações no contexto e continuar requisição
        log.Printf("Sessão válida para usuário: %s", token.Subject())

        // Define o usuário no contexto
        ctx := context.WithValue(r.Context(), "user", token.Subject())

        // Chama o handler original
        next(w, r.WithContext(ctx))
    }
}


package handlers

import (
	"felipejsm/tp-admin/internal/services"
	"fmt"
	"html/template"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

type PatientHandler struct {
	Service   *services.PatientService
	Templates *template.Template
}

func NewPatientHandler(service *services.PatientService, templates *template.Template) *PatientHandler {
	return &PatientHandler{
		Service:   service,
		Templates: templates,
	}
}

func (h *PatientHandler) HandleGetPatient(w http.ResponseWriter, r *http.Request) {
	// temporariamente fazer um new no renderer
	rnd := renderer.New(renderer.Options{
		TemplateDir:      "view",
		ParseGlobPattern: "internal/templates/*.html",
	})

	fmt.Printf("Template @ Patient -> %s", h.Templates.Name())
	if r.Method == http.MethodGet && r.URL.Path == "/patients" {
		//pegar ids da sessão ou req
		data, err := h.Service.GetPatientDetail(1, 1)
		fmt.Println("Dados patients: Name ->  " + data.Name)
		if err != nil {
			http.Error(w, "Paciente não encontrado", http.StatusNotFound)
			return
		}
		templateName := "patient"
		fmt.Printf("Data full: %v", data)
		err = rnd.View(w, http.StatusOK, templateName, nil)
		/*
			err = h.Templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
				"Content": templateName,
				"Data":    data,
			})
		*/
		if err != nil {
			http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

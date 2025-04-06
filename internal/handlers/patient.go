package handlers

import (
	"felipejsm/tp-admin/internal/services"
	"fmt"
	"html/template"
	"net/http"
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
	fmt.Printf("Template @ Patient -> %s", h.Templates.Name())
	if r.Method == http.MethodGet && r.URL.Path == "/patient" {
		// Recuperar o ID do terapeuta do contexto
		therapistID := r.Context().Value("therapist_id").(uint)

		data, err := h.Service.GetPatientDetail(1, int(therapistID))
		fmt.Println("Dados patients: Name ->  " + data.Name)
		if err != nil {
			http.Error(w, "Paciente não encontrado", http.StatusNotFound)
			return
		}
		templateName := "patient_detail"
		fmt.Printf("Data full: %v", data)
		err = h.Templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
			"TemplateName": templateName,
			"Data":         data,
		})
		if err != nil {
			http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

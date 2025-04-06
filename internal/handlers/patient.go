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

// HandleNewPatient renders the patient registration form
func (h *PatientHandler) HandleNewPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Path == "/patients/new" {
		// Render the patient registration form
		err := h.Templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
			"TemplateName": "patient_new",
		})
		if err != nil {
			http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// HandleCreatePatient processes the patient registration form submission
func (h *PatientHandler) HandleCreatePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && r.URL.Path == "/patients" {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
			return
		}

		// Get the therapist ID from the context
		therapistID := r.Context().Value("therapist_id").(uint)

		// Create a new patient from the form data
		// This will be implemented in the service layer
		// For now, we'll just redirect to the patients list
		http.Redirect(w, r, "/patients", http.StatusSeeOther)
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

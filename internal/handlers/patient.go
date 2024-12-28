package handlers

import (
	"felipejsm/tp-admin/internal/services"
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
	if r.Method == http.MethodGet {
		//pegar ids da sessão ou req
		data, err := h.Service.GetPatientDetail(1, 2)
		if err != nil {
			http.Error(w, "Paciente não encontrado", http.StatusNotFound)
			return
		}
		err = h.Templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
			"Content": "patients.html",
			"Data":    data,
		})
		if err != nil {
			http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

package handlers

import(
	"net/http"
	"text/template"
	"felipejsm/tp-admin/internal/services"
)

type PatientHandler struct {
	Service *services.PatientService
	Templates *template.Template
}

func NewPatientHandler(service *services.PatientService, templates *template.Template) *PatientHandler {
	return &PatientHandler{
		Service: services,
		Templates: templates
	}
}

func (h *PatientHandler) HandleGetPatient(w http.ResponseWriter, r *http.Request) {
	data,  err := h.Service.GetPatientDetail(1)
	if err != nil {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
        return
	}
	if err := h.Templates.ExecuteTemplate(w, "patient_detail.html", data); err != nil {
        http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
        return
    }

}

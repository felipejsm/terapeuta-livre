package handlers

import (
	"felipejsm/tp-admin/internal/services"
	"html/template"
	"net/http"
)

type TherapistHandler struct {
	Service   *services.TherapistService
	Templates *template.Template
}

func NewTherapistHandler(service *services.TherapistService, templates *template.Template) *TherapistHandler {
	return &TherapistHandler{
		Service:   service,
		Templates: templates,
	}
}

func (h *TherapistHandler) HandleGetTherapist(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		data, err := h.Service.GetTherapistDetail(1)
		if err != nil {
			http.Error(w, "Terapeuta não encontrado", http.StatusNotFound)
			return
		}
		err = h.Templates.ExecuteTemplate(w, "therapist.html", data)
		if err != nil {
			http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)

		}
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

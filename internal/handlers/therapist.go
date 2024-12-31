package handlers

import (
	"felipejsm/tp-admin/internal/services"
	"fmt"
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
	fmt.Printf("Template @ Therapist -> %s", h.Templates.Name())
	if r.Method == http.MethodGet && r.URL.Path == "/therapist" {

		data, err := h.Service.GetTherapistDetail(1)
		if err != nil {
			http.Error(w, "Terapeuta não encontrado", http.StatusNotFound)
			return
		}
		fmt.Printf("Renderizando layout com Content: %s e Data: %+v\n", "therapist.html", data)
		err = h.Templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
			"Content": "therapist.html",
			"Data":    data,
		})
		if err != nil {
			http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)

		}
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

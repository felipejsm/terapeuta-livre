package handlers

import (
	"html/template"
	"net/http"
)

type LayoutHandler struct {
	Templates *template.Template
}

func NewLayoutHandler(template *template.Template) *LayoutHandler {
	return &LayoutHandler{Templates: template}
}

func (h *LayoutHandler) HandleLayout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := h.Templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
			"Content": "layout.html",
			"Data":    nil,
		})
		if err != nil {
			http.Error(w, "Erro ao renderizar layout principal", http.StatusInternalServerError)
		}
	} else {
		http.NotFound(w, r)
	}
}

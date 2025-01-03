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
		templateName := "layout.html"
		err := h.Templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
			"Content": templateName,
			"Data":    nil,
		})
		if err != nil {
			http.Error(w, "Erro ao renderizar layout principal", http.StatusInternalServerError)
		}
	} else {
		http.NotFound(w, r)
	}
}

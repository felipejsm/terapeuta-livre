package handlers

import (
	"html/template"
	"net/http"
)

type LayoutHandler struct {
	Templates *template.Template
}

func (h *LayoutHandler) HandleLayout(w http.ResponseWriter, r *http.Request) {

	if err := h.Templates.ExecuteTemplate(w, "layout.html", nil); err != nil {
		http.Error(w, "Erro ao renderizar layout principal", http.StatusInternalServerError)
		return
	}

}

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
	// Rota para servir arquivos estáticos
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := h.Templates.ExecuteTemplate(w, "layout.html", nil); err != nil {
		http.Error(w, "Erro ao renderizar layout principal", http.StatusInternalServerError)
		return
	}

}

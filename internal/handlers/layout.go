package handlers

import (
	"html/template"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

type LayoutHandler struct {
	Templates *template.Template
}

func NewLayoutHandler(template *template.Template) *LayoutHandler {
	return &LayoutHandler{Templates: template}
}

func (h *LayoutHandler) HandleLayout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rnd := renderer.New(renderer.Options{
			TemplateDir:      "view",
			ParseGlobPattern: "internal/templates/*.html",
		})
		templateName := "layout"
		err := rnd.View(w, http.StatusOK, templateName, nil)
		/*
			err := h.Templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
				"Content": templateName,
				"Data":    nil,
			})
		*/
		if err != nil {
			http.Error(w, "Erro ao renderizar layout principal", http.StatusInternalServerError)
		}
	} else {
		http.NotFound(w, r)
	}
}

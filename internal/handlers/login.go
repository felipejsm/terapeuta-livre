package handlers

import (
    "net/http"
    "html/template"
)

type LoginHandler struct {
    Templates *template.Template
}

func NewLoginHandler(templates *template.Template) *LoginHandler {
    return &LoginHandler{
        Templates: templates,
    }
}

func (h *LoginHandler) HandleLogin (w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        templateName := "login"
        err := h.Templates.ExecuteTemplate(w, "login.html", map[string]interface{}{
            "TemplateName": templateName,
            "Data": nil,
        })
        if err != nil {
            http.Error(w, "Erro ao renderizar login", http.StatusInternalServerError)
        } else {
            http.NotFound(w, r)
        }
    }
} 

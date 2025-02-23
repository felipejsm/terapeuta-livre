package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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
    fmt.Println("/login")
    if r.Method == http.MethodGet {
        templateName := "login"
        err := h.Templates.ExecuteTemplate(w, "login", map[string]interface{}{
            "TemplateName": templateName,
            "Data": nil,
        })
        if err != nil {
            http.Error(w, err.Error() , http.StatusInternalServerError)
            log.Println(err)
        } 
    } else {
        http.NotFound(w, r)
    }
} 

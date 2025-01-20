package handlers

import (
	"felipejsm/tp-admin/internal/services"
	"html/template"
	"net/http"
)

type FileMetadataHandler struct{
	Service *services.FileMetadataService
	Template *template.Template
}

func NewFileMetadataHandler(service *services.FileMetadataService, templates *template.Template) *FileMetadataHandler {
	return &FileMetadataHandler{
		Service: service,
		Template: templates,
	}
}

func (h *FileMetadataHandler) HandleGetFileMetadata(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := h.Service.GetFilesMetadata(1)
		if err != nil {
			http.Error(w, "Sem Arquivos", http.StatusNoContent)
			return
		}
		templateName := "file_metadata"
		err = h.Template.ExecuteTemplate(w, "layout.html", map[string]interface{}{
			"TemplateName": templateName,
			"Data": data,
		})
		if err != nil {
			http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)

		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}

}

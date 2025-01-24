package handlers

import (
	"felipejsm/tp-admin/internal/services"
	"html/template"
	"net/http"
    "io"
    "encoding/json"
)

type FileMetadataHandler struct{
	Service *services.FileMetadataService
    FileService *services.FileService
	Template *template.Template
}

func NewFileMetadataHandler(service *services.FileMetadataService, templates *template.Template, fileService *services.FileService) *FileMetadataHandler {
	return &FileMetadataHandler{
		Service: service,
        FileService: fileService,
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

		} else  if r.Method == http.MethodPost && r.Header.Get("Content-Type") == "multipart/form-data"{
            r.Body = http.MaxBytesReader(w, r.Body, 10 << 20)
            if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "O arquivo excede o limite de 10MB", http.StatusBadRequest)
		return
	}

            file, header, err := r.FormFile("arquivo")
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo", http.StatusBadRequest)
		return
	}
	defer file.Close()
            fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo", http.StatusInternalServerError)
		return
	}
            h.FileService.UploadFile(1, fileBytes)
w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Arquivo salvo com sucesso",
		"nome":    header.Filename,
	})
        } else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}

}

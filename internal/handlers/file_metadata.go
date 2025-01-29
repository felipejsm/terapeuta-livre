package handlers

import (
	"encoding/json"
	"felipejsm/tp-admin/internal/services"
	"fmt"
	"html/template"
	"io"
	"net/http"
    "strconv"
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
func (h *FileMetadataHandler) HandleFileDownload(w http.ResponseWriter, r *http.Request, fileID string) {
    fmt.Printf("[HandleFileDownload] Func init with id %v and method %v", fileID, r.Method)
    id, err := strconv.Atoi(fileID)
     if r.Method == http.MethodGet {
        if err != nil {
            fmt.Printf("Error during id conversion")
        }
        file, err := h.FileService.DownloadFile(id) 
         if err != nil {
                http.Error(w, "Arquivo não encontrado", http.StatusNotFound)
                return
            }
            // Define cabeçalhos apropriados para o download
            w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "file_name.pdf"))
            mimeType := http.DetectContentType(file.FileData)
            w.Header().Set("Content-Type", mimeType)
            w.Header().Set("Content-Length",     fmt.Sprintf("%d", len(file.FileData)))

            // Escreve os bytes do arquivo na resposta
            _, err = w.Write(file.FileData)
            if err != nil {
                http.Error(w, "Erro ao enviar o arquivo", http.StatusInternalServerError)
            }    
            fmt.Printf("[DownloadFile] Func end")
    } else if r.Method == http.MethodDelete {
        fmt.Printf("File Delete begin")
        _, err := h.FileService.DeleteFile(id)
        if err != nil {
            http.Error(w, "Erro ao deletar arquivo", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    }
}

func (h *FileMetadataHandler) HandleGetFileMetadata(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
			fmt.Printf("GET")
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

			}
        } else if r.Method == http.MethodPost {
            fmt.Print("File intercepted")
            r.Body = http.MaxBytesReader(w, r.Body, 10 << 20)
            if err := r.ParseMultipartForm(10 << 20); err != nil {
				http.Error(w, "O arquivo excede o limite de 10MB", http.StatusBadRequest)
				return
			}

			file, header, err := r.FormFile("files")
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

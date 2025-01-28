package services

import (
	"felipejsm/tp-admin/internal/models"
	"felipejsm/tp-admin/internal/repositories"
	"fmt"
)

type FileService struct {
	Repo *repositories.FileRepository
}

func NewFileService(repo *repositories.FileRepository) *FileService {
	return &FileService{
		Repo: repo,
	}
}

func (s *FileService) DeleteFile(id int) (string, error) {
    var result string
    result, err := s.Repo.DeleteFile(id)
    if err != nil {
        return "nok", err
    }
    return result, nil
}

func (s *FileService) DownloadFile(id int) (*models.File, error) {
    var fileDownload models.File
    fileDownload, err := s.Repo.DownloadFile(id)
    if err != nil {
        fmt.Printf("[Service] Erro @ DownloadFile %v", err)
        return nil, err
    }

    return &fileDownload, nil 

}

func (s *FileService) UploadFile(metadataId int, file []byte) (*models.File, error) {
	var fileUpload models.File
	fileUpload, err := s.Repo.UploadFile(metadataId, file)
	if err != nil {
		fmt.Printf("[Service] Error @ UploadFile %v", err)
		return nil, err
	}
	return &fileUpload, nil

}

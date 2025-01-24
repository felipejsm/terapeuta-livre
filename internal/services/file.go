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

func (s *FileService) UploadFile(metadataId int, file []byte) (*models.File, error) {
	var fileUpload models.File
	fileUpload, err := s.Repo.UploadFile(metadataId, file)
	if err != nil {
		fmt.Printf("[Service] Error @ UploadFile %v", err)
		return nil, err
	}
	return &fileUpload, nil

}

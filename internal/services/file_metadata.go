package services

import (
	"felipejsm/tp-admin/internal/models"
	repository "felipejsm/tp-admin/internal/repositories"
	"fmt"
)

type FileMetadataService struct {
	Repo *repository.FileMetadataRepository
}

func NewFileMetadataService(repo *repository.FileMetadataRepository) *FileMetadataService {
	return &FileMetadataService{
		Repo: repo,
	}
}

func (s *FileMetadataService) GetFilesMetadata(ownerId int) (*[]models.FileMetadata, error) {
	var fileMetadata []models.FileMetadata
	fileMetadata, err := s.Repo.FindAllByOwnerId(ownerId)
	if err != nil {
		fmt.Printf("[Service] Error @ GetFilesMetadata %v", err)
		return nil, err
	}
	return &fileMetadata, nil
}

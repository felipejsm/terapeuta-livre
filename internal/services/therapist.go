package services

import (
	"felipejsm/tp-admin/internal/dtos"
	"felipejsm/tp-admin/internal/models"
	repository "felipejsm/tp-admin/internal/repositories"
)

type TherapistService struct {
	Repo *repository.TherapistRepository
}

func NewTherapistService(repo *repository.TherapistRepository) *TherapistService {
	return &TherapistService{
		Repo: repo,
	}
}

func (s *TherapistService) GetTherapistDetail(id int) (*dtos.TherapistDto, error) {
	var therapist models.Therapist
	var files []models.FileMetadata
	var patients []models.Patient
	therapist, err := s.Repo.FindById(id)
	files, err = s.Repo.FindAllFilesByTherapistId(id)
	patients, err = s.Repo.FindAllPatientsById(id)
	patientsDto := dtos.PatientsToDtos(patients)
	response := dtos.TherapistToDto(therapist, patientsDto, files)
	return response, err
}

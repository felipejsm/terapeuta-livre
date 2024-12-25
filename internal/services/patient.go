package services

import (
	"felipejsm/tp-admin/internal/dtos"
	"felipejsm/tp-admin/internal/models"
	repository "felipejsm/tp-admin/internal/repositories"
)

type PatientService struct {
	Repo *repository.PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{
		Repo: repo,
	}

}

func (s *PatientService) GetPatientDetail(patientId int, therapistId int) (*dtos.PatientDto, error) {
	var patient models.Patient
	var files []models.FileMetadata
	patient, err := s.Repo.FindByIdAndTherapistId(patientId, therapistId)
	if err != nil {
		return nil, err
	}
	files, err = s.Repo.FindAllFilesByPatientId(patient.ID)
	response := dtos.PatientToDto(patient, files)
	return response, nil
}

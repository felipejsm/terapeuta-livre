package services

import (
	"felipejsm/tp-admin/internal/dtos"
	"felipejsm/tp-admin/internal/models"
	repository "felipejsm/tp-admin/internal/repositories"

	"gorm.io/gorm"
)

type PatientService struct {
	repo *repository.PatientRepository
}

func NewPatientService(db *gorm.DB) *PatientService {
	return &PatientService{
		repo: repository.NewPatientRepository(db),
	}

}

func (s *PatientService) GetPatientDetail(patientId int, therapistId int) (*dtos.PatientDto, error) {
	var patient models.Patient
	var files []models.FileMetadata
	patient, err := s.repo.FindByIdAndTherapistId(patientId, therapistId)
	if err != nil {
		return nil, err
	}
	files, err = s.repo.FindAllFilesByPatientId(patient.ID)
	response := dtos.PatientToDto(patient, files)
	return response, nil
}

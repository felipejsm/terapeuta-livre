package services

import (
	"felipejsm/tp-admin/internal/dtos"
	"felipejsm/tp-admin/internal/models"
	repository "felipejsm/tp-admin/internal/repositories"
	"fmt"
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
	fmt.Printf("@@ Antes da pesquisa")
	patient, err := s.Repo.FindByIdAndTherapistId(patientId, therapistId)
	fmt.Printf("#### Paciente Full: %v", patient)
	if err != nil {
		fmt.Printf("[Service]Error @ GetPatientDetail %v", err)
		return nil, err
	}
	fmt.Printf("\nPatient Id: %v\nTherapist Id: %v", patientId, therapistId)
	files, err = s.Repo.FindAllFilesByPatientId(patient.ID)
	fmt.Printf("Após repo com patient.Name %v", patient.Name)
	response := dtos.PatientToDto(patient, files)
	return response, nil
}

package services

import (
	"felipejsm/tp-admin/internal/dtos"
	repo "felipejsm/tp-admin/internal/repositories"
	model "felipejsm/tp-admin/internal/models"
)

func GetPatientDetail(patientId string) (dtos.PatientDto, error) {
	var patient model.Patient
	repo.PatientRepository.DB.
}

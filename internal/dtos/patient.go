package dtos

import "felipejsm/tp-admin/internal/models"

type PatientDto struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name"`
	Email       string                `json:"email"`
	IdTherapist int                   `json:"id_therapist"`
	Files       []models.FileMetadata `json:"files"`
}

func PatientToDto(patient models.Patient, files []models.FileMetadata) PatientDto {
	return PatientDto{
		ID:          patient.ID,
		Name:        patient.Name,
		Email:       patient.Email,
		IdTherapist: patient.TherapistId,
		Files:       files,
	}

}

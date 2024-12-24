package dtos

import "felipejsm/tp-admin/internal/models"

type PatientDto struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name"`
	Email       string                `json:"email"`
	IdTherapist int                   `json:"id_therapist"`
	Files       []models.FileMetadata `json:"files"`
}

func PatientsToDtos(patients []models.Patient) []PatientDto {
	var patientsDtos []PatientDto
	for _, v := range patients {
		patientDto := PatientDto{
			ID:          v.ID,
			Name:        v.Name,
			Email:       v.Email,
			IdTherapist: v.TherapistId,
			Files:       nil,
		}
		patientsDtos = append(patientsDtos, patientDto)
	}
	return patientsDtos
}
func PatientToDto(patient models.Patient, files []models.FileMetadata) *PatientDto {
	return &PatientDto{
		ID:          patient.ID,
		Name:        patient.Name,
		Email:       patient.Email,
		IdTherapist: patient.TherapistId,
		Files:       files,
	}

}

func DtoToPatient(patientDto PatientDto) (models.Patient, []models.FileMetadata) {
	var patient = models.Patient{
		ID:          patientDto.ID,
		Name:        patientDto.Name,
		Email:       patientDto.Email,
		TherapistId: patientDto.IdTherapist,
	}
	return patient, patientDto.Files
}

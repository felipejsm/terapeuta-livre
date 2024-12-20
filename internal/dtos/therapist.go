package dtos

import "felipejsm/tp-admin/internal/models"

type TherapistDto struct {
	ID       uint                  `json:"id"`
	Name     string                `json:"name"`
	Email    string                `json:"email"`
	Patients []PatientDto          `json:"patients"`
	Files    []models.FileMetadata `json:"files"`
}

func TherapistToDto(therapist models.Therapist, patients []PatientDto, files []models.FileMetadata) *TherapistDto {
	return &TherapistDto{
		ID:       therapist.ID,
		Name:     therapist.Name,
		Email:    therapist.Email,
		Patients: patients,
		Files:    files,
	}
}

func DtoToTherapist(therapistDto TherapistDto) (models.Therapist, []PatientDto, []models.FileMetadata) {
	var therapist = models.Therapist{
		ID:    therapistDto.ID,
		Name:  therapistDto.Name,
		Email: therapistDto.Email,
	}
	return therapist, therapistDto.Patients, therapistDto.Files
}

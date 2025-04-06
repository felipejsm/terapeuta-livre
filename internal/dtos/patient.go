package dtos

import (
	"felipejsm/tp-admin/internal/models"
	"time"
)

type PatientDto struct {
	ID                    uint                  `json:"id"`
	Name                  string                `json:"name"`
	Email                 string                `json:"email"`
	TherapistId           int                   `json:"therapist_id"`
	BirthDate             time.Time             `json:"birth_date"`
	Gender                string                `json:"gender"`
	Phone                 string                `json:"phone"`
	CPF                   string                `json:"cpf"`
	RG                    string                `json:"rg"`
	Address               string                `json:"address"`
	EmergencyContactName  string                `json:"emergency_contact_name"`
	EmergencyContactPhone string                `json:"emergency_contact_phone"`
	HealthInsurance       string                `json:"health_insurance"`
	HealthInsuranceNumber string                `json:"health_insurance_number"`
	Active                bool                  `json:"active"`
	CreatedAt             time.Time             `json:"created_at"`
	UpdatedAt             time.Time             `json:"updated_at"`
	Notes                 string                `json:"notes"`
	ProfilePictureKey     string                `json:"profile_picture_key"`
	MaritalStatus         string                `json:"marital_status"`
	Profession            string                `json:"profession"`
	Files                 []models.FileMetadata `json:"files"`
}

func PatientsToDtos(patients []models.Patient) []PatientDto {
	var patientsDtos []PatientDto
	for _, v := range patients {
		patientDto := PatientDto{
			ID:                    v.ID,
			Name:                  v.Name,
			Email:                 v.Email,
			TherapistId:           v.TherapistId,
			BirthDate:             v.BirthDate,
			Gender:                v.Gender,
			Phone:                 v.Phone,
			CPF:                   v.CPF,
			RG:                    v.RG,
			Address:               v.Address,
			EmergencyContactName:  v.EmergencyContactName,
			EmergencyContactPhone: v.EmergencyContactPhone,
			HealthInsurance:       v.HealthInsurance,
			HealthInsuranceNumber: v.HealthInsuranceNumber,
			Active:                v.Active,
			CreatedAt:             v.CreatedAt,
			UpdatedAt:             v.UpdatedAt,
			Notes:                 v.Notes,
			ProfilePictureKey:     v.ProfilePictureKey,
			MaritalStatus:         v.MaritalStatus,
			Profession:            v.Profession,
		}
		patientsDtos = append(patientsDtos, patientDto)
	}
	return patientsDtos
}

func PatientToDto(patient models.Patient, files []models.FileMetadata) *PatientDto {
	return &PatientDto{
		ID:                    patient.ID,
		Name:                  patient.Name,
		Email:                 patient.Email,
		TherapistId:           patient.TherapistId,
		BirthDate:             patient.BirthDate,
		Gender:                patient.Gender,
		Phone:                 patient.Phone,
		CPF:                   patient.CPF,
		RG:                    patient.RG,
		Address:               patient.Address,
		EmergencyContactName:  patient.EmergencyContactName,
		EmergencyContactPhone: patient.EmergencyContactPhone,
		HealthInsurance:       patient.HealthInsurance,
		HealthInsuranceNumber: patient.HealthInsuranceNumber,
		Active:                patient.Active,
		CreatedAt:             patient.CreatedAt,
		UpdatedAt:             patient.UpdatedAt,
		Notes:                 patient.Notes,
		ProfilePictureKey:     patient.ProfilePictureKey,
		MaritalStatus:         patient.MaritalStatus,
		Profession:            patient.Profession,
		Files:                 files,
	}
}

func DtoToPatient(patientDto PatientDto) (models.Patient, []models.FileMetadata) {
	var patient = models.Patient{
		ID:                    patientDto.ID,
		Name:                  patientDto.Name,
		Email:                 patientDto.Email,
		TherapistId:           patientDto.TherapistId,
		BirthDate:             patientDto.BirthDate,
		Gender:                patientDto.Gender,
		Phone:                 patientDto.Phone,
		CPF:                   patientDto.CPF,
		RG:                    patientDto.RG,
		Address:               patientDto.Address,
		EmergencyContactName:  patientDto.EmergencyContactName,
		EmergencyContactPhone: patientDto.EmergencyContactPhone,
		HealthInsurance:       patientDto.HealthInsurance,
		HealthInsuranceNumber: patientDto.HealthInsuranceNumber,
		Active:                patientDto.Active,
		CreatedAt:             patientDto.CreatedAt,
		UpdatedAt:             patientDto.UpdatedAt,
		Notes:                 patientDto.Notes,
		ProfilePictureKey:     patientDto.ProfilePictureKey,
		MaritalStatus:         patientDto.MaritalStatus,
		Profession:            patientDto.Profession,
	}
	return patient, patientDto.Files
}

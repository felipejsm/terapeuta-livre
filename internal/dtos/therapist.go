package dtos

import (
	"felipejsm/tp-admin/internal/models"
	"time"
)

type TherapistDto struct {
	ID                  uint                  `json:"id"`
	Name                string                `json:"name"`
	Email               string                `json:"email"`
	Login               string                `json:"login"`
	CPF                 string                `json:"cpf"`
	Phone               string                `json:"phone"`
	CRP                 string                `json:"crp"`
	Specialization      string                `json:"specialization"`
	Active              bool                  `json:"active"`
	CreatedAt           time.Time             `json:"created_at"`
	UpdatedAt           time.Time             `json:"updated_at"`
	LastLogin           time.Time             `json:"last_login"`
	ProfilePictureKey   string                `json:"profile_picture_key"`
	Timezone            string                `json:"timezone"`
	ReceiveNotifications bool                  `json:"receive_notifications"`
	Patients            []PatientDto          `json:"patients"`
	Files               []models.FileMetadata `json:"files"`
}

func TherapistToDto(therapist models.Therapist, patients []PatientDto, files []models.FileMetadata) *TherapistDto {
	return &TherapistDto{
		ID:                  therapist.ID,
		Name:                therapist.Name,
		Email:               therapist.Email,
		Login:               therapist.Login,
		CPF:                 therapist.CPF,
		Phone:               therapist.Phone,
		CRP:                 therapist.CRP,
		Specialization:      therapist.Specialization,
		Active:              therapist.Active,
		CreatedAt:           therapist.CreatedAt,
		UpdatedAt:           therapist.UpdatedAt,
		LastLogin:           therapist.LastLogin,
		ProfilePictureKey:   therapist.ProfilePictureKey,
		Timezone:            therapist.Timezone,
		ReceiveNotifications: therapist.ReceiveNotifications,
		Patients:            patients,
		Files:               files,
	}
}

func DtoToTherapist(therapistDto TherapistDto) (models.Therapist, []PatientDto, []models.FileMetadata) {
	var therapist = models.Therapist{
		ID:                  therapistDto.ID,
		Name:                therapistDto.Name,
		Email:               therapistDto.Email,
		Login:               therapistDto.Login,
		CPF:                 therapistDto.CPF,
		Phone:               therapistDto.Phone,
		CRP:                 therapistDto.CRP,
		Specialization:      therapistDto.Specialization,
		Active:              therapistDto.Active,
		CreatedAt:           therapistDto.CreatedAt,
		UpdatedAt:           therapistDto.UpdatedAt,
		LastLogin:           therapistDto.LastLogin,
		ProfilePictureKey:   therapistDto.ProfilePictureKey,
		Timezone:            therapistDto.Timezone,
		ReceiveNotifications: therapistDto.ReceiveNotifications,
	}
	return therapist, therapistDto.Patients, therapistDto.Files
}

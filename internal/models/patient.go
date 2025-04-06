package models

import "time"

type Patient struct {
	ID                    uint      `gorm:"primarykey"`
	Name                  string    `gorm:"not null"`
	Email                 string    `gorm:"unique;not null"`
	TherapistId           int       `gorm:"not null"`
	BirthDate             time.Time
	Gender                string
	Phone                 string
	CPF                   string    `gorm:"unique"`
	RG                    string
	Address               string
	EmergencyContactName  string
	EmergencyContactPhone string
	HealthInsurance       string
	HealthInsuranceNumber string
	Active                bool      `gorm:"default:true"`
	CreatedAt             time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt             time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Notes                 string
	ProfilePictureKey     string
	MaritalStatus         string
	Profession            string
}

func (Patient) TableName() string {
	return "tb_patient"
}

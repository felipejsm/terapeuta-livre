package models

import "time"

type Therapist struct {
	ID                  uint      `gorm:"primarykey"`
	Name                string    `gorm:"not null"`
	Email               string    `gorm:"unique;not null"`
	Login               string    `gorm:"unique;not null"`
	Password            string    `gorm:"not null"`
	CPF                 string    `gorm:"unique"`
	Phone               string
	CRP                 string
	Specialization      string
	Active              bool      `gorm:"default:true"`
	CreatedAt           time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	LastLogin           time.Time
	ProfilePictureKey   string
	Timezone            string    `gorm:"default:'America/Sao_Paulo'"`
	ReceiveNotifications bool      `gorm:"default:true"`
}

func (Therapist) TableName() string {
	return "tb_therapist"
}

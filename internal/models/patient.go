package models

type Patient struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Email       string
	TherapistId int
}

func (Patient) TableName() string {
	return "tb_patient"
}

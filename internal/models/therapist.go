package models

type Therapist struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `gorm:not null`
	Email    string `gorm:not null`
	Login    string
	Password string
}

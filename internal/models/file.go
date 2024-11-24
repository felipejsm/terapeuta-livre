package models

import "time"

type File struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	ObjectKey string
	OwnerId   uint
	OwnerType string
}

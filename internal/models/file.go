package models

import "time"

type File struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Size      int64
	CreatedAt time.Time
	ObjectKey string
	OwnerId   uint
	OwnerType string
}

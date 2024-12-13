package models

import "time"

type FileMetadata struct {
	ID        uint `gorm:"primaryKey"`
	FileName  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	ObjectKey string
	Extension string
	OwnerId   int
	FileSize  int
	OwnerType string
}

func (FileMetadata) TableName() string {
	return "tb_file_metadata"
}

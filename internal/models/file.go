package models

type File struct {
	ID         uint `gorm:"primarykey"`
	MetadataId int
	FileData   []byte `gorm:"type:bytea"`
}

func (File) TableName() string {
	return "tb_file"
}

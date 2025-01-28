package repositories

import (
	"felipejsm/tp-admin/internal/models"

	"gorm.io/gorm"
)

type FileRepository struct {
	DB *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{
		DB: db,
	}
}

func (r *FileRepository) DeleteFile(id int) (string, error) {

    result := r.DB.Raw("DELETE FROM tb_file WHERE id = ?", id)
    return "Ok", result.Error
}

func (r *FileRepository) UploadFile(metadataId int, file []byte) (models.File, error) {
	var fileUpload models.File

	result := r.DB.Raw("INSERT INTO tb_file(metadata_id, file_data) VALUES(?,?)", metadataId, file).Scan(&fileUpload)
	return fileUpload, result.Error
}

func (r *FileRepository) DownloadFile(id int) (models.File, error){
    var file models.File

    result := r.DB.Raw("SELECT * FROM tb_file WHERE id = ? ", id).Scan(&file)
    return file, result.Error

}

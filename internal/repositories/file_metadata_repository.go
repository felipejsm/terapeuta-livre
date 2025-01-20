package repositories

import(
	"felipejsm/tp-admin/internal/models"
	"gorm.io/gorm"
)

type FileMetadataRepository struct {
	DB *gorm.DB
}

func NewFileMetadataRepository(db *gorm.DB) *FileMetadataRepository {
	return &FileMetadataRepository{DB: db}
}

func (r *FileMetadataRepository) FindAllByOwnerId(ownerId int) ([]models.FileMetadata, error) {

	var filesMetadata []models.FileMetadata

	result := r.DB.Raw("SELECT * FROM tb_file_metadata WHERE owner_id = ?", ownerId).Scan(&filesMetadata)
	return filesMetadata, result.Error
}


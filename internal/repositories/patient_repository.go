package repositories

import (
	"felipejsm/tp-admin/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type PatientRepository struct {
	DB *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{DB: db}
}

func (r *PatientRepository) FindAllByTherapistId(therapistId int) ([]models.Patient, error) {
	var patients []models.Patient
	result := r.DB.Raw("SELECT * FROM tb_patient WHERE id_therapist = ?", therapistId).Scan(&patients)
	return patients, result.Error
}

func (r *PatientRepository) findFileMetadataByObjectKey(objectKey string) (models.FileMetadata, error) {
	var fileMetadata models.FileMetadata
	result := r.DB.Raw("SELECT * FROM tb_file_metadata WHERE object_key = ?", objectKey).Scan(&fileMetadata)
	return fileMetadata, result.Error
}

func (r *PatientRepository) FindFileByMetadataId(metadataId int) (models.File, error) {
	var file models.File

	result := r.DB.Raw("SELECT * FROM tb_file tf where tf.metadata_id = ?", metadataId).Scan(&file)
	return file, result.Error
}

func (r *PatientRepository) FindByIdAndTherapistId(id int, therapistId int) (models.Patient, error) {
	var patient models.Patient
	fmt.Printf("Before RAW")
	result := r.DB.Raw("SELECT * FROM tb_patient WHERE id = ? AND therapist_id = ?", id, therapistId).Scan(&patient)
	fmt.Printf("Patient %v", patient.Name)
	return patient, result.Error
}

func (r *PatientRepository) FindAllFilesByPatientId(id uint) ([]models.FileMetadata, error) {
	var files []models.FileMetadata
	result := r.DB.Raw("SELECT * FROM tb_file_metadata WHERE owner_id = ?", id).Scan(&files)
	fmt.Printf("Files FULL: %v", files)
	return files, result.Error
}

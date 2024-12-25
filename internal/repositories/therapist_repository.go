package repositories

import (
	"felipejsm/tp-admin/internal/models"

	"gorm.io/gorm"
)

type TherapistRepository struct {
	DB *gorm.DB
}

func NewTherapistRepository(db *gorm.DB) *TherapistRepository {
	return &TherapistRepository{DB: db}
}
func (r *TherapistRepository) FindAllFilesByTherapistId(id int) ([]models.FileMetadata, error) {
	var files []models.FileMetadata
	result := r.DB.Raw("SELECT * FROM tb_file_metadata WHERE owner_id = ?", id).Scan(&files)
	return files, result.Error
}

func (r *TherapistRepository) FindAllPatientsById(id int) ([]models.Patient, error) {
	var patients []models.Patient
	result := r.DB.Raw("SELECT * FROM tb_patient WHERE therapist_id = ?", id).Scan(&patients)
	return patients, result.Error

}

func (r *TherapistRepository) FindById(id int) (models.Therapist, error) {
	var therapist models.Therapist
	result := r.DB.Raw("SELECT * FROM tb_therapist WHERE id = ?", id).Scan(&therapist)
	return therapist, result.Error
}

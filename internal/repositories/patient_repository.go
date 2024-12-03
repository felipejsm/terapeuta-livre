package repositories

import (
	"felipejsm/tp-admin/internal/models"

	"gorm.io/gorm"
)

type PatientRepository struct {
	DB *gorm.DB
}

func (r *PatientRepository) FindAllByTherapistId(therapistId string) ([]models.Patient, error) {
	var patients []models.Patient
	result := r.DB.Raw("SELECT * FROM tb_patient WHERE id_therapist = ?", therapistId).Scan(&patients)
	return patients, result.Error
}

func (r *PatientRepository) FindFileByFileIdAndOwnerId(fileId string, ownerId string) (models.File, error) {
	var file models.File

	result := r.DB.Raw("SELECT * FROM tb_file tf where tf.id = ? AND tf.owner_id = ?", fileId, ownerId).Scan(&file)
	return file, result.Error
}

func (r *PatientRepository) FindById(id string, therapistId string) (models.Patient, error) {
	var patient models.Patient
	result := r.DB.Raw("SELECT * FROM tb_patient WHERE id = ? AND therapist_id = ?", id, therapistId)
	return patient, result.Error
}

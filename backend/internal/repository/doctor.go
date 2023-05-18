package repository

import "backend/internal/models"

type DoctorRepository interface {
	Create(doctor *models.Doctor) error
	GetDoctorByLogin(login string) (*models.Doctor, error)
	GetDoctorById(id uint64) (*models.Doctor, error)
	GetAllDoctors() ([]models.Doctor, error)
	UpdateShedule(id uint64, newStart uint64, newEnd uint64) error
	Delete(id uint64) error
}

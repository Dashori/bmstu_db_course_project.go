package services

import "backend/internal/models"

type DoctorService interface {
	SetRole() error
	Create(client *models.Doctor, password string) (*models.Doctor, error)
	Login(login, password string) (*models.Doctor, error)
	UpdateShedule(id uint64, newStart uint64, newEnd uint64) error
	GetAllDoctors() ([]models.Doctor, error)
	GetDoctorById(id uint64) (*models.Doctor, error)
	GetDoctorByLogin(login string) (*models.Doctor, error)
}

package repository

import "backend/internal/models"

type RecordRepository interface {
	Create(record *models.Record) error
	GetRecord(id uint64) (*models.Record, error)
	GetAllByClient(id uint64) ([]models.Record, error)
	GetAllByDoctor(id uint64) ([]models.Record, error)
	GetAllRecordFilter(doctorId uint64, clientId uint64) ([]models.Record, error)
	GetAllRecords() ([]models.Record, error)
	Delete(id uint64) error
}

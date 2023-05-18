package services

import "backend/internal/models"

type RecordService interface {
	CreateRecord(record *models.Record) error
	DeleteRecord(recordId uint64) error
	GetRecord(recordId uint64) (*models.Record, error)
	GetAllRecords(doctorId uint64, clientId uint64) ([]models.Record, error)
}

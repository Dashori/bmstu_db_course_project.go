package repository

import "backend/internal/models"

type ClientRepository interface {
	SetRole() error
	Create(client *models.Client) error
	GetClientByLogin(login string) (*models.Client, error)
	GetClientById(id uint64) (*models.Client, error)
	GetAllClient() ([]models.Client, error)
	Delete(id uint64) error
}

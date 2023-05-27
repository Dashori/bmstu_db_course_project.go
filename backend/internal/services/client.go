package services

import "backend/internal/models"

type ClientService interface {
	SetRole() error
	Create(client *models.Client, password string) (*models.Client, error)
	Login(login, password string) (*models.Client, error)
	GetClientById(id uint64) (*models.Client, error)
	GetClientByLogin(login string) (*models.Client, error)
}

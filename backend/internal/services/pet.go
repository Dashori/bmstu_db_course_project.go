package services

import "backend/internal/models"

type PetService interface {
	Create(pet *models.Pet, login string) error
	Delete(petId uint64, clientId uint64) error
	GetAllByClient(id uint64) ([]models.Pet, error)
	GetPet(petId uint64) (*models.Pet, error)
}

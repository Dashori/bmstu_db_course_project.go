package repository

import "backend/internal/models"

type PetRepository interface {
	Create(pet *models.Pet) error
	GetPet(id uint64) (*models.Pet, error)
	GetAllByClient(id uint64) ([]models.Pet, error)
	GetAllPets() ([]models.Pet, error)
	Delete(id uint64) error
}

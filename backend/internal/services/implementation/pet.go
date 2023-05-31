package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/repoErrors"
	"backend/internal/pkg/errors/servicesErrors"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
)

type PetServiceImplementation struct {
	PetRepository    repository.PetRepository
	ClientRepository repository.ClientRepository
	logger           *log.Logger
}

func NewPetServiceImplementation(
	PetRepository repository.PetRepository,
	ClientRepository repository.ClientRepository,
	logger *log.Logger) services.PetService {

	return &PetServiceImplementation{
		PetRepository:    PetRepository,
		ClientRepository: ClientRepository,
		logger:           logger,
	}
}

func (p *PetServiceImplementation) Create(pet *models.Pet, login string) error { // login client
	p.logger.Debug("PET! Start create new pet", "client login", login, "petName", pet.Name)

	client, err := p.ClientRepository.GetClientByLogin(login)
	if err != nil {
		p.logger.Warn("PET! Error with client when create new pet", "client login", login, "petName", pet.Name, "err", err)
		return serviceErrors.ClientDoesNotExists
	}

	pets, err := p.PetRepository.GetAllByClient(client.ClientId)
	if err != nil && err != repoErrors.EntityDoesNotExists {
		p.logger.Warn("PET! Error in repository method GetAllByClient", "client login", login, "err", err)
		return err
	}

	// у клиента/доктора login считается уникальным, а у животного уникальным считать name нельзя,
	// так как наверняка найдется 500 Барсиков. Поэтому имя уникально для каждого владельца
	// следовательно, нужно проверять связку name + clientId
	// для этого получаем всех питомцев клиента и сравниваем клички
	for _, j := range pets {
		if j.Name == pet.Name {
			p.logger.Warn("PET! Pet already exists", "client login", login, "petName", pet.Name)
			return serviceErrors.PetAlreadyExists
		}
	}

	pet.ClientId = client.ClientId
	err = p.PetRepository.Create(pet)
	if err != nil {
		p.logger.Warn("PET! Error in repository method Create", "client login", login, "petName", pet.Name, "petName", pet.Name, "err", err)
	}

	p.logger.Info("PET! Successfully Create new pet", "client login", login, "petName", pet.Name)
	return nil
}

func (p *PetServiceImplementation) Delete(petId uint64, clientId uint64) error {
	pet, err := p.PetRepository.GetPet(petId)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		p.logger.Warn("PET! Pet does not exists", "petId", petId)
		return serviceErrors.PetDoesNotExists
	} else if err != nil {
		p.logger.Warn("PET! Error in repository method GetPet", "petId", petId, "err", err)
		return err
	}

	if pet.ClientId != clientId {
		p.logger.Warn("PET! Error, client is not pet owner and can't delete", "cliendId", clientId, "petId", petId, "petName", pet.Name)
		return serviceErrors.NotClientPet
	}

	err = p.PetRepository.Delete(petId)

	if err != nil {
		p.logger.Warn("PET! Error in repository method delete", "cliendId", clientId, "petId", petId, "petName", pet.Name, "err", err)
		return err
	}

	p.logger.Info("PET! Successfully delete pet", "cliendId", clientId, "petId", petId, "petName", pet.Name)
	return nil
}

func (p *PetServiceImplementation) GetPet(petId uint64) (*models.Pet, error) {
	pet, err := p.PetRepository.GetPet(petId)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		p.logger.Warn("PET! Pet does not exists", "petId", petId)
		return nil, serviceErrors.PetDoesNotExists
	} else if err != nil {
		p.logger.Warn("PET! Error in repository method GetPet", "petId", petId, "err", err)
		return nil, err
	}

	p.logger.Debug("PET! Successfully repository method GetPet", "petId", petId, "petName", pet.Name)
	return pet, nil
}

func (p *PetServiceImplementation) GetAllByClient(id uint64) ([]models.Pet, error) {
	pets, err := p.PetRepository.GetAllByClient(id)

	if err != nil {
		p.logger.Warn("PET! Error in repository method GetAllByClient", "id", id, "err", err)
		return nil, err
	}

	p.logger.Info("PET! Successfully repository method GetAllByClient", "clientId", id)
	return pets, nil
}

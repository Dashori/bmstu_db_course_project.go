package postgres_repo

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/dbErrors"
	"backend/internal/pkg/errors/repoErrors"
	"backend/internal/repository"
	"database/sql"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
)

type PetPostgres struct {
	PetId    uint64 `db:"id_pet"`
	Name     string `db:"name"`
	Type     string `db:"type"`
	Age      uint64 `db:"age"`
	Health   uint64 `db:"health"`
	ClientId uint64 `db:"id_client"`
}

type PetPostgresRepository struct {
	db *sqlx.DB
}

func NewPetPostgresRepository(db *sqlx.DB) repository.PetRepository {
	return &PetPostgresRepository{db: db}
}

func (p *PetPostgresRepository) Create(pet *models.Pet) error {
	query := `insert into pets(name, type, age, health, id_client)
	values ($1, $2, $3, $4, $5);`

	_, err := p.db.Exec(query, pet.Name, pet.Type, pet.Age, pet.Health, pet.ClientId)

	if err != nil {
		return dbErrors.ErrorInsert
	}

	return nil
}

func (p *PetPostgresRepository) GetPet(id uint64) (*models.Pet, error) {
	query := `select * from pets where pets.id_pet = $1;`
	petBD := &PetPostgres{}

	err := p.db.Get(petBD, query, id)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, dbErrors.ErrorSelect
	}

	petModels := &models.Pet{}
	err = copier.Copy(petModels, petBD)

	if err != nil {
		return nil, dbErrors.ErrorCopy
	}

	return petModels, nil
}

func (p *PetPostgresRepository) GetAllByClient(id uint64) ([]models.Pet, error) {
	query := `select * from pets where id_client = $1;`
	var petPostgres []PetPostgres

	err := p.db.Select(&petPostgres, query, id)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, dbErrors.ErrorSelect
	}

	petModels := []models.Pet{}

	for _, r := range petPostgres {
		pet := &models.Pet{}
		err = copier.Copy(pet, r)

		if err != nil {
			return nil, dbErrors.ErrorCopy
		}

		petModels = append(petModels, *pet)
	}

	return petModels, nil
}

func (p *PetPostgresRepository) GetAllPets() ([]models.Pet, error) {
	query := `select * from pets;`

	var petPostgres []PetPostgres

	err := p.db.Select(&petPostgres, query)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, dbErrors.ErrorSelect
	}

	petModels := []models.Pet{}

	for i := range petPostgres {
		pet := &models.Pet{}
		err = copier.Copy(pet, &petPostgres[i])

		if err != nil {
			return nil, dbErrors.ErrorCopy
		}

		petModels = append(petModels, *pet)
	}

	return petModels, nil
}

func (p *PetPostgresRepository) Delete(id uint64) error {
	query := `delete from pets where id_pet = $1`

	_, err := p.db.Exec(query, id)

	if err != nil {
		return dbErrors.ErrorDelete
	}

	return nil
}

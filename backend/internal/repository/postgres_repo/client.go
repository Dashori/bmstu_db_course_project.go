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

type ClientPostgres struct {
	ClientId uint64 `db:"id_client"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

type ClientPostgresRepository struct {
	db *sqlx.DB
}

func NewClientPostgresRepository(db *sqlx.DB) repository.ClientRepository {
	return &ClientPostgresRepository{db: db}
}

func (c *ClientPostgresRepository) Create(client *models.Client) error {
	query := `insert into clients(login, password) values($1, $2);`

	_, err := c.db.Exec(query, client.Login, client.Password)

	if err != nil {
		return dbErrors.ErrorInsert
	}

	return nil
}

func (c *ClientPostgresRepository) GetClientByLogin(login string) (*models.Client, error) {
	query := `select * from clients where login = $1;`
	clientDB := &ClientPostgres{}

	err := c.db.Get(clientDB, query, login)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, dbErrors.ErrorSelect
	}

	clientModels := &models.Client{}
	err = copier.Copy(clientModels, clientDB)

	if err != nil {
		return nil, dbErrors.ErrorCopy
	}

	return clientModels, nil
}

func (c *ClientPostgresRepository) GetClientById(id uint64) (*models.Client, error) {
	query := `select * from clients where id_client = $1;`
	clientDB := &ClientPostgres{}

	err := c.db.Get(clientDB, query, id)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, dbErrors.ErrorSelect
	}

	clientModels := &models.Client{}
	err = copier.Copy(clientModels, clientDB)

	if err != nil {
		return nil, dbErrors.ErrorCopy
	}

	return clientModels, nil
}

func (c *ClientPostgresRepository) GetAllClient() ([]models.Client, error) {
	query := `select * from clients;`
	clientDB := []ClientPostgres{}

	err := c.db.Select(&clientDB, query)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, dbErrors.ErrorSelect
	}

	clientModels := []models.Client{}

	for i := range clientDB {
		client := models.Client{}
		err = copier.Copy(client, &clientDB[i])

		if err != nil {
			return nil, dbErrors.ErrorCopy
		}

		clientModels = append(clientModels, client)
	}

	return clientModels, nil
}

func (c *ClientPostgresRepository) Delete(id uint64) error {
	query := `delete from clients where id_client = $1`

	_, err := c.db.Exec(query, id)
	if err != nil {
		return dbErrors.ErrorDelete
	}

	return nil
}

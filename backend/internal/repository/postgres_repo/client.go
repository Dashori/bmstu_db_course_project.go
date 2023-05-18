package postgres_repo

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/bdErrors"
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
		return bdErrors.ErrorInsert
	}

	return nil
}

func (c *ClientPostgresRepository) GetClientByLogin(login string) (*models.Client, error) {
	query := `select * from clients where login = $1;`
	clientBD := &ClientPostgres{}

	err := c.db.Get(clientBD, query, login)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	clientModels := &models.Client{}
	err = copier.Copy(clientModels, clientBD)

	if err != nil {
		return nil, bdErrors.ErrorCopy
	}

	return clientModels, nil
}

func (c *ClientPostgresRepository) GetClientById(id uint64) (*models.Client, error) {
	query := `select * from clients where id_client = $1;`
	clientBD := &ClientPostgres{}

	err := c.db.Get(clientBD, query, id)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	clientModels := &models.Client{}
	err = copier.Copy(clientModels, clientBD)

	if err != nil {
		return nil, bdErrors.ErrorCopy
	}

	return clientModels, nil
}

func (c *ClientPostgresRepository) GetAllClient() ([]models.Client, error) {
	query := `select * from clients;`
	clientBD := []ClientPostgres{}

	err := c.db.Select(&clientBD, query)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	clientModels := []models.Client{}

	for i := range clientBD {
		client := models.Client{}
		err = copier.Copy(client, &clientBD[i])

		if err != nil {
			return nil, bdErrors.ErrorCopy
		}

		clientModels = append(clientModels, client)
	}

	return clientModels, nil
}

func (c *ClientPostgresRepository) Delete(id uint64) error {
	query := `delete from clients where id_client = $1`

	_, err := c.db.Exec(query, id)
	if err != nil {
		return bdErrors.ErrorDelete
	}

	return nil
}

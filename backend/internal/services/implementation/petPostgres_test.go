package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/repository/postgres_repo"
	"backend/internal/services"
	"context"
	"database/sql"
	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"os"
	"testing"
)

type petServiceFieldsPostgres struct {
	petRepository    *repository.PetRepository
	clientRepository *repository.ClientRepository
	logger           *log.Logger
}

func createPetServiceFieldsPostgres(dbTest *sql.DB) *petServiceFieldsPostgres {
	fields := new(petServiceFieldsPostgres)

	repositoryFields := postgres_repo.PostgresRepositoryFields{DB: dbTest}

	petRepo := postgres_repo.CreatePetPostgresRepository(&repositoryFields)
	fields.petRepository = &petRepo

	clientRepo := postgres_repo.CreateClientPostgresRepository(&repositoryFields)
	fields.clientRepository = &clientRepo

	fields.logger = log.New(os.Stderr)
	fields.logger.SetLevel(log.FatalLevel)

	return fields
}

func createPetServicePostgres(fields *petServiceFieldsPostgres) services.PetService {
	return NewPetServiceImplementation(*fields.petRepository, *fields.clientRepository, fields.logger)
}

var testPetCreatePostgresSuccess = []struct {
	TestName        string
	InputData       struct{}
	Prepare         func(fields *petServiceFieldsPostgres)
	CheckOutput     func(t *testing.T, err error)
	CheckOutputHelp func(t *testing.T, err error)
}{
	{
		TestName:  "pet create and delete success",
		InputData: struct{}{},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
		CheckOutputHelp: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testPetCreatePostgresFailure = []struct {
	TestName  string
	InputData struct {
		pet   *models.Pet
		login string
	}
	Prepare     func(fields *petServiceFieldsPostgres)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "pet creare failure",
		InputData: struct {
			pet   *models.Pet
			login string
		}{pet: &models.Pet{Name: "Havrosha", ClientId: 1}, login: "Ffdpfpsgf"},
		CheckOutput: func(t *testing.T, err error) {
			require.Error(t, err)
		},
	},
}

func TestPetServiceImplementationCreatePostgres(t *testing.T) {

	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	for _, tt := range testPetCreatePostgresSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields := createPetServiceFieldsPostgres(db)

			clients := fields.clientRepository
			pets := fields.petRepository

			err := (*clients).Create(&models.Client{Login: "ChicagoTest", Password: "12345"})
			tt.CheckOutputHelp(t, err)

			client, err := (*clients).GetClientByLogin("ChicagoTest")
			tt.CheckOutputHelp(t, err)

			err = (*pets).Create(&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: client.ClientId})
			tt.CheckOutput(t, err)

			// трюк чтоб узнать id питомца Havrosha
			clientPets, err := (*pets).GetAllByClient(client.ClientId)
			tt.CheckOutputHelp(t, err)
			petId := clientPets[0].PetId

			err = (*pets).Delete(petId)
			tt.CheckOutputHelp(t, err)

			err = (*clients).Delete(client.ClientId)
			tt.CheckOutputHelp(t, err)
		})
	}

	for _, tt := range testPetCreatePostgresFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields := createPetServiceFieldsPostgres(db)

			pet := createPetServicePostgres(fields)

			err := pet.Create(tt.InputData.pet, tt.InputData.login)

			tt.CheckOutput(t, err)
		})
	}
}

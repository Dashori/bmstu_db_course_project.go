package postgres_repo

import (
	"backend/internal/models"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"testing"
)

var testPetPostgresRepositoryCreateSuccess = []struct {
	TestName  string
	InputData struct {
		pet *models.Pet
	}
	CheckOutput     func(t *testing.T, err error)
	CheckOutputHelp func(t *testing.T, err error)
}{
	{
		TestName: "create success test",
		InputData: struct {
			pet *models.Pet
		}{&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: 5}},

		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
		CheckOutputHelp: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestPetPostgresRepositoryCreate(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	for _, tt := range testPetPostgresRepositoryCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields := PostgresRepositoryFields{DB: db}
			// fields, err := CreatePostgresRepositoryFieldsTest(configFileName, pathToConfig)

			clientRepository := CreateClientPostgresRepository(&fields)
			petRepository := CreatePetPostgresRepository(&fields)

			err := clientRepository.Create(&models.Client{Login: "ChicagoTest", Password: "12345"})
			tt.CheckOutputHelp(t, err)

			client, err := clientRepository.GetClientByLogin("ChicagoTest")
			tt.CheckOutputHelp(t, err)

			err = petRepository.Create(&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: client.ClientId})
			tt.CheckOutputHelp(t, err)

			// трюк чтоб узнать id питомца Havrosha
			pets, err := petRepository.GetAllByClient(client.ClientId)
			tt.CheckOutputHelp(t, err)

			err = petRepository.Delete(pets[0].PetId)
			tt.CheckOutputHelp(t, err)

			err = clientRepository.Delete(client.ClientId)
			tt.CheckOutputHelp(t, err)
		})
	}
}

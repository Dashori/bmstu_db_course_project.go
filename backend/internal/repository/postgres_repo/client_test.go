package postgres_repo

import (
	"backend/internal/models"
	"github.com/stretchr/testify/require"
	"testing"
)

var testClientPostgresRepositoryCreateSuccess = []struct {
	TestName  string
	InputData struct {
		client *models.Client
	}
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "create success test",
		InputData: struct {
			client *models.Client
		}{&models.Client{Login: "ChicagoTest", Password: "12345"}},

		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testClientPostgresRepositoryCreateFailure = []struct {
	TestName  string
	InputData struct {
		client *models.Client
	}
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "create failure test",
		InputData: struct {
			client *models.Client
		}{&models.Client{Login: "ChicagoTest", Password: "12345"}},

		CheckOutput: func(t *testing.T, err error) {
			require.Error(t, err)
		},
	},
}

func TestClientPostgresRepositoryCreate(t *testing.T) {
	for _, tt := range testClientPostgresRepositoryCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields, err := CreatePostgresRepositoryFieldsTest(configFileName, pathToConfig)

			clientRepository := CreateClientPostgresRepository(fields)

			err = clientRepository.Create(tt.InputData.client)

			tt.CheckOutput(t, err)

			client, err := clientRepository.GetClientByLogin("ChicagoTest")

			if err == nil {
				err = clientRepository.Delete(client.ClientId)
			}

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testClientPostgresRepositoryCreateFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields, err := CreatePostgresRepositoryFieldsTest(configFileName, pathToConfig)

			clientRepository := CreateClientPostgresRepository(fields)

			clientRepository.Create(tt.InputData.client)

			clientRepository.Create(tt.InputData.client)

			err = clientRepository.Create(tt.InputData.client)

			tt.CheckOutput(t, err)

			client, err := clientRepository.GetClientByLogin("ChicagoTest")

			if err == nil {
				err = clientRepository.Delete(client.ClientId)
			}
		})
	}
}
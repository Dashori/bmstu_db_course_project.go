package postgres_repo

import (
	"backend/internal/models"
	"github.com/stretchr/testify/require"
	"testing"
)

var testPetPostgresRepositoryCreateSuccess = []struct {
	TestName  string
	InputData struct {
		pet *models.Pet
	}
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "create success test",
		InputData: struct {
			pet *models.Pet
		}{&models.Pet{PetId: uint64(5), Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: 5}},

		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestPetPostgresRepositoryCreate(t *testing.T) {
	for _, tt := range testPetPostgresRepositoryCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields, err := CreatePostgresRepositoryFieldsTest(configFileName, pathToConfig)

			clientRepository := CreateClientPostgresRepository(fields)
			err = clientRepository.Create(&models.Client{ClientId: uint64(5), Login: "Chicago", Password: "12345"})

			if err == nil {
				petRepository := CreatePetPostgresRepository(fields)

				err = petRepository.Create(tt.InputData.pet)

				tt.CheckOutput(t, err)

				petRepository.Delete(uint64(5))

				clientRepository.Delete(uint64(5))
			}
		})
	}
}

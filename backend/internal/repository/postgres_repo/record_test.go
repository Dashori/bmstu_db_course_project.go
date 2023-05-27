package postgres_repo

import (
	"backend/internal/models"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"testing"
	"time"
)

var testRecordPostgresRepositoryCreateSuccess = []struct {
	TestName  string
	InputData struct {
	}
	CheckOutput     func(t *testing.T, err error)
	CheckOutputHelp func(t *testing.T, err error)
}{
	{
		TestName: "Create test",
		InputData: struct {
		}{},

		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
		CheckOutputHelp: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestRecordPostgresRepositoryGetDoctor(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	for _, tt := range testRecordPostgresRepositoryCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields := PostgresRepositoryFields{DB: db}

			recordRepository := CreateRecordPostgresRepository(&fields)
			clientRepository := CreateClientPostgresRepository(&fields)
			doctorRepository := CreateDoctorPostgresRepository(&fields)
			petRepository := CreatePetPostgresRepository(&fields)

			err := clientRepository.Create(&models.Client{Login: "ChicagoTest", Password: "12345"})
			tt.CheckOutputHelp(t, err)

			client, err := clientRepository.GetClientByLogin("ChicagoTest")
			tt.CheckOutputHelp(t, err)

			err = petRepository.Create(&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: client.ClientId})
			tt.CheckOutputHelp(t, err)

			err = doctorRepository.Create(&models.Doctor{Login: "ChicagoTest", Password: "12345"})
			tt.CheckOutputHelp(t, err)

			doctor, err := doctorRepository.GetDoctorByLogin("ChicagoTest")
			tt.CheckOutputHelp(t, err)

			// трюк чтоб узнать id питомца Havrosha
			pets, err := petRepository.GetAllByClient(client.ClientId)
			tt.CheckOutputHelp(t, err)
			petId := pets[0].PetId

			err = recordRepository.Create(&models.Record{
				PetId: petId, ClientId: client.ClientId, DoctorId: doctor.DoctorId,
				DatetimeStart: time.Date(3000, 8, 15, 14, 00, 00, 00, time.UTC),
				DatetimeEnd:   time.Date(3000, 8, 15, 14, 00, 00, 00, time.UTC)})

			tt.CheckOutput(t, err)

			err = petRepository.Delete(pets[0].PetId) // при удалении pet удалится и запись в records
			tt.CheckOutputHelp(t, err)
		})
	}
}

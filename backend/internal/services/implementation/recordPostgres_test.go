package servicesImplementation

import (
	"backend/internal/models"
	// "backend/internal/repository"

	// "backend/internal/services"
	"context"

	// "github.com/charmbracelet/log"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"

	"testing"
	"time"
)

var testRecordCreatePostgresSuccess = []struct {
	TestName        string
	InputData       struct{}
	Prepare         func(fields *RecordServiceFieldsPostgres)
	CheckOutput     func(t *testing.T, err error)
	CheckOutputHelp func(t *testing.T, err error)
}{
	{
		TestName:  "record creare and delete success",
		InputData: struct{}{},

		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},

		CheckOutputHelp: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testRecordCreatePostgresFailure = []struct {
	TestName        string
	InputData       struct{}
	Prepare         func(fields *recordServiceFields)
	CheckOutput     func(t *testing.T, err error)
	CheckOutputHelp func(t *testing.T, err error)
}{
	{
		TestName:  "record create failure",
		InputData: struct{}{},

		CheckOutput: func(t *testing.T, err error) {
			require.Error(t, err)
		},
		CheckOutputHelp: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestRecordServiceImplementationCreatePostgres(t *testing.T) {

	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	for _, tt := range testRecordCreatePostgresSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields := СreateRecordServiceFieldsPostgres(db)
			records := CreateRecordServicePostgres(fields)

			clients := fields.ClientRepository
			doctors := fields.DoctorRepository
			pets := fields.PetRepository

			err := (*clients).Create(&models.Client{Login: "ChicagoTest", Password: "12345"})
			tt.CheckOutputHelp(t, err)

			client, err := (*clients).GetClientByLogin("ChicagoTest")
			tt.CheckOutputHelp(t, err)

			err = (*pets).Create(&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: client.ClientId})
			tt.CheckOutputHelp(t, err)

			err = (*doctors).Create(&models.Doctor{Login: "ChicagoTest", Password: "12345", StartTime: 10, EndTime: 23})
			tt.CheckOutputHelp(t, err)

			doctor, err := (*doctors).GetDoctorByLogin("ChicagoTest")
			tt.CheckOutputHelp(t, err)

			// трюк чтоб узнать id питомца Havrosha
			clientPets, err := (*pets).GetAllByClient(client.ClientId)
			tt.CheckOutputHelp(t, err)
			petId := clientPets[0].PetId

			err = records.CreateRecord(&models.Record{
				PetId: petId, ClientId: client.ClientId, DoctorId: doctor.DoctorId,
				DatetimeStart: time.Date(3000, 8, 15, 14, 00, 00, 00, time.UTC),
				DatetimeEnd:   time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC)})

			tt.CheckOutput(t, err)

			err = (*pets).Delete(petId) // при удалении pet удалится и запись в records
			tt.CheckOutputHelp(t, err)

			err = (*doctors).Delete(doctor.DoctorId)
			tt.CheckOutputHelp(t, err)

			err = (*clients).Delete(client.ClientId)
			tt.CheckOutputHelp(t, err)
		})
	}

	for _, tt := range testRecordCreatePostgresFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields := СreateRecordServiceFieldsPostgres(db)
			records := CreateRecordServicePostgres(fields)

			clients := fields.ClientRepository
			doctors := fields.DoctorRepository
			pets := fields.PetRepository

			err := (*clients).Create(&models.Client{Login: "ChicagoTest", Password: "12345"})
			tt.CheckOutputHelp(t, err)

			client, err := (*clients).GetClientByLogin("ChicagoTest")
			tt.CheckOutputHelp(t, err)

			err = (*pets).Create(&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: client.ClientId})
			tt.CheckOutputHelp(t, err)

			err = (*doctors).Create(&models.Doctor{Login: "ChicagoTest", Password: "12345", StartTime: 10, EndTime: 23})
			tt.CheckOutputHelp(t, err)

			doctor, err := (*doctors).GetDoctorByLogin("ChicagoTest")
			tt.CheckOutputHelp(t, err)

			// трюк чтоб узнать id питомца Havrosha
			clientPets, err := (*pets).GetAllByClient(client.ClientId)
			tt.CheckOutputHelp(t, err)
			petId := clientPets[0].PetId

			err = records.CreateRecord(&models.Record{
				PetId: petId, ClientId: client.ClientId, DoctorId: doctor.DoctorId,
				DatetimeStart: time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC),
				DatetimeEnd:   time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC)})

			tt.CheckOutput(t, err)

			err = (*pets).Delete(petId) // при удалении pet удалится и запись в records
			tt.CheckOutputHelp(t, err)

			err = (*doctors).Delete(doctor.DoctorId)
			tt.CheckOutputHelp(t, err)

			err = (*clients).Delete(client.ClientId)
			tt.CheckOutputHelp(t, err)
		})
	}
}

package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/repository/postgres_repo"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/require"
	"os"
	"database/sql"
	"testing"
	// "fmt"
	"context"
	"time"
	"github.com/testcontainers/testcontainers-go"
)
	// "github.com/testcontainers/testcontainers-go"
    // "github.com/testcontainers/testcontainers-go/wait")

type recordServiceFieldsPostgres struct {
	recordRepository *repository.RecordRepository
	doctorRepository *repository.DoctorRepository
	clientRepository *repository.ClientRepository
	petRepository    *repository.PetRepository
	logger           *log.Logger
}

func createRecordServiceFieldsPostgres(dbTest *sql.DB) *recordServiceFieldsPostgres {
	fields := new(recordServiceFieldsPostgres)

	// repositoryFields, err := postgres_repo.CreatePostgresRepositoryFieldsTest(configFileName, pathToConfig)

	// if err != nil {
	// 	return nil
	// }

	repositoryFields := postgres_repo.PostgresRepositoryFields{DB : dbTest}

	recordRepo := postgres_repo.CreateRecordPostgresRepository(&repositoryFields)
	fields.recordRepository = &recordRepo

	doctorRepo := postgres_repo.CreateDoctorPostgresRepository(&repositoryFields)
	fields.doctorRepository = &doctorRepo

	clientRepo := postgres_repo.CreateClientPostgresRepository(&repositoryFields)
	fields.clientRepository = &clientRepo

	petRepo := postgres_repo.CreatePetPostgresRepository(&repositoryFields)
	fields.petRepository = &petRepo

	fields.logger = log.New(os.Stderr)
	fields.logger.SetLevel(log.FatalLevel)

	return fields
}

func createRecordServicePostgres(fields *recordServiceFieldsPostgres) services.RecordService {
	return NewRecordServiceImplementation(*fields.recordRepository, *fields.doctorRepository,
		*fields.clientRepository, *fields.petRepository, fields.logger)
}

var testRecordCreatePostgresSuccess = []struct {
	TestName        string
	InputData       struct{}
	Prepare         func(fields *recordServiceFieldsPostgres)
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
			fields := createRecordServiceFieldsPostgres(db)
			records := createRecordServicePostgres(fields)

			clients := fields.clientRepository
			doctors := fields.doctorRepository
			pets := fields.petRepository

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
			fields := createRecordServiceFieldsPostgres(db)
			records := createRecordServicePostgres(fields)

			clients := fields.clientRepository
			doctors := fields.doctorRepository
			pets := fields.petRepository

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

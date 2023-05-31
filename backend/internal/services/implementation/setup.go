package servicesImplementation

import (
	"backend/internal/repository"
	"backend/internal/repository/postgres_repo"
	"context"
	"database/sql"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	// "database/sql"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	"os"
)

const (
	USER     = "dashori"
	PASSWORD = "parasha"
	DBNAME   = "postgres"
)

func SetupTestDatabase() (testcontainers.Container, *sql.DB) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       DBNAME,
			"POSTGRES_PASSWORD": PASSWORD,
			"POSTGRES_USER":     USER,
		},
	}

	dbContainer, _ := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")

	dsnPGConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port.Int(), USER, PASSWORD, DBNAME)
	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		return dbContainer, nil
	}

	err = db.Ping()
	if err != nil {
		return dbContainer, nil
	}
	db.SetMaxOpenConns(10)

	text, err := os.ReadFile("../../../db/postgreSQL/init.sql")
	if err != nil {
		return dbContainer, nil
	}

	if _, err := db.Exec(string(text)); err != nil {
		return dbContainer, nil
	}

	fmt.Println(string(text))

	return dbContainer, db
}

func СreateRecordServiceFieldsPostgres(dbTest *sql.DB) *RecordServiceFieldsPostgres {
	fields := new(RecordServiceFieldsPostgres)

	repositoryFields := postgres_repo.PostgresRepositoryFields{DB: dbTest}

	recordRepo := postgres_repo.CreateRecordPostgresRepository(&repositoryFields)
	fields.RecordRepository = &recordRepo

	doctorRepo := postgres_repo.CreateDoctorPostgresRepository(&repositoryFields)
	fields.DoctorRepository = &doctorRepo

	clientRepo := postgres_repo.CreateClientPostgresRepository(&repositoryFields)
	fields.ClientRepository = &clientRepo

	petRepo := postgres_repo.CreatePetPostgresRepository(&repositoryFields)
	fields.PetRepository = &petRepo

	fields.logger = log.New(os.Stderr)
	fields.logger.SetLevel(log.FatalLevel)

	return fields
}

type RecordServiceFieldsPostgres struct {
	RecordRepository *repository.RecordRepository
	DoctorRepository *repository.DoctorRepository
	ClientRepository *repository.ClientRepository
	PetRepository    *repository.PetRepository
	logger           *log.Logger
}

func CreateRecordServicePostgres(fields *RecordServiceFieldsPostgres) services.RecordService {
	return NewRecordServiceImplementation(*fields.RecordRepository, *fields.DoctorRepository,
		*fields.ClientRepository, *fields.PetRepository, fields.logger)
}

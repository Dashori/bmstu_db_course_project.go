package postgres_repo

import (
	"backend/cmd/modes/flags"
	"backend/config"
	"backend/internal/pkg/errors/bdErrors"
	"backend/internal/repository"
	"database/sql"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"os"
)

type PostgresRepositoryFields struct {
	DB     *sql.DB
	Config config.Config
}

func CreatePostgresRepositoryFields(Postgres flags.PostgresFlags, logger *log.Logger) (*PostgresRepositoryFields, error) {
	fields := new(PostgresRepositoryFields)
	var err error

	fields.Config.Postgres = Postgres

	fields.DB, err = fields.Config.Postgres.InitDB(logger)
	if err != nil {
		logger.Error("POSTGRES! Error parse config for postgreSQL")
		return nil, bdErrors.ErrorParseConfig
	}

	logger.Info("POSTGRES! Successfully create postgres repository fields")

	return fields, nil
}

func CreatePostgresRepositoryFieldsTest(configFileName string, pathToConfig string) (*PostgresRepositoryFields, error) {
	fields := new(PostgresRepositoryFields)

	err := fields.Config.ParseConfig(configFileName, pathToConfig)
	if err != nil {
		return nil, bdErrors.ErrorParseConfig
	}

	fields.DB, err = fields.Config.Postgres.InitDB(log.New(os.Stderr))
	if err != nil {
		return nil, bdErrors.ErrorParseConfig
	}

	return fields, nil
}

func CreateClientPostgresRepository(fields *PostgresRepositoryFields) repository.ClientRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewClientPostgresRepository(dbx)
}

func CreateDoctorPostgresRepository(fields *PostgresRepositoryFields) repository.DoctorRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewDoctorPostgresRepository(dbx)
}

func CreatePetPostgresRepository(fields *PostgresRepositoryFields) repository.PetRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewPetPostgresRepository(dbx)
}

func CreateRecordPostgresRepository(fields *PostgresRepositoryFields) repository.RecordRepository {
	dbx := sqlx.NewDb(fields.DB, "pgx")

	return NewRecordPostgresRepository(dbx)
}

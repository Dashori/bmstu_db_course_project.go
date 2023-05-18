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
	db     *sql.DB
	config config.Config
}

func CreatePostgresRepositoryFields(Postgres flags.PostgresFlags, logger *log.Logger) (*PostgresRepositoryFields, error) {
	fields := new(PostgresRepositoryFields)
	var err error

	fields.config.Postgres = Postgres

	fields.db, err = fields.config.Postgres.InitDB(logger)
	if err != nil {
		logger.Error("POSTGRES! Error parse config for postgreSQL")
		return nil, bdErrors.ErrorParseConfig
	}

	logger.Info("POSTGRES! Successfully create postgres repository fields")

	return fields, nil
}

func CreatePostgresRepositoryFieldsTest(configFileName string, pathToConfig string) (*PostgresRepositoryFields, error) {
	fields := new(PostgresRepositoryFields)

	err := fields.config.ParseConfig(configFileName, pathToConfig)
	if err != nil {
		return nil, bdErrors.ErrorParseConfig
	}

	fields.db, err = fields.config.Postgres.InitDB(log.New(os.Stderr))
	if err != nil {
		return nil, bdErrors.ErrorParseConfig
	}

	return fields, nil
}

func CreateClientPostgresRepository(fields *PostgresRepositoryFields) repository.ClientRepository {
	dbx := sqlx.NewDb(fields.db, "pgx")

	return NewClientPostgresRepository(dbx)
}

func CreateDoctorPostgresRepository(fields *PostgresRepositoryFields) repository.DoctorRepository {
	dbx := sqlx.NewDb(fields.db, "pgx")

	return NewDoctorPostgresRepository(dbx)
}

func CreatePetPostgresRepository(fields *PostgresRepositoryFields) repository.PetRepository {
	dbx := sqlx.NewDb(fields.db, "pgx")

	return NewPetPostgresRepository(dbx)
}

func CreateRecordPostgresRepository(fields *PostgresRepositoryFields) repository.RecordRepository {
	dbx := sqlx.NewDb(fields.db, "pgx")

	return NewRecordPostgresRepository(dbx)
}

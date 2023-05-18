package postgres_repo

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/bdErrors"
	"backend/internal/pkg/errors/repoErrors"
	"backend/internal/repository"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

type RecordPostgres struct {
	RecordId      uint64    `db:"id_record"`
	PetId         uint64    `db:"id_pet"`
	ClientId      uint64    `db:"id_client"`
	DoctorId      uint64    `db:"id_doctor"`
	DatetimeStart time.Time `db:"time_start"`
	DatetimeEnd   time.Time `db:"time_end"`
}

type RecordPostgresRepository struct {
	db *sqlx.DB
}

func NewRecordPostgresRepository(db *sqlx.DB) repository.RecordRepository {
	return &RecordPostgresRepository{db: db}
}

func (r *RecordPostgresRepository) Create(record *models.Record) error {
	query := `insert into records (id_pet, id_client, id_doctor, time_start, time_end)
	values ($1, $2, $3, $4, $5);`

	_, err := r.db.Exec(query, record.PetId, record.ClientId, record.DoctorId, record.DatetimeStart, record.DatetimeEnd)

	if err != nil {
		return bdErrors.ErrorInsert
	}

	return nil
}

func copyRecord(r RecordPostgres) models.Record {
	record := models.Record{RecordId: r.RecordId, PetId: r.PetId,
		ClientId: r.ClientId, DoctorId: r.DoctorId,

		DatetimeStart: time.Date(
			r.DatetimeStart.Year(),
			r.DatetimeStart.Month(),
			r.DatetimeStart.Day(),
			r.DatetimeStart.Hour(),
			r.DatetimeStart.Minute(),
			r.DatetimeStart.Second(),
			r.DatetimeStart.Nanosecond(),
			time.UTC),

		DatetimeEnd: time.Date(
			r.DatetimeEnd.Year(),
			r.DatetimeEnd.Month(),
			r.DatetimeEnd.Day(),
			r.DatetimeEnd.Hour(),
			r.DatetimeEnd.Minute(),
			r.DatetimeEnd.Second(),
			r.DatetimeEnd.Nanosecond(),
			time.UTC),
	}

	return record
}

func (r *RecordPostgresRepository) GetRecord(id uint64) (*models.Record, error) {
	query := `select * from records where records.id_record = $1;`
	recordDB := &RecordPostgres{}

	err := r.db.Get(recordDB, query, id)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	record := copyRecord(*recordDB)

	return &record, nil
}

func (r *RecordPostgresRepository) GetAllByClient(id uint64) ([]models.Record, error) {
	query := `select * from records where id_client = $1;`
	var recordsPostgres = []RecordPostgres{}
	err := r.db.Select(&recordsPostgres, query, id)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	recordModels := []models.Record{}

	for _, r := range recordsPostgres {

		record := copyRecord(r)

		recordModels = append(recordModels, record)
	}

	return recordModels, nil
}

func (r *RecordPostgresRepository) GetAllByDoctor(id uint64) ([]models.Record, error) {
	query := `select * from records where id_doctor = $1`
	var recordsPostgres []RecordPostgres
	err := r.db.Select(&recordsPostgres, query, id)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	recordModels := []models.Record{}

	for _, r := range recordsPostgres {

		record := copyRecord(r)

		recordModels = append(recordModels, record)
	}

	return recordModels, nil
}

func (r *RecordPostgresRepository) GetAllRecordFilter(doctorId uint64, clientId uint64) ([]models.Record, error) {
	query := `select * from records where records.id_doctor = $1 and records.id_client = $2;`

	var recordsPostgres []RecordPostgres
	err := r.db.Select(&recordsPostgres, query, doctorId, clientId)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	recordModels := []models.Record{}

	for _, r := range recordsPostgres {

		record := copyRecord(r)

		recordModels = append(recordModels, record)
	}

	return recordModels, nil
}

func (r *RecordPostgresRepository) GetAllRecords() ([]models.Record, error) {
	query := `select * from records;`

	var recordsPostgres []RecordPostgres
	err := r.db.Select(&recordsPostgres, query)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	recordModels := []models.Record{}

	for _, r := range recordsPostgres {
		record := copyRecord(r)

		recordModels = append(recordModels, record)
	}

	return recordModels, nil
}

func (r *RecordPostgresRepository) Delete(id uint64) error {
	query := `delete from records where id_record = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return bdErrors.ErrorDelete
	}

	return nil
}

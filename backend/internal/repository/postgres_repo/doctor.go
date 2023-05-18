package postgres_repo

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/bdErrors"
	"backend/internal/pkg/errors/repoErrors"
	"backend/internal/repository"
	"database/sql"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
)

type DoctorPostgres struct {
	DoctorId  uint64 `db:"id_doctor"`
	Login     string `db:"login"`
	Password  string `db:"password"`
	StartTime uint64 `db:"start_time"`
	EndTime   uint64 `db:"end_time"`
}

type DoctorPostgresRepository struct {
	db *sqlx.DB
}

func NewDoctorPostgresRepository(db *sqlx.DB) repository.DoctorRepository {
	return &DoctorPostgresRepository{db: db}
}

func (d *DoctorPostgresRepository) Create(doctor *models.Doctor) error {
	query := `insert into doctors(login, password, start_time, end_time) values($1, $2, $3, $4);`

	_, err := d.db.Exec(query, doctor.Login, doctor.Password, doctor.StartTime, doctor.EndTime)

	if err != nil {
		return bdErrors.ErrorInsert
	}

	return nil
}

func (d *DoctorPostgresRepository) GetDoctorByLogin(login string) (*models.Doctor, error) {
	query := `select * from doctors where login = $1;`
	doctorBD := &DoctorPostgres{}

	err := d.db.Get(doctorBD, query, login)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	doctorModels := &models.Doctor{}
	err = copier.Copy(doctorModels, doctorBD)

	if err != nil {
		return nil, bdErrors.ErrorCopy
	}

	return doctorModels, nil
}

func (d *DoctorPostgresRepository) GetDoctorById(id uint64) (*models.Doctor, error) {
	query := `select * from doctors where id_doctor = $1;`
	doctorBD := &DoctorPostgres{}

	err := d.db.Get(doctorBD, query, id)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	doctorModels := &models.Doctor{}
	err = copier.Copy(doctorModels, doctorBD)

	if err != nil {
		return nil, bdErrors.ErrorCopy
	}

	return doctorModels, nil
}

func (d *DoctorPostgresRepository) GetAllDoctors() ([]models.Doctor, error) {
	query := `select id_doctor, login, start_time, end_time from doctors;`
	doctorDB := []DoctorPostgres{}

	err := d.db.Select(&doctorDB, query)

	if err == sql.ErrNoRows {
		return nil, repoErrors.EntityDoesNotExists
	} else if err != nil {
		return nil, bdErrors.ErrorSelect
	}

	doctorModels := []models.Doctor{}

	for _, r := range doctorDB {
		doctor := &models.Doctor{}
		err = copier.Copy(doctor, r)

		if err != nil {
			return nil, bdErrors.ErrorCopy
		}

		doctorModels = append(doctorModels, *doctor)
	}

	return doctorModels, nil
}

func (d *DoctorPostgresRepository) UpdateShedule(id uint64, newStart uint64, newEnd uint64) error {
	query := `update doctors set start_time = $1, end_time = $2 where doctors.id_doctor = $3`

	_, err := d.db.Exec(query, newStart, newEnd, id)

	if err != nil {
		return bdErrors.ErrorUpdate
	}

	return nil
}

func (d *DoctorPostgresRepository) Delete(id uint64) error {
	query := `delete from doctors where id_doctor = $1`

	_, err := d.db.Exec(query, id)

	if err != nil {
		return bdErrors.ErrorDelete
	}

	return nil
}
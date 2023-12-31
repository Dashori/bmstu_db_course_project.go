package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/repoErrors"
	"backend/internal/pkg/errors/servicesErrors"
	"backend/internal/pkg/hasher"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
)

type DoctorServiceImplementation struct {
	DoctorRepository repository.DoctorRepository
	hasher           hasher.Hasher
	logger           *log.Logger
}

func NewDoctorServiceImplementation(
	DoctorRepository repository.DoctorRepository,
	hasher hasher.Hasher,
	logger *log.Logger) services.DoctorService {

	return &DoctorServiceImplementation{
		DoctorRepository: DoctorRepository,
		hasher:           hasher,
		logger:           logger,
	}
}

func (c *DoctorServiceImplementation) SetRole() error {
	err := c.DoctorRepository.SetRole()

	return err
}

func checkShedule(start uint64, end uint64) error {
	if start >= end {
		return serviceErrors.ErrorWrongNewShedule
	}

	if start > 24 || end > 24 {
		return serviceErrors.ErrorWrongNewShedule
	}
	return nil
}

func (d *DoctorServiceImplementation) GetDoctorByLogin(login string) (*models.Doctor, error) {
	doctor, err := d.DoctorRepository.GetDoctorByLogin(login)

	if err != nil {
		d.logger.Warn("DOCTOR! Error in repository GetDoctorByLogin", "login", login, "error", err)
		return nil, err
	}

	d.logger.Debug("DOCTOR! Successfully GetDoctorByLogin", "login", login)
	return doctor, nil
}

func (d *DoctorServiceImplementation) Create(doctor *models.Doctor, password string) (*models.Doctor, error) {
	d.logger.Debug("DOCTOR! Start create doctor with", "login", doctor.Login)
	_, err := d.DoctorRepository.GetDoctorByLogin(doctor.Login)

	if err != nil && err != repoErrors.EntityDoesNotExists {
		d.logger.Warn("DOCTOR! Error in repository method GetDoctorByLogin", "err", err)
		return nil, err
	} else if err == nil {
		d.logger.Warn("DOCTOR! Error, doctor with this login already exists", "login", doctor.Login)
		return nil, serviceErrors.DoctorAlreadyExists
	}

	passwordHash, err := d.hasher.GetHash(password)
	if err != nil {
		d.logger.Warn("DOCTOR! Error get hash for password", "login", doctor.Login)
		return nil, serviceErrors.ErrorHash
	}

	doctor.Password = string(passwordHash)
	err = checkShedule(doctor.StartTime, doctor.EndTime)
	if err != nil {
		d.logger.Warn("DOCTOR! Error shedule time for", "doctor", doctor.Login, "start time", doctor.StartTime,
			"end time", doctor.EndTime)
		return nil, err
	}

	err = d.DoctorRepository.Create(doctor)
	if err != nil {
		d.logger.Warn("DOCTOR! Error in repository method Create", "doctor", doctor.Login, "error", err)
		return nil, err
	}

	newDoctor, err := d.GetDoctorByLogin(doctor.Login)
	if err != nil {
		d.logger.Warn("DOCTOR! Error in repository method GetDoctorByLogin", "doctor", doctor.Login, "error", err)
	}

	d.logger.Info("DOCTOR! Successfully create doctor", "doctor", doctor.Login)
	return newDoctor, nil
}

func (d *DoctorServiceImplementation) Login(login, password string) (*models.Doctor, error) {
	d.logger.Debug("DOCTOR! Start login with", "login", login)

	tempDoctor, err := d.DoctorRepository.GetDoctorByLogin(login)
	if err != nil && err == repoErrors.EntityDoesNotExists {
		d.logger.Warn("DOCTOR! Error, doctor with this login does not exists", "login", login, "error", err)
		return nil, serviceErrors.DoctorDoesNotExists
	} else if err != nil {
		d.logger.Warn("DOCTOR! Error in repository method GetDoctorByLogin", "login", login, "error", err)
		return nil, err
	}

	if !d.hasher.CheckUnhashedValue(tempDoctor.Password, password) { //== false
		d.logger.Warn("DOCTOR! Error doctor password", "login", login)
		return nil, serviceErrors.InvalidPassword
	}

	d.logger.Info("DOCTOR! Successfully login doctor", "login", login)
	return tempDoctor, nil
}

func (d *DoctorServiceImplementation) UpdateShedule(id uint64, newStart uint64, newEnd uint64) error {
	_, err := d.GetDoctorById(id)
	if err != nil {
		return err
	}

	err = checkShedule(newStart, newEnd)

	if err != nil {
		d.logger.Warn("DOCTOR! Error new shedule for doctor", "doctorId", id,
			"new startTime", newStart, "new endTime", newEnd)
		return err
	}

	err = d.DoctorRepository.UpdateShedule(id, newStart, newEnd)
	if err != nil {
		d.logger.Warn("DOCTOR! Error in repository method UpdateShedule", "doctorId", id,
			"new startTime", newStart, "new endTime", newEnd, "err", err)
		return err
	}

	d.logger.Info("DOCTOR! Successfully update shedule", "doctorId", id,
		"new startTime", newStart, "new endTime", newEnd)
	return nil
}

func (d *DoctorServiceImplementation) GetAllDoctors() ([]models.Doctor, error) {
	doctors, err := d.DoctorRepository.GetAllDoctors()

	if err != nil {
		d.logger.Warn("DOCTOR! Error in repository method GetDoctorById", "err", err)
		return nil, err
	}

	d.logger.Info("DOCTOR! Successsully repository method GetAllDoctors")
	return doctors, nil
}

func (d *DoctorServiceImplementation) GetDoctorById(id uint64) (*models.Doctor, error) {
	doctor, err := d.DoctorRepository.GetDoctorById(id)

	if err != nil && err == repoErrors.EntityDoesNotExists {
		d.logger.Warn("DOCTOR! Error, doctor with this id does not exists", "id", id, "error", err)
		return nil, serviceErrors.DoctorDoesNotExists
	} else if err != nil {
		d.logger.Warn("DOCTOR! Error in repository method GetDoctorById", "id", id, "error", err)
		return nil, err
	}

	d.logger.Info("DOCTOR! Successsully repository method GetDoctorById", "id", id)
	return doctor, nil
}

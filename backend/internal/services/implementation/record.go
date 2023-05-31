package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/servicesErrors"
	"backend/internal/repository"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	"time"
)

type RecordServiceImplementation struct {
	RecordRepository repository.RecordRepository
	DoctorRepository repository.DoctorRepository
	ClientRepository repository.ClientRepository
	PetRepository    repository.PetRepository
	logger           *log.Logger
}

func NewRecordServiceImplementation(
	RecordRepository repository.RecordRepository,
	DoctorRepository repository.DoctorRepository,
	ClientRepository repository.ClientRepository,
	PetRepository repository.PetRepository,
	logger *log.Logger) services.RecordService {

	return &RecordServiceImplementation{
		RecordRepository: RecordRepository,
		DoctorRepository: DoctorRepository,
		ClientRepository: ClientRepository,
		PetRepository:    PetRepository,
		logger:           logger,
	}
}

func CheckDoctorTime(doctor *models.Doctor, record *models.Record) bool {
	datetimeStartHour := uint64(record.DatetimeStart.Hour()) // час начала приема
	datetimeEndHour := uint64(record.DatetimeEnd.Hour())     // час конца приема

	// неверное начало приема
	if datetimeStartHour < doctor.StartTime || datetimeStartHour > doctor.EndTime {
		return false
	}

	// неверный конец приема
	if datetimeEndHour < doctor.StartTime || datetimeEndHour > doctor.EndTime {
		return false
	}

	return true
}

func CheckOtherRecords(records []models.Record, record *models.Record) bool {
	datetimeStart := record.DatetimeStart

	for _, rec := range records {
		if rec.DatetimeStart.Year() == datetimeStart.Year() &&
			rec.DatetimeStart.Month() == datetimeStart.Month() &&
			rec.DatetimeStart.Day() == datetimeStart.Day() {

			if rec.DatetimeStart.Hour() == datetimeStart.Hour() {
				return false
			}
		}
	}

	return true
}

func CheckTime(record *models.Record) bool {
	datetimeStart := record.DatetimeStart
	datetimeEnd := record.DatetimeEnd

	// не один день
	if datetimeStart.Year() != datetimeEnd.Year() ||
		datetimeStart.Month() != datetimeEnd.Month() ||
		datetimeStart.Day() != datetimeEnd.Day() {
		return false
	}

	// начало позже конца
	if datetimeStart.Hour() >= datetimeEnd.Hour() {
		return false
	}

	if datetimeEnd.Hour()-datetimeStart.Hour() != 1 {
		return false
	}

	// начало и конец только в 00
	if datetimeStart.Minute() != 00 || datetimeEnd.Minute() != 00 {
		return false
	}

	today := time.Now()

	return datetimeStart.After(today)
}

func isClientPetOwner(pets []models.Pet, petId uint64) bool {
	var clientIsPetOwner bool = false

	for _, r := range pets {
		if r.PetId == petId {
			clientIsPetOwner = true
			break
		}
	}

	return clientIsPetOwner
}

func (r *RecordServiceImplementation) CreateRecord(record *models.Record) error {
	r.logger.Debug("RECORD! Start CreateRecord with params", "client", record.ClientId, "pet", record.PetId,
		"doctor", record.DoctorId, "DatetimeStart", record.DatetimeStart, "DatetimeEnd", record.DatetimeEnd)

	if !CheckTime(record) { // == false
		r.logger.Warn("RECORD! Error in CreateRecord", "start time", record.DatetimeStart,
			"end time", record.DatetimeEnd)
		return serviceErrors.ErrorCreateRecordTime
	}

	pets, err := r.PetRepository.GetAllByClient(record.ClientId)
	if err != nil {
		r.logger.Warn("RECORD! Error in repository method GetAllByClient", "error", err)
		return err
	}

	if !isClientPetOwner(pets, record.PetId) { //== false
		r.logger.Warn("RECORD! Client is not the pet owner", "client", record.ClientId, "pet", record.PetId)
		return serviceErrors.NotClientPet
	}

	doctor, err := r.DoctorRepository.GetDoctorById(record.DoctorId)
	if err != nil {
		r.logger.Warn("RECORD! Error in repository method GetDoctorById", "error", err)
		return serviceErrors.DoctorDoesNotExists
	}

	if !CheckDoctorTime(doctor, record) { //== false
		r.logger.Warn("RECORD! Error doctor time for new record", "doctorId", doctor.DoctorId)
		return serviceErrors.ErrorDoctorTime
	}

	records, err := r.RecordRepository.GetAllByDoctor(record.DoctorId)
	if err != nil {
		r.logger.Warn("RECORD! Error in repository method GetAllByDoctor", "error", err)
		return err
	}

	if !CheckOtherRecords(records, record) { //== false
		r.logger.Warn("RECORD! Error, other record has same time")
		return serviceErrors.TimeIsTaken
	}

	err = r.RecordRepository.Create(record)
	if err != nil {
		r.logger.Warn("RECORD! Error in repository method Create", "error", err)
		return err
	}

	r.logger.Info("RECORD! Successfully CreateRecord for", "client", record.ClientId, "pet", record.PetId)

	return nil
}

func (r *RecordServiceImplementation) DeleteRecord(recordId uint64) error {
	_, err := r.RecordRepository.GetRecord(recordId)
	if err != nil {
		r.logger.Warn("RECORD! Error in repository method GetRecord with ", "id", recordId, "error", err)
		return serviceErrors.RecordDoesNotExists
	}

	err = r.RecordRepository.Delete(recordId)
	if err != nil {
		r.logger.Warn("RECORD! Error in repository method Delete with ", "id", recordId, "error", err)
		return err
	}

	r.logger.Info("RECORD! Success Delete record with ", "id", recordId)
	return nil
}

func (r *RecordServiceImplementation) GetRecord(recordId uint64) (*models.Record, error) {
	record, err := r.RecordRepository.GetRecord(recordId)

	if err != nil {
		r.logger.Warn("RECORD! Error in repository method GetRecord with ", "id", recordId, "error", err)
		return nil, serviceErrors.RecordDoesNotExists
	}

	r.logger.Debug("RECORD! Success GetRecord with ", "id", recordId)

	return record, nil
}

func (r *RecordServiceImplementation) GetAllRecords(doctorId uint64, clientId uint64) ([]models.Record, error) {
	var err error
	var records []models.Record

	r.logger.Debug("RECORD! Start GetAllRecords with", "clientId", clientId, "doctorId", doctorId)

	if clientId == 0 && doctorId == 0 {
		records, err = r.RecordRepository.GetAllRecords()
		if err != nil {
			r.logger.Warn("Error in repository method GetAllRecords", "error", err)
			return nil, err
		}
	} else if clientId == 0 && doctorId != 0 {
		records, err = r.RecordRepository.GetAllByDoctor(doctorId)
		if err != nil {
			r.logger.Warn("Error in repository method GetAllByDoctor with", "doctorId", doctorId, "error", err)
			return nil, err
		}
	} else if clientId != 0 && doctorId == 0 {
		records, err = r.RecordRepository.GetAllByClient(clientId)
		if err != nil {
			r.logger.Warn("Error in repository method GetAllByClient with", "clientId", clientId, "error", err)
			return nil, err
		}
	} else {
		records, err = r.RecordRepository.GetAllRecordFilter(doctorId, clientId)
		if err != nil {
			r.logger.Warn("Error in repository method GetAllRecordFilter with", "clientId", clientId, "doctorId", doctorId,
				"error", err)
			return nil, err
		}
	}

	r.logger.Info("RECORD! Success GetAllRecords with", "clientId", clientId, "doctorId", doctorId)

	return records, err
}

func (r *RecordServiceImplementation) CreateRecordResearch(record *models.Record) (error, time.Duration) {

	if !CheckTime(record) { // == false

		return serviceErrors.ErrorCreateRecordTime, 0
	}

	pets, err := r.PetRepository.GetAllByClient(record.ClientId)
	if err != nil {
		return err, 0
	}

	if !isClientPetOwner(pets, record.PetId) { //== false
		return serviceErrors.NotClientPet, 0
	}

	start := time.Now()

	doctor, err := r.DoctorRepository.GetDoctorById(record.DoctorId)
	if err != nil {
		return serviceErrors.DoctorDoesNotExists, 0
	}

	if !CheckDoctorTime(doctor, record) { //== false
		return serviceErrors.ErrorDoctorTime, 0
	}

	records, err := r.RecordRepository.GetAllByDoctor(record.DoctorId)
	if err != nil {
		return err, 0
	}

	if !CheckOtherRecords(records, record) { //== false
		return serviceErrors.TimeIsTaken, 0
	}

	err = r.RecordRepository.Create(record)

	duration := time.Since(start)

	if err != nil {
		return err, 0
	}

	return nil, duration
}



func (r *RecordServiceImplementation) CreateRecordResearchTrigger(record *models.Record) (error, time.Duration) {

	if !CheckTime(record) { // == false

		return serviceErrors.ErrorCreateRecordTime, 0
	}

	pets, err := r.PetRepository.GetAllByClient(record.ClientId)
	if err != nil {
		return err, 0
	}

	if !isClientPetOwner(pets, record.PetId) { //== false
		return serviceErrors.NotClientPet, 0
	}

	start := time.Now()

	err = r.RecordRepository.Create(record)

	duration := time.Since(start)

	if err != nil {
		return err, 0
	}

	return nil, duration
}

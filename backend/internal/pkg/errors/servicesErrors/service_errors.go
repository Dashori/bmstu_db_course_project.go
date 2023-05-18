package serviceErrors

import "errors"

var (
	// records
	ErrorCreateRecordTime = errors.New("Service error! Неверное время для записи!")
	TimeIsTaken           = errors.New("Service error! Данное время уже занято!")
	ErrorDoctorTime       = errors.New("Service error! Неверное время для записи к этому доктору!")

	NotClientPet = errors.New("Service error! Данный питомец не принадлежит вам!")

	// entity does not exists
	DoctorDoesNotExists = errors.New("Service error! Такого доктора не существует!")
	ClientDoesNotExists = errors.New("Service error! Такого клиента не существует!")
	RecordDoesNotExists = errors.New("Service error! Такой записи не существует!")
	PetDoesNotExists    = errors.New("Service error! Такого питомца не существует!")

	// entity already exists
	PetAlreadyExists    = errors.New("Service error! Питомец уже существует в базе!")
	DoctorAlreadyExists = errors.New("Service error! Доктор уже существует в базе!")
	ClientAlreadyExists = errors.New("Service error! Клиент уже существует в базе!")

	// Create + login
	ErrorGetClientByLogin = errors.New("Service error! Ошибка при получении клиента по логину!")
	ErrorGetDoctorByLogin = errors.New("Service error! Ошибка при получении доктора по логину!")
	ErrorHash             = errors.New("Service error! Ошибка получения хэша для пароля!")
	InvalidPassword       = errors.New("Service error! Неверный пароль!")

	ErrorWrongNewShedule = errors.New("Service error! Неверное время для расписания!")
)

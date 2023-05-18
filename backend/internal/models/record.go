package models

import "time"

type Record struct {
	RecordId      uint64
	PetId         uint64
	ClientId      uint64
	DoctorId      uint64
	DatetimeStart time.Time // содержит и дату и время
	DatetimeEnd   time.Time // содержит и дату и время
}

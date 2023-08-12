package models

import "time"

type Record struct {
	RecordId      uint64
	PetId         uint64
	PetName  	  string
	ClientId      uint64
	ClientLogin   string
	DoctorId      uint64
	DoctorLogin   string
	DatetimeStart time.Time // содержит и дату и время
	DatetimeEnd   time.Time // содержит и дату и время
}

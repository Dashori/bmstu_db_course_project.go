package models

import "time"

type ErrorBody struct {
	Err string `json:"error"`
}

type NewShedule struct {
	Start uint64
	End   uint64
}

type Doctors struct {
	Doctors []struct {
		DoctorId  uint64 `json:"DoctorId"`
		Login     string `json:"Login"`
		StartTime uint64 `json:"StartTime"`
		EndTime   uint64 `json:"EndTime"`
	} `json:"doctors"`
}

type Pets struct {
	Pets []struct {
		PetId    uint64 `json:"PetId"`
		Name     string `json:"Name"`
		Type     string `json:"Type"`
		Age      uint64 `json:"Age"`
		Health   uint64 `json:"Health"`
		ClientId uint64 `json:"ClientId"`
	} `json:"pets"`
}

type Records struct {
	Records []struct {
		RecordId      uint64    `json:"RecordId"`
		PetId         uint64    `json:"PetId"`
		ClientId      uint64    `json:"ClientId"`
		DoctorId      uint64    `json:"DoctorId"`
		DatetimeStart time.Time `json:"DatetimeStart"`
		DatetimeEnd   time.Time `json:"DatetimeEnd"`
	} `json:"records"`
}

type Client struct {
	ClientId uint64 `json:"ClientId"`
	Login    string `json:"Login"`
	Password string
	Token    string `json:"Token"`
}

type Doctor struct {
	DoctorId  uint64 `json:"DoctorId"`
	Login     string `json:"Login"`
	Password  string
	Token     string `json:"Token"`
	StartTime uint64 `json:"StartTime"`
	EndTime   uint64 `json:"EndTime"`
}

type Pet struct {
	PetId    uint64 `json:"PetId"`
	Name     string `json:"Name"`
	Type     string `json:"Type"`
	Age      uint64 `json:"Age"`
	Health   uint64 `json:"Health"`
	ClientId uint64 `json:"ClientId"`
}

type Record struct {
	RecordId      uint64
	PetId         uint64
	ClientId      uint64
	DoctorId      uint64
	Year          int
	Month         int
	Day           int
	DatetimeStart int
}

package models

type Doctor struct {
	DoctorId  uint64
	Login     string
	Password  string
	StartTime uint64
	EndTime   uint64
}

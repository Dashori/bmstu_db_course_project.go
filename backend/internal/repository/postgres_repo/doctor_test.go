package postgres_repo

import (
	"backend/internal/models"
	"github.com/stretchr/testify/require"
	"testing"
)

var testDoctorPostgresRepositoryCreateSuccess = []struct {
	TestName  string
	InputData struct {
		doctor *models.Doctor
	}
	CheckOutput     func(t *testing.T, err error)
	CheckOutputHelp func(t *testing.T, err error)
}{
	{
		TestName: "create success test",
		InputData: struct {
			doctor *models.Doctor
		}{&models.Doctor{Login: "ChicagoTest", Password: "12345", StartTime: 15, EndTime: 19}},

		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
		CheckOutputHelp: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestDoctorPostgresRepositoryCreate(t *testing.T) {
	for _, tt := range testDoctorPostgresRepositoryCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields, err := CreatePostgresRepositoryFieldsTest(configFileName, pathToConfig)

			doctorRepository := CreateDoctorPostgresRepository(fields)

			err = doctorRepository.Create(tt.InputData.doctor)

			tt.CheckOutput(t, err)

			doctor, err := doctorRepository.GetDoctorByLogin("ChicagoTest")

			if err == nil {
				err = doctorRepository.Delete(doctor.DoctorId)
			}

			tt.CheckOutputHelp(t, err)
		})
	}
}

var testDoctorPostgresRepositoryGetId = []struct {
	TestName  string
	InputData struct {
		doctor *models.Doctor
	}
	CheckOutput     func(t *testing.T, doctor *models.Doctor, err error)
	CheckOutputHelp func(t *testing.T, err error)
}{
	{
		TestName: "get by id success test",
		InputData: struct {
			doctor *models.Doctor
		}{&models.Doctor{Login: "ChicagoTest", Password: "12345", StartTime: 15, EndTime: 19}},

		CheckOutput: func(t *testing.T, doctor *models.Doctor, err error) {
			require.NoError(t, err)
			require.Equal(t, doctor.Login, "ChicagoTest")
			require.Equal(t, doctor.Password, "12345")
			require.Equal(t, doctor.StartTime, uint64(15))
			require.Equal(t, doctor.EndTime, uint64(19))
		},
		CheckOutputHelp: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestDoctorPostgresRepositoryGetId(t *testing.T) {
	for _, tt := range testDoctorPostgresRepositoryGetId {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			fields, err := CreatePostgresRepositoryFieldsTest(configFileName, pathToConfig)

			doctorRepository := CreateDoctorPostgresRepository(fields)

			err = doctorRepository.Create(tt.InputData.doctor)

			tt.CheckOutputHelp(t, err)

			doctor, err := doctorRepository.GetDoctorByLogin("ChicagoTest")

			tt.CheckOutputHelp(t, err)

			doctor, err = doctorRepository.GetDoctorById(doctor.DoctorId)

			tt.CheckOutput(t, doctor, err)

			err = doctorRepository.Delete(doctor.DoctorId)

			tt.CheckOutputHelp(t, err)
		})
	}
}

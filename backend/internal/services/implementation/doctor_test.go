package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/dbErrors"
	"backend/internal/pkg/errors/repoErrors"
	"backend/internal/pkg/errors/servicesErrors"
	mock_hasher "backend/internal/pkg/hasher/mocks"
	mock_repository "backend/internal/repository/mocks"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type doctorServiceFields struct {
	doctorRepositoryMock *mock_repository.MockDoctorRepository
	hasherMock           *mock_hasher.MockHasher
	logger               *log.Logger
}

func createDoctorServiceFields(controller *gomock.Controller) *doctorServiceFields {
	fields := new(doctorServiceFields)

	fields.hasherMock = mock_hasher.NewMockHasher(controller)
	fields.doctorRepositoryMock = mock_repository.NewMockDoctorRepository(controller)
	fields.logger = log.New(os.Stderr)
	fields.logger.SetLevel(log.FatalLevel)

	return fields
}

func createDoctorService(fields *doctorServiceFields) services.DoctorService {
	return NewDoctorServiceImplementation(fields.doctorRepositoryMock, fields.hasherMock, fields.logger)
}

//-------------------------------------------------------------------------------------------------
// create

var testDoctorCreateSuccess = []struct {
	TestName  string
	InputData struct {
		doctor   *models.Doctor
		password string
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple create success",
		InputData: struct {
			doctor   *models.Doctor
			password string
		}{doctor: &models.Doctor{Login: "Chepigo", StartTime: 10, EndTime: 12}, password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").Return(nil, repoErrors.EntityDoesNotExists)
			fields.hasherMock.EXPECT().GetHash("12345").Return([]byte("12345_hash"), nil)
			fields.doctorRepositoryMock.EXPECT().Create(
				&models.Doctor{Login: "Chepigo", Password: "12345_hash", StartTime: 10, EndTime: 12})
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo")
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testDoctorCreateFailure = []struct {
	TestName  string
	InputData struct {
		doctor   *models.Doctor
		password string
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "create error, get doctor by login ",
		InputData: struct {
			doctor   *models.Doctor
			password string
		}{doctor: &models.Doctor{Login: "Chepigo"}, password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").Return(nil, repoErrors.ErrorGetDoctorByLogin)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, repoErrors.ErrorGetDoctorByLogin)
		},
	},
	{
		TestName: "create error, doctor already exist",
		InputData: struct {
			doctor   *models.Doctor
			password string
		}{doctor: &models.Doctor{Login: "Chepigo"}, password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").Return(&models.Doctor{}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.DoctorAlreadyExists)
		},
	},
	{
		TestName: "create error, error hash",
		InputData: struct {
			doctor   *models.Doctor
			password string
		}{doctor: &models.Doctor{Login: "Chepigo"}, password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").Return(nil, repoErrors.EntityDoesNotExists)
			fields.hasherMock.EXPECT().GetHash("12345").Return(nil, serviceErrors.ErrorHash)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorHash)
		},
	},
	{
		TestName: "create error, error shedule, start after end",
		InputData: struct {
			doctor   *models.Doctor
			password string
		}{doctor: &models.Doctor{Login: "Chepigo", StartTime: 11, EndTime: 10}, password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").Return(nil, repoErrors.EntityDoesNotExists)
			fields.hasherMock.EXPECT().GetHash("12345").Return([]byte("12345_hash"), nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorWrongNewShedule)
		},
	},
	{
		TestName: "create error, error shedule > 24 ",
		InputData: struct {
			doctor   *models.Doctor
			password string
		}{doctor: &models.Doctor{Login: "Chepigo", StartTime: 23, EndTime: 56}, password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").Return(nil, repoErrors.EntityDoesNotExists)
			fields.hasherMock.EXPECT().GetHash("12345").Return([]byte("12345_hash"), nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorWrongNewShedule)
		},
	},
	{
		TestName: "create error, error repo create",
		InputData: struct {
			doctor   *models.Doctor
			password string
		}{doctor: &models.Doctor{Login: "Chepigo", StartTime: 10, EndTime: 12}, password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").Return(nil, repoErrors.EntityDoesNotExists)
			fields.hasherMock.EXPECT().GetHash("12345").Return([]byte("12345_hash"), nil)
			fields.doctorRepositoryMock.EXPECT().Create(
				&models.Doctor{Login: "Chepigo", Password: "12345_hash", StartTime: 10, EndTime: 12}).
				Return(dbErrors.ErrorInsert)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, dbErrors.ErrorInsert)
		},
	},
}

func TestDoctorServiceImplementationCreate(t *testing.T) {
	t.Parallel()

	for _, tt := range testDoctorCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			_, err := doctorService.Create(tt.InputData.doctor, tt.InputData.password)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testDoctorCreateFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			_, err := doctorService.Create(tt.InputData.doctor, tt.InputData.password)

			tt.CheckOutput(t, err)

		})
	}
}

//-------------------------------------------------------------------------------------------------
// Login

var testDoctorLoginSuccess = []struct {
	TestName  string
	InputData struct {
		login    string
		password string
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple login",
		InputData: struct {
			login    string
			password string
		}{login: "Chepigo", password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").Return(&models.Doctor{
				DoctorId: 1,
				Login:    "Chepigo",
				Password: "12345"}, nil)
			fields.hasherMock.EXPECT().CheckUnhashedValue("12345", "12345").Return(true)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testDoctorLoginFailure = []struct {
	TestName  string
	InputData struct {
		login    string
		password string
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "error login, entity does not exist",
		InputData: struct {
			login    string
			password string
		}{login: "Chepigo", password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").
				Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.DoctorDoesNotExists)
		},
	},

	{
		TestName: "error login, bad password",
		InputData: struct {
			login    string
			password string
		}{login: "Chepigo", password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").
				Return(&models.Doctor{
					DoctorId: 1,
					Login:    "Chepigo",
					Password: "123456"}, nil)
			fields.hasherMock.EXPECT().CheckUnhashedValue("123456", "12345").Return(false)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.InvalidPassword)
		},
	},

	{
		TestName: "error login, repo error",
		InputData: struct {
			login    string
			password string
		}{login: "Chepigo", password: "12345"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").
				Return(nil, dbErrors.ErrorSelect)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, dbErrors.ErrorSelect)
		},
	},
}

func TestDoctorServiceImplementationLogin(t *testing.T) {
	t.Parallel()

	for _, tt := range testDoctorLoginSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			_, err := doctorService.Login(tt.InputData.login, tt.InputData.password)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testDoctorLoginFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			_, err := doctorService.Login(tt.InputData.login, tt.InputData.password)

			tt.CheckOutput(t, err)
		})
	}
}

//-------------------------------------------------------------------------------------------------
// Update shedule

var testDoctorUpdateSheduleSuccess = []struct {
	TestName  string
	InputData struct {
		id       uint64
		newStart uint64
		newEnd   uint64
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple update",
		InputData: struct {
			id       uint64
			newStart uint64
			newEnd   uint64
		}{id: 1, newStart: 10, newEnd: 18},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).Return(&models.Doctor{
				DoctorId: 1,
				Login:    "Chepigo",
				Password: "12345"}, nil)

			fields.doctorRepositoryMock.EXPECT().UpdateShedule(uint64(1), uint64(10), uint64(18)).Return(nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testDoctorUpdateSheduleFailure = []struct {
	TestName  string
	InputData struct {
		id       uint64
		newStart uint64
		newEnd   uint64
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "failure update, entity does not exists",
		InputData: struct {
			id       uint64
			newStart uint64
			newEnd   uint64
		}{id: 1, newStart: 10, newEnd: 18},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.DoctorDoesNotExists)
		},
	},
	{
		TestName: "failure update, bad time, start > end",
		InputData: struct {
			id       uint64
			newStart uint64
			newEnd   uint64
		}{id: 1, newStart: 18, newEnd: 10},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).Return(&models.Doctor{
				DoctorId: 1,
				Login:    "Chepigo",
				Password: "12345"}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorWrongNewShedule)
		},
	},
	{
		TestName: "failure update, bad return from repo",
		InputData: struct {
			id       uint64
			newStart uint64
			newEnd   uint64
		}{id: 1, newStart: 18, newEnd: 10},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).Return(nil,
				dbErrors.ErrorSelect)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, dbErrors.ErrorSelect)
		},
	},
}

func TestDoctorServiceImplementationUpdateShedule(t *testing.T) {
	t.Parallel()

	for _, tt := range testDoctorUpdateSheduleSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			err := doctorService.UpdateShedule(tt.InputData.id, tt.InputData.newStart, tt.InputData.newEnd)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testDoctorUpdateSheduleFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			err := doctorService.UpdateShedule(tt.InputData.id, tt.InputData.newStart, tt.InputData.newEnd)

			tt.CheckOutput(t, err)
		})
	}
}

//-------------------------------------------------------------------------------------------------
// GetAllDoctors

var testDoctorGetAllDoctorsSuccess = []struct {
	TestName  string
	InputData struct {
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple get all doctors",
		InputData: struct {
		}{},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetAllDoctors().Return([]models.Doctor{
				{
					DoctorId: 1,
					Login:    "Chepigo",
					Password: "12345"}}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testDoctorGetAllDoctorsFailure = []struct {
	TestName    string
	InputData   struct{}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "failure get all doctors, repo error",
		InputData: struct {
		}{},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetAllDoctors().Return(nil, dbErrors.ErrorSelect)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, dbErrors.ErrorSelect)
		},
	},
}

func TestDoctorServiceImplementationGetAllDoctors(t *testing.T) {
	t.Parallel()

	for _, tt := range testDoctorGetAllDoctorsSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			_, err := doctorService.GetAllDoctors()

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testDoctorGetAllDoctorsFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			_, err := doctorService.GetAllDoctors()

			tt.CheckOutput(t, err)
		})
	}
}

//-------------------------------------------------------------------------------------------------
// GetDoctorById(id uint64) (*models.Doctor, error)

var testDoctorGetDoctorByIdSuccess = []struct {
	TestName  string
	InputData struct {
		id uint64
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple get doctor by id",
		InputData: struct {
			id uint64
		}{id: 1},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).Return(
				&models.Doctor{
					DoctorId: 1,
					Login:    "Chepigo",
					Password: "12345"}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testDoctorGetDoctorByIdFailure = []struct {
	TestName  string
	InputData struct {
		id uint64
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "failure get doctor by id, entity does not exists",
		InputData: struct {
			id uint64
		}{id: 1},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.DoctorDoesNotExists)
		},
	},
	{
		TestName: "failure get doctor by id, repo error",
		InputData: struct {
			id uint64
		}{id: 1},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).Return(nil, dbErrors.ErrorSelect)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, dbErrors.ErrorSelect)
		},
	},
}

func TestDoctorServiceImplementationGetDoctorById(t *testing.T) {
	t.Parallel()

	for _, tt := range testDoctorGetDoctorByIdSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			_, err := doctorService.GetDoctorById(tt.InputData.id)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testDoctorGetDoctorByIdFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			_, err := doctorService.GetDoctorById(tt.InputData.id)

			tt.CheckOutput(t, err)
		})
	}
}

//-------------------------------------------------------------------------------------------------
// GetDoctorByLogin(login string) (*models.Doctor, error)

var testDoctorGetDoctorByLoginSuccess = []struct {
	TestName  string
	InputData struct {
		login string
	}
	Prepare     func(fields *doctorServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple get doctor by id",
		InputData: struct {
			login string
		}{login: "Chepigo"},

		Prepare: func(fields *doctorServiceFields) {
			fields.doctorRepositoryMock.EXPECT().GetDoctorByLogin("Chepigo").Return(
				&models.Doctor{
					DoctorId: 1,
					Login:    "Chepigo",
					Password: "12345"}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestDoctorServiceImplementationGetDoctorByLogin(t *testing.T) {
	t.Parallel()

	for _, tt := range testDoctorGetDoctorByLoginSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createDoctorServiceFields(ctrl)
			tt.Prepare(fields)

			doctorService := createDoctorService(fields)

			_, err := doctorService.GetDoctorByLogin(tt.InputData.login)

			tt.CheckOutput(t, err)
		})
	}
}

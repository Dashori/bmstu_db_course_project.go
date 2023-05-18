package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/bdErrors"
	"backend/internal/pkg/errors/repoErrors"
	"backend/internal/pkg/errors/servicesErrors"
	mock_repository "backend/internal/repository/mocks"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

type recordServiceFields struct {
	recordRepositoryMock *mock_repository.MockRecordRepository
	doctorRepositoryMock *mock_repository.MockDoctorRepository
	clientRepositoryMock *mock_repository.MockClientRepository
	petRepositoryMock    *mock_repository.MockPetRepository
	logger               *log.Logger
}

func createRecordServiceFields(controller *gomock.Controller) *recordServiceFields {
	fields := new(recordServiceFields)

	fields.recordRepositoryMock = mock_repository.NewMockRecordRepository(controller)
	fields.doctorRepositoryMock = mock_repository.NewMockDoctorRepository(controller)
	fields.clientRepositoryMock = mock_repository.NewMockClientRepository(controller)
	fields.petRepositoryMock = mock_repository.NewMockPetRepository(controller)
	fields.logger = log.New(os.Stderr)
	fields.logger.SetLevel(log.FatalLevel)

	return fields
}

func createRecordService(fields *recordServiceFields) services.RecordService {
	return NewRecordServiceImplementation(fields.recordRepositoryMock, fields.doctorRepositoryMock,
		fields.clientRepositoryMock, fields.petRepositoryMock, fields.logger)
}

//-------------------------------------------------------------------------------------------------
// get record

var testRecordGetSuccess = []struct {
	TestName  string
	InputData struct {
		recordId uint64
	}
	Prepare     func(fields *recordServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "record get success",
		InputData: struct {
			recordId uint64
		}{recordId: 1},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetRecord(uint64(1)).Return(&models.Record{RecordId: 1}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testRecordGetFailure = []struct {
	TestName  string
	InputData struct {
		recordId uint64
	}
	Prepare     func(fields *recordServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "record get failure, entity does not exists",
		InputData: struct {
			recordId uint64
		}{recordId: 1},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetRecord(uint64(1)).Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.RecordDoesNotExists)
		},
	},
}

func TestRecordServiceImplementationGet(t *testing.T) {
	t.Parallel()

	for _, tt := range testRecordGetSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createRecordServiceFields(ctrl)
			tt.Prepare(fields)

			recordService := createRecordService(fields)

			_, err := recordService.GetRecord(tt.InputData.recordId)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testRecordGetFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createRecordServiceFields(ctrl)
			tt.Prepare(fields)

			recordService := createRecordService(fields)

			_, err := recordService.GetRecord(tt.InputData.recordId)

			tt.CheckOutput(t, err)
		})
	}
}

//-------------------------------------------------------------------------------------------------
// delete

var testRecordDeleteSuccess = []struct {
	TestName  string
	InputData struct {
		recordId uint64
	}
	Prepare     func(fields *recordServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "record delete success",
		InputData: struct {
			recordId uint64
		}{recordId: 1},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetRecord(uint64(1)).Return(&models.Record{RecordId: 1}, nil)
			fields.recordRepositoryMock.EXPECT().Delete(uint64(1)).Return(nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testRecordDeleteFailure = []struct {
	TestName  string
	InputData struct {
		recordId uint64
	}
	Prepare     func(fields *recordServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "record delete failure",
		InputData: struct {
			recordId uint64
		}{recordId: 1},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetRecord(uint64(1)).Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.RecordDoesNotExists)
		},
	},
}

func TestRecordServiceImplementationDelete(t *testing.T) {
	t.Parallel()

	for _, tt := range testRecordDeleteSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createRecordServiceFields(ctrl)
			tt.Prepare(fields)

			recordService := createRecordService(fields)

			err := recordService.DeleteRecord(tt.InputData.recordId)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testRecordDeleteFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createRecordServiceFields(ctrl)
			tt.Prepare(fields)

			recordService := createRecordService(fields)

			err := recordService.DeleteRecord(tt.InputData.recordId)

			tt.CheckOutput(t, err)

		})
	}
}

//-------------------------------------------------------------------------------------------------
// Create record

var testRecordCreateSuccess = []struct {
	TestName  string
	InputData struct {
		record *models.Record
	}
	Prepare     func(fields *recordServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "record create success",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{DoctorId: 1, PetId: 1, ClientId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 14, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {

			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(
				[]models.Pet{{PetId: 1, ClientId: 1}}, nil)

			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).
				Return(&models.Doctor{DoctorId: 1, StartTime: 0, EndTime: 23}, nil)

			fields.recordRepositoryMock.EXPECT().GetAllByDoctor(uint64(1)).Return([]models.Record{
				{DoctorId: 1,
					DatetimeStart: time.Date(3000, 8, 15, 17, 00, 00, 00, time.UTC),
					DatetimeEnd:   time.Date(3000, 8, 15, 18, 00, 00, 00, time.UTC)}}, nil)

			fields.recordRepositoryMock.EXPECT().Create(
				&models.Record{ClientId: 1, PetId: 1, DoctorId: 1,
					DatetimeStart: time.Date(3000, 8, 15, 14, 00, 00, 00, time.UTC),
					DatetimeEnd:   time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC)})
		},

		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testRecordCreateFailure = []struct {
	TestName  string
	InputData struct {
		record *models.Record
	}
	Prepare     func(fields *recordServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "record create failure, different days",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 14, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 9, 15, 15, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorCreateRecordTime)
		},
	},
	{
		TestName: "record create failure, start after end",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorCreateRecordTime)
		},
	},
	{
		TestName: "record create failure, error minute start",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 15, 01, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorCreateRecordTime)
		},
	},
	{
		TestName: "record create failure, date before now",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1,
			DatetimeStart: time.Date(1900, 8, 15, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(1900, 8, 15, 17, 30, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {},

		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorCreateRecordTime)
		},
	},
	{
		TestName: "record create failure, error get all by client",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).
				Return(nil, bdErrors.ErrorSelect)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, bdErrors.ErrorSelect)
		},
	},
	{
		TestName: "record create failure, not client pet",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1, PetId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).
				Return([]models.Pet{{PetId: 2}}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.NotClientPet)
		},
	},
	{
		TestName: "record create failure, doctor does not exists",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1, PetId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(
				[]models.Pet{{PetId: 1, ClientId: 1}}, nil)

			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).
				Return(nil, repoErrors.EntityDoesNotExists)

		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.DoctorDoesNotExists)
		},
	},
	{
		TestName: "record create failure, error doctor time for START",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1, PetId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 10, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 11, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(
				[]models.Pet{{PetId: 1, ClientId: 1}}, nil)

			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).
				Return(&models.Doctor{DoctorId: 1, StartTime: 14, EndTime: 17}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorDoctorTime)
		},
	},
	{
		TestName: "record create failure, error doctor time for END",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1, PetId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(
				[]models.Pet{{PetId: 1, ClientId: 1}}, nil)

			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).
				Return(&models.Doctor{DoctorId: 1, StartTime: 14, EndTime: 15}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorDoctorTime)
		},
	},
	{
		TestName: "record create failure, error get all records by doctor",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1, PetId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(
				[]models.Pet{{PetId: 1, ClientId: 1}}, nil)

			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).
				Return(&models.Doctor{DoctorId: 1, StartTime: 10, EndTime: 20}, nil)

			fields.recordRepositoryMock.EXPECT().GetAllByDoctor(uint64(1)).
				Return(nil, repoErrors.ErrorGetAllByDoctor)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, repoErrors.ErrorGetAllByDoctor)
		},
	},
	{
		TestName: "record create failure, time is taken",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{ClientId: 1, DoctorId: 1, PetId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(
				[]models.Pet{{PetId: 1, ClientId: 1}}, nil)

			fields.doctorRepositoryMock.EXPECT().GetDoctorById(uint64(1)).
				Return(&models.Doctor{DoctorId: 1, StartTime: 0, EndTime: 23}, nil)

			fields.recordRepositoryMock.EXPECT().GetAllByDoctor(uint64(1)).Return([]models.Record{
				{RecordId: 1, DoctorId: 1, PetId: 1, ClientId: 1,
					DatetimeStart: time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC),
					DatetimeEnd:   time.Date(3000, 8, 15, 16, 00, 00, 00, time.UTC)}}, nil)
		},

		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.TimeIsTaken)
		},
	},
	{
		TestName: "record create failure, not hour ",
		InputData: struct {
			record *models.Record
		}{record: &models.Record{RecordId: 1, DoctorId: 1,
			DatetimeStart: time.Date(3000, 8, 15, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(3000, 8, 15, 17, 30, 00, 00, time.UTC)},
		},

		Prepare: func(fields *recordServiceFields) {},

		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorCreateRecordTime)
		},
	},
}

func TestRecordServiceImplementationCreate(t *testing.T) {
	t.Parallel()

	for _, tt := range testRecordCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createRecordServiceFields(ctrl)
			tt.Prepare(fields)

			recordService := createRecordService(fields)

			err := recordService.CreateRecord(tt.InputData.record)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testRecordCreateFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createRecordServiceFields(ctrl)
			tt.Prepare(fields)

			recordService := createRecordService(fields)

			err := recordService.CreateRecord(tt.InputData.record)

			tt.CheckOutput(t, err)

		})
	}
}

//-------------------------------------------------------------------------------------------------
// GetAllRecords

var testRecordGetAllRecordsSuccess = []struct {
	TestName  string
	InputData struct {
		doctorId uint64
		clientId uint64
	}
	Prepare     func(fields *recordServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "record get all records full success",
		InputData: struct {
			doctorId uint64
			clientId uint64
		}{doctorId: 1, clientId: 1},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetAllRecordFilter(uint64(1), uint64(1)).Return(
				[]models.Record{{ClientId: 1, DoctorId: 1}}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
	{
		TestName: "record get all records only doctor success",
		InputData: struct {
			doctorId uint64
			clientId uint64
		}{doctorId: 1, clientId: 0},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetAllByDoctor(uint64(1)).Return(
				[]models.Record{{DoctorId: 1}}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
	{
		TestName: "record get all records only client success",
		InputData: struct {
			doctorId uint64
			clientId uint64
		}{doctorId: 0, clientId: 1},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(
				[]models.Record{{ClientId: 1}}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
	{
		TestName: "record get all records without filters success",
		InputData: struct {
			doctorId uint64
			clientId uint64
		}{doctorId: 0, clientId: 0},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetAllRecords().Return([]models.Record{}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testRecordGetAllRecordsFailure = []struct {
	TestName  string
	InputData struct {
		doctorId uint64
		clientId uint64
	}
	Prepare     func(fields *recordServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "record get all records failure",
		InputData: struct {
			doctorId uint64
			clientId uint64
		}{doctorId: 1, clientId: 1},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetAllRecordFilter(uint64(1), uint64(1)).Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, repoErrors.EntityDoesNotExists)
		},
	},
	{
		TestName: "record get all records failure, client records does not exists",
		InputData: struct {
			doctorId uint64
			clientId uint64
		}{doctorId: 0, clientId: 1},

		Prepare: func(fields *recordServiceFields) {
			fields.recordRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, repoErrors.EntityDoesNotExists)
		},
	},
}

func TestRecordServiceImplementationGetAllRecords(t *testing.T) {
	t.Parallel()

	for _, tt := range testRecordGetAllRecordsSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createRecordServiceFields(ctrl)
			tt.Prepare(fields)

			recordService := createRecordService(fields)

			_, err := recordService.GetAllRecords(tt.InputData.doctorId, tt.InputData.clientId)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testRecordGetAllRecordsFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createRecordServiceFields(ctrl)
			tt.Prepare(fields)

			recordService := createRecordService(fields)

			_, err := recordService.GetAllRecords(tt.InputData.doctorId, tt.InputData.clientId)

			tt.CheckOutput(t, err)

		})
	}
}

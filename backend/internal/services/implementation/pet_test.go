package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/dbErrors"
	"backend/internal/pkg/errors/repoErrors"
	"backend/internal/pkg/errors/servicesErrors"
	mock_repository "backend/internal/repository/mocks"
	"backend/internal/services"
	"github.com/charmbracelet/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type petServiceFields struct {
	petRepositoryMock    *mock_repository.MockPetRepository
	clientRepositoryMock *mock_repository.MockClientRepository
	logger               *log.Logger
}

func createPetServiceFields(controller *gomock.Controller) *petServiceFields {
	fields := new(petServiceFields)

	fields.petRepositoryMock = mock_repository.NewMockPetRepository(controller)
	fields.clientRepositoryMock = mock_repository.NewMockClientRepository(controller)
	fields.logger = log.New(os.Stderr)
	fields.logger.SetLevel(log.FatalLevel)

	return fields
}

func createPetService(fields *petServiceFields) services.PetService {
	return NewPetServiceImplementation(fields.petRepositoryMock, fields.clientRepositoryMock, fields.logger)
}

//-------------------------------------------------------------------------------------------------
// create

var testPetCreateSuccess = []struct {
	TestName  string
	InputData struct {
		pet   *models.Pet
		login string
	}
	Prepare     func(fields *petServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple create pet",
		InputData: struct {
			pet   *models.Pet
			login string
		}{pet: &models.Pet{PetId: 1, Name: "Havrosha", ClientId: 1}, login: "Chepigo"},
		Prepare: func(fields *petServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").Return(
				&models.Client{ClientId: uint64(1)}, nil)
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return([]models.Pet{}, nil)
			fields.petRepositoryMock.EXPECT().Create(&models.Pet{PetId: 1, Name: "Havrosha", ClientId: 1})
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testPetCreateFailure = []struct {
	TestName  string
	InputData struct {
		pet   *models.Pet
		login string
	}
	Prepare     func(fields *petServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "create error, client does not exists",
		InputData: struct {
			pet   *models.Pet
			login string
		}{pet: &models.Pet{PetId: 1, Name: "Havrosha", ClientId: 1}, login: "Chepigo"},

		Prepare: func(fields *petServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ClientDoesNotExists)
		},
	},
	{
		TestName: "create error, get all by client",
		InputData: struct {
			pet   *models.Pet
			login string
		}{pet: &models.Pet{PetId: 1, Name: "Havrosha", ClientId: 1}, login: "Chepigo"},

		Prepare: func(fields *petServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").Return(&models.Client{ClientId: uint64(1)}, nil)
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(nil, repoErrors.ErrorGetAllByClient)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, repoErrors.ErrorGetAllByClient)
		},
	},
	{
		TestName: "create error, pet already exist",
		InputData: struct {
			pet   *models.Pet
			login string
		}{pet: &models.Pet{PetId: 1, Name: "Havrosha", ClientId: 1}, login: "Chepigo"},

		Prepare: func(fields *petServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").Return(&models.Client{ClientId: uint64(1)}, nil)
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(
				[]models.Pet{{PetId: 1, Name: "Havrosha", ClientId: 1}}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.PetAlreadyExists)
		},
	},
}

func TestPetServiceImplementationCreate(t *testing.T) {
	t.Parallel()

	for _, tt := range testPetCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createPetServiceFields(ctrl)
			tt.Prepare(fields)

			petService := createPetService(fields)

			err := petService.Create(tt.InputData.pet, tt.InputData.login)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testPetCreateFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createPetServiceFields(ctrl)
			tt.Prepare(fields)

			petService := createPetService(fields)

			err := petService.Create(tt.InputData.pet, tt.InputData.login)

			tt.CheckOutput(t, err)

		})
	}
}

//-------------------------------------------------------------------------------------------------
// delete

var testPetDeleteSuccess = []struct {
	TestName  string
	InputData struct {
		petId    uint64
		clientId uint64
	}
	Prepare     func(fields *petServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "delete success",
		InputData: struct {
			petId    uint64
			clientId uint64
		}{petId: 1, clientId: 1},

		Prepare: func(fields *petServiceFields) {
			fields.petRepositoryMock.EXPECT().GetPet(uint64(1)).Return(&models.Pet{PetId: 1, ClientId: 1}, nil)
			fields.petRepositoryMock.EXPECT().Delete(uint64(1)).Return(nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testPetDeleteFailure = []struct {
	TestName  string
	InputData struct {
		petId    uint64
		clientId uint64
	}
	Prepare     func(fields *petServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "delete error, entity does not exists",
		InputData: struct {
			petId    uint64
			clientId uint64
		}{petId: 1, clientId: 1},

		Prepare: func(fields *petServiceFields) {
			fields.petRepositoryMock.EXPECT().GetPet(uint64(1)).Return(nil, repoErrors.EntityDoesNotExists)

		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.PetDoesNotExists)
		},
	},
	{
		TestName: "delete error, not client pet",
		InputData: struct {
			petId    uint64
			clientId uint64
		}{petId: 1, clientId: 2},

		Prepare: func(fields *petServiceFields) {
			fields.petRepositoryMock.EXPECT().GetPet(uint64(1)).Return(&models.Pet{PetId: 1, ClientId: 1}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.NotClientPet)
		},
	},
}

func TestPetServiceImplementationDelete(t *testing.T) {
	t.Parallel()

	for _, tt := range testPetDeleteSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createPetServiceFields(ctrl)
			tt.Prepare(fields)

			petService := createPetService(fields)

			err := petService.Delete(tt.InputData.petId, tt.InputData.clientId)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testPetDeleteFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createPetServiceFields(ctrl)
			tt.Prepare(fields)

			petService := createPetService(fields)

			err := petService.Delete(tt.InputData.petId, tt.InputData.clientId)

			tt.CheckOutput(t, err)

		})
	}
}

//-------------------------------------------------------------------------------------------------
// get

var testGetPetSuccess = []struct {
	TestName  string
	InputData struct {
		petId uint64
	}
	Prepare     func(fields *petServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "get pet success",
		InputData: struct {
			petId uint64
		}{petId: 1},

		Prepare: func(fields *petServiceFields) {
			fields.petRepositoryMock.EXPECT().GetPet(uint64(1)).Return(&models.Pet{PetId: 1}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testGetPetFailure = []struct {
	TestName  string
	InputData struct {
		petId uint64
	}
	Prepare     func(fields *petServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "get pet failure, entity does not exists",
		InputData: struct {
			petId uint64
		}{petId: uint64(1)},

		Prepare: func(fields *petServiceFields) {
			fields.petRepositoryMock.EXPECT().GetPet(uint64(1)).Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.PetDoesNotExists)
		},
	},
}

func TestPetServiceImplementationGet(t *testing.T) {
	t.Parallel()

	for _, tt := range testGetPetSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createPetServiceFields(ctrl)
			tt.Prepare(fields)

			petService := createPetService(fields)

			_, err := petService.GetPet(tt.InputData.petId)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testGetPetFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createPetServiceFields(ctrl)
			tt.Prepare(fields)

			petService := createPetService(fields)

			_, err := petService.GetPet(tt.InputData.petId)

			tt.CheckOutput(t, err)

		})
	}
}

//-------------------------------------------------------------------------------------------------
// GetAllByClient

var testGetAllByClientSuccess = []struct {
	TestName  string
	InputData struct {
		id uint64
	}
	Prepare     func(fields *petServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "get all by client success",
		InputData: struct {
			id uint64
		}{id: 1},

		Prepare: func(fields *petServiceFields) {
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return([]models.Pet{}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testGetAllByClientFailure = []struct {
	TestName  string
	InputData struct {
		id uint64
	}
	Prepare     func(fields *petServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "get all by client repo failure",
		InputData: struct {
			id uint64
		}{id: 1},

		Prepare: func(fields *petServiceFields) {
			fields.petRepositoryMock.EXPECT().GetAllByClient(uint64(1)).Return(nil, dbErrors.ErrorSelect)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, dbErrors.ErrorSelect)
		},
	},
}

func TestPetServiceImplementationGetAllByClient(t *testing.T) {
	t.Parallel()

	for _, tt := range testGetAllByClientSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createPetServiceFields(ctrl)
			tt.Prepare(fields)

			petService := createPetService(fields)

			_, err := petService.GetAllByClient(tt.InputData.id)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testGetAllByClientFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createPetServiceFields(ctrl)
			tt.Prepare(fields)

			petService := createPetService(fields)

			_, err := petService.GetAllByClient(tt.InputData.id)

			tt.CheckOutput(t, err)

		})
	}
}

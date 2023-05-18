package servicesImplementation

import (
	"backend/internal/models"
	"backend/internal/pkg/errors/bdErrors"
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

type clientServiceFields struct {
	clientRepositoryMock *mock_repository.MockClientRepository
	hasherMock           *mock_hasher.MockHasher
	logger               *log.Logger
}

func createClientServiceFields(controller *gomock.Controller) *clientServiceFields {
	fields := new(clientServiceFields)

	fields.hasherMock = mock_hasher.NewMockHasher(controller)
	fields.clientRepositoryMock = mock_repository.NewMockClientRepository(controller)
	fields.logger = log.New(os.Stderr)
	fields.logger.SetLevel(log.FatalLevel)

	return fields
}

func createClientService(fields *clientServiceFields) services.ClientService {
	return NewClientServiceImplementation(fields.clientRepositoryMock, fields.hasherMock, fields.logger)
}

//-------------------------------------------------------------------------------------------------
// create

var testClientCreateSuccess = []struct {
	TestName  string
	InputData struct {
		client   *models.Client
		password string
	}
	Prepare     func(fields *clientServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple create",
		InputData: struct {
			client   *models.Client
			password string
		}{client: &models.Client{Login: "Chepigo"}, password: "12345"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").
				Return(nil, repoErrors.EntityDoesNotExists)
			fields.hasherMock.EXPECT().GetHash("12345").Return([]byte("12345_hash"), nil)
			fields.clientRepositoryMock.EXPECT().Create(
				&models.Client{Login: "Chepigo", Password: "12345_hash"}).
				Return(nil)
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").
				Return(&models.Client{Login: "Chepigo", Password: "12345_hash"}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testClientCreateFailure = []struct {
	TestName  string
	InputData struct {
		client   *models.Client
		password string
	}
	Prepare     func(fields *clientServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "create error, error get client by login",
		InputData: struct {
			client   *models.Client
			password string
		}{client: &models.Client{Login: "Chepigo"}, password: "12345"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").
				Return(nil, repoErrors.ErrorGetClientByLogin)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, repoErrors.ErrorGetClientByLogin)
		},
	},
	{
		TestName: "create error, client already exist",
		InputData: struct {
			client   *models.Client
			password string
		}{client: &models.Client{Login: "Chepigo"}, password: "12345"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").Return(&models.Client{}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ClientAlreadyExists)
		},
	},
	{
		TestName: "create error, bad hash",
		InputData: struct {
			client   *models.Client
			password string
		}{client: &models.Client{Login: "Chepigo"}, password: "12345"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").
				Return(nil, repoErrors.EntityDoesNotExists)

			fields.hasherMock.EXPECT().GetHash("12345").Return(nil, serviceErrors.ErrorHash)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ErrorHash)
		},
	},
	{
		TestName: "create error, bad response from repository",
		InputData: struct {
			client   *models.Client
			password string
		}{client: &models.Client{Login: "Chepigo"}, password: "12345"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").
				Return(nil, repoErrors.EntityDoesNotExists)

			fields.hasherMock.EXPECT().GetHash("12345").Return([]byte("12345_hash"), nil)
			fields.clientRepositoryMock.EXPECT().Create(&models.Client{Login: "Chepigo", Password: "12345_hash"}).
				Return(bdErrors.ErrorInsert)

		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, bdErrors.ErrorInsert)
		},
	},
}

func TestClientServiceImplementationCreate(t *testing.T) {
	t.Parallel()

	for _, tt := range testClientCreateSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createClientServiceFields(ctrl)
			tt.Prepare(fields)

			clientService := createClientService(fields)

			_, err := clientService.Create(tt.InputData.client, tt.InputData.password)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testClientCreateFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createClientServiceFields(ctrl)
			tt.Prepare(fields)

			clientService := createClientService(fields)

			_, err := clientService.Create(tt.InputData.client, tt.InputData.password)

			tt.CheckOutput(t, err)

		})
	}
}

//-------------------------------------------------------------------------------------------------
// Login

var testClientLoginSuccess = []struct {
	TestName  string
	InputData struct {
		login    string
		password string
	}
	Prepare     func(fields *clientServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple login",
		InputData: struct {
			login    string
			password string
		}{login: "Chepigo", password: "12345"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").Return(&models.Client{
				ClientId: 1,
				Login:    "Chepigo",
				Password: "12345"}, nil)
			fields.hasherMock.EXPECT().CheckUnhashedValue("12345", "12345").Return(true)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testClientLoginFailure = []struct {
	TestName  string
	InputData struct {
		login    string
		password string
	}
	Prepare     func(fields *clientServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "login failure, entity does not exist",
		InputData: struct {
			login    string
			password string
		}{login: "Chepigo", password: "12345"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").
				Return(nil, repoErrors.EntityDoesNotExists)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.ClientDoesNotExists)
		},
	},
	{
		TestName: "login failure, bad password",
		InputData: struct {
			login    string
			password string
		}{login: "Chepigo", password: "12345"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").Return(&models.Client{
				ClientId: 1,
				Login:    "Chepigo",
				Password: "123456"}, nil)
			fields.hasherMock.EXPECT().CheckUnhashedValue("123456", "12345").Return(false)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, serviceErrors.InvalidPassword)
		},
	},
	{
		TestName: "login failure, repository error response",
		InputData: struct {
			login    string
			password string
		}{login: "Chepigo", password: "12345"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").Return(nil, bdErrors.ErrorSelect)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, bdErrors.ErrorSelect)
		},
	},
}

func TestClientServiceImplementationLogin(t *testing.T) {
	t.Parallel()

	for _, tt := range testClientLoginSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createClientServiceFields(ctrl)
			tt.Prepare(fields)

			clientService := createClientService(fields)

			_, err := clientService.Login(tt.InputData.login, tt.InputData.password)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testClientLoginFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createClientServiceFields(ctrl)
			tt.Prepare(fields)

			clientService := createClientService(fields)

			_, err := clientService.Login(tt.InputData.login, tt.InputData.password)

			tt.CheckOutput(t, err)
		})
	}
}

//-------------------------------------------------------------------------------------------------
// GetClientById(id uint64) (*models.Client, error)

var testGetClientByIdSuccess = []struct {
	TestName  string
	InputData struct {
		id uint64
	}
	Prepare     func(fields *clientServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple get client by id",
		InputData: struct {
			id uint64
		}{id: 1},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientById(uint64(1)).Return(&models.Client{
				ClientId: 1,
				Login:    "Chepigo",
				Password: "12345"}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

var testGetClientByIdFailure = []struct {
	TestName  string
	InputData struct {
		id uint64
	}
	Prepare     func(fields *clientServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "bad repo response get client by id",
		InputData: struct {
			id uint64
		}{id: 1},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientById(uint64(1)).Return(nil,
				bdErrors.ErrorSelect)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.ErrorIs(t, err, bdErrors.ErrorSelect)
		},
	},
}

func TestClientServiceImplementationGetClientById(t *testing.T) {
	t.Parallel()

	for _, tt := range testGetClientByIdSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createClientServiceFields(ctrl)
			tt.Prepare(fields)

			clientService := createClientService(fields)

			_, err := clientService.GetClientById(tt.InputData.id)

			tt.CheckOutput(t, err)
		})
	}

	for _, tt := range testGetClientByIdFailure {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createClientServiceFields(ctrl)
			tt.Prepare(fields)

			clientService := createClientService(fields)

			_, err := clientService.GetClientById(tt.InputData.id)

			tt.CheckOutput(t, err)
		})
	}
}

//-------------------------------------------------------------------------------------------------
// GetClientByLogin(login string) (*models.Client, error)

var testGetClientByLoginSuccess = []struct {
	TestName  string
	InputData struct {
		login string
	}
	Prepare     func(fields *clientServiceFields)
	CheckOutput func(t *testing.T, err error)
}{
	{
		TestName: "simple get client by login",
		InputData: struct {
			login string
		}{login: "Chepigo"},

		Prepare: func(fields *clientServiceFields) {
			fields.clientRepositoryMock.EXPECT().GetClientByLogin("Chepigo").Return(&models.Client{
				ClientId: 1,
				Login:    "Chepigo",
				Password: "12345"}, nil)
		},
		CheckOutput: func(t *testing.T, err error) {
			require.NoError(t, err)
		},
	},
}

func TestClientServiceImplementationGetClientByLogin(t *testing.T) {
	t.Parallel()

	for _, tt := range testGetClientByLoginSuccess {
		tt := tt
		t.Run(tt.TestName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := createClientServiceFields(ctrl)
			tt.Prepare(fields)

			clientService := createClientService(fields)

			_, err := clientService.GetClientByLogin(tt.InputData.login)

			tt.CheckOutput(t, err)
		})
	}
}

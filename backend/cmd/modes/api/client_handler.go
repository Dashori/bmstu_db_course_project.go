package api

import (
	token "backend/cmd/modes/api/utils"
	"backend/internal/models"
	"github.com/gin-gonic/gin"
	"time"
)

func (t *services) createClient(c *gin.Context) {
	var client *models.Client
	err := c.ShouldBindJSON(&client)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	res, err := t.Services.ClientService.Create(client, client.Password)
	if errorHandler(c, err) != true {
		return
	}

	token, err := token.GenerateToken(res.ClientId, "client")
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	jsonClientCreateResponse(c, res, token)
}

func (t *services) loginClient(c *gin.Context) {

	var client *models.Client
	err := c.ShouldBindJSON(&client)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	res, err := t.Services.ClientService.Login(client.Login, client.Password)
	if errorHandler(c, err) != true {
		return
	}

	token, err := token.GenerateToken(res.ClientId, "client")

	if err != nil {
		jsonBadRequestResponse(c, err)
		return
	}

	jsonClientLoginOkResponse(c, res, token)
}

func (t *services) infoClient(c *gin.Context) {

	user_id, role, err := token.ExtractTokenIdAndRole(c)
	if errorHandlerClientAuth(c, err, role) != true {
		return
	}

	res, err := t.Services.ClientService.GetClientById(user_id)
	if errorHandler(c, err) != true {
		return
	}

	jsonClientInfoOkResponse(c, res)
}

func (t *services) ClientRecords(c *gin.Context) {

	user_id, role, err := token.ExtractTokenIdAndRole(c)
	if errorHandlerClientAuth(c, err, role) != true {
		return
	}

	client, err := t.Services.ClientService.GetClientById(user_id)
	if errorHandler(c, err) != true {
		return
	}

	records, err := t.Services.RecordService.GetAllRecords(0, client.ClientId)
	if errorHandler(c, err) != true {
		return
	}

	jsonRecordsOkResponse(c, records)
}

func (t *services) ClientPets(c *gin.Context) {

	user_id, role, err := token.ExtractTokenIdAndRole(c)
	if errorHandlerClientAuth(c, err, role) != true {
		return
	}

	client, err := t.Services.ClientService.GetClientById(user_id)
	if errorHandler(c, err) != true {
		return
	}

	pets, err := t.Services.PetService.GetAllByClient(client.ClientId)
	if errorHandler(c, err) != true {
		return
	}

	jsonPetsOkResponse(c, pets)
}

func (t *services) NewPet(c *gin.Context) {

	user_id, role, err := token.ExtractTokenIdAndRole(c)
	if errorHandlerClientAuth(c, err, role) != true {
		return
	}

	client, err := t.Services.ClientService.GetClientById(user_id)
	if errorHandler(c, err) != true {
		return
	}

	var pet *models.Pet
	err = c.ShouldBindJSON(&pet)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	err = t.Services.PetService.Create(pet, client.Login)
	if errorHandler(c, err) != true {
		return
	}

	jsonPetCreatedResponse(c, pet)
}

func (t *services) DeletePet(c *gin.Context) {

	user_id, role, err := token.ExtractTokenIdAndRole(c)
	if errorHandlerClientAuth(c, err, role) != true {
		return
	}

	client, err := t.Services.ClientService.GetClientById(user_id)
	if errorHandler(c, err) != true {
		return
	}

	var pet *models.Pet
	err = c.ShouldBindJSON(&pet)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	err = t.Services.PetService.Delete(uint64(pet.PetId), client.ClientId)
	if errorHandler(c, err) != true {
		return
	}

	jsonPetOkResponse(c, pet)
}

type Record struct {
	PetId     int
	DoctorId  int
	Year      int
	Month     int
	Day       int
	StartTime int
}

func (t *services) NewRecord(c *gin.Context) {

	user_id, role, err := token.ExtractTokenIdAndRole(c)
	if errorHandlerClientAuth(c, err, role) != true {
		return
	}

	client, err := t.Services.ClientService.GetClientById(user_id)
	if errorHandler(c, err) != true {
		return
	}

	var record Record
	err = c.ShouldBindJSON(&record)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	datetimeStart := time.Date(record.Year, time.Month(record.Month), record.Day, record.StartTime, 00, 00, 00, time.UTC)
	datetimeEnd := time.Date(record.Year, time.Month(record.Month), record.Day, record.StartTime+1, 00, 00, 00, time.UTC)

	newRecord := models.Record{PetId: uint64(record.PetId), ClientId: uint64(client.ClientId),
		DoctorId: uint64(record.DoctorId), DatetimeStart: datetimeStart, DatetimeEnd: datetimeEnd}

	err = t.Services.RecordService.CreateRecord(&newRecord)
	if errorHandler(c, err) != true {
		return
	}

	jsonRecordCreatedResponse(c, &newRecord)
}

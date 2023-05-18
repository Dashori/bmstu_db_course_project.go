package api

import (
	"backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonStatusOkResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// errors

func jsonInternalServerErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func jsonBadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func jsonUnauthorizedResponse(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
}

func jsonBadRoleResponse(c *gin.Context, role string) {
	c.JSON(http.StatusForbidden, gin.H{"error": role})
}

// client

func jsonClientCreateResponse(c *gin.Context, client *models.Client, token string) {
	c.JSON(http.StatusCreated, gin.H{"ClientId": client.ClientId, "Login": client.Login, "Token": token})
}

func jsonClientLoginOkResponse(c *gin.Context, client *models.Client, token string) {
	c.JSON(http.StatusOK, gin.H{"ClientId": client.ClientId, "Login": client.Login, "Token": token})
}

func jsonClientInfoOkResponse(c *gin.Context, client *models.Client) {
	c.JSON(http.StatusOK, gin.H{"ClientId": client.ClientId, "Login": client.Login})
}

// doctor

func jsonDoctorCreateResponse(c *gin.Context, doctor *models.Doctor, token string) {
	c.JSON(http.StatusCreated, gin.H{"DoctorId": doctor.DoctorId, "Login": doctor.Login,
		"StartTime": doctor.StartTime, "EndTime": doctor.EndTime, "Token": token})
}

func jsonDoctorLoginOkResponse(c *gin.Context, doctor *models.Doctor, token string) {
	c.JSON(http.StatusOK, gin.H{"DoctorId": doctor.DoctorId, "Login": doctor.Login, "Token": token})
}

func jsonDoctorInfoOkResponse(c *gin.Context, doctor *models.Doctor) {
	c.JSON(http.StatusOK, gin.H{"DoctorId": doctor.DoctorId, "Login": doctor.Login,
		"StartTime": doctor.StartTime, "EndTime": doctor.EndTime})
}

func jsonDoctorsOkResponse(c *gin.Context, doctors []models.Doctor) {
	c.JSON(http.StatusOK, gin.H{"doctors": doctors})
}

// records

func jsonRecordsOkResponse(c *gin.Context, records []models.Record) {
	c.JSON(http.StatusOK, gin.H{"records": records})
}

func jsonRecordCreatedResponse(c *gin.Context, record *models.Record) {
	c.JSON(http.StatusCreated, gin.H{"record": record})
}

// pets

func jsonPetsOkResponse(c *gin.Context, pets []models.Pet) {
	c.JSON(http.StatusOK, gin.H{"pets": pets})
}

func jsonPetCreatedResponse(c *gin.Context, pet *models.Pet) {
	c.JSON(http.StatusCreated, gin.H{"pet": pet})
}

func jsonPetOkResponse(c *gin.Context, pet *models.Pet) {
	c.JSON(http.StatusOK, gin.H{"pet": pet})
}

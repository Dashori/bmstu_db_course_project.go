package api

import (
	token "backend/cmd/modes/api/utils"
	"backend/internal/models"
	"github.com/gin-gonic/gin"
)

func (t *services) createDoctor(c *gin.Context) {
	var doctor *models.Doctor
	err := c.ShouldBindJSON(&doctor)
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	res, err := t.Services.DoctorService.Create(doctor, doctor.Password)
	if errorHandler(c, err) != true {
		return
	}

	token, err := token.GenerateToken(res.DoctorId, "doctor")
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	jsonDoctorCreateResponse(c, res, token)
}

func (t *services) loginDoctor(c *gin.Context) {
	var doctor *models.Doctor
	err := c.ShouldBindJSON(&doctor)
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	res, err := t.Services.DoctorService.Login(doctor.Login, doctor.Password)
	if errorHandler(c, err) != true {
		return
	}

	token, err := token.GenerateToken(res.DoctorId, "doctor")
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	jsonDoctorLoginOkResponse(c, res, token)
}

func (t *services) getAllDoctors(c *gin.Context) {

	doctors, err := t.Services.DoctorService.GetAllDoctors()
	if errorHandler(c, err) != true {
		return
	}

	jsonDoctorsOkResponse(c, doctors)
}

func (t *services) doctorInfo(c *gin.Context) {

	user_id, role, err := token.ExtractTokenIdAndRole(c)
	if errorHandlerDoctorAuth(c, err, role) != true {
		return
	}

	res, err := t.Services.DoctorService.GetDoctorById(user_id)
	if errorHandler(c, err) != true {
		return
	}

	jsonDoctorInfoOkResponse(c, res)
}

func (t *services) doctorRecords(c *gin.Context) {

	user_id, role, err := token.ExtractTokenIdAndRole(c)
	if errorHandlerDoctorAuth(c, err, role) != true {
		return
	}

	doctor, err := t.Services.DoctorService.GetDoctorById(user_id)
	if errorHandler(c, err) != true {
		return
	}

	records, err := t.Services.RecordService.GetAllRecords(doctor.DoctorId, 0)
	if errorHandler(c, err) != true {
		return
	}

	jsonRecordsOkResponse(c, records)
}

type newTime struct {
	StartTime uint64
	EndTime   uint64
}

func (t *services) doctorShedule(c *gin.Context) {
	var shedule *newTime
	err := c.ShouldBindJSON(&shedule)
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	user_id, role, err := token.ExtractTokenIdAndRole(c)
	if errorHandlerDoctorAuth(c, err, role) != true {
		return
	}

	doctor, err := t.Services.DoctorService.GetDoctorById(user_id)
	if errorHandler(c, err) != true {
		return
	}

	err = t.Services.DoctorService.UpdateShedule(doctor.DoctorId,
		uint64(shedule.StartTime), uint64(shedule.EndTime))
	if errorHandler(c, err) != true {
		return
	}

	jsonStatusOkResponse(c)
}

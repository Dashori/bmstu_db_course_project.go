package api

import (
	"github.com/gin-gonic/gin"

	dbErrors "backend/internal/pkg/errors/dbErrors"
	servicesErrors "backend/internal/pkg/errors/servicesErrors"
)

func errorHandler(c *gin.Context, err error) bool {

	if err == nil {
		return true
	}

	if err == dbErrors.ErrorParseConfig ||
		err == dbErrors.ErrorInitDB ||
		err == dbErrors.ErrorInsert ||
		err == dbErrors.ErrorDelete ||
		err == dbErrors.ErrorSelect ||
		err == dbErrors.ErrorUpdate ||
		err == dbErrors.ErrorCopy {

		jsonInternalServerErrorResponse(c, err)
		return false
	}

	if err == servicesErrors.ErrorCreateRecordTime ||
		err == servicesErrors.TimeIsTaken ||
		err == servicesErrors.ErrorDoctorTime ||
		err == servicesErrors.NotClientPet ||
		err == servicesErrors.DoctorDoesNotExists ||
		err == servicesErrors.ClientDoesNotExists ||
		err == servicesErrors.RecordDoesNotExists ||
		err == servicesErrors.PetDoesNotExists ||
		err == servicesErrors.PetAlreadyExists ||
		err == servicesErrors.DoctorAlreadyExists ||
		err == servicesErrors.ClientAlreadyExists ||
		err == servicesErrors.ErrorHash ||
		err == servicesErrors.InvalidPassword ||
		err == servicesErrors.ErrorWrongNewShedule {

		jsonBadRequestResponse(c, err)
		return false
	}

	return true
}

func errorHandlerClientAuth(c *gin.Context, err error, role string) bool {

	if err != nil {
		jsonUnauthorizedResponse(c, err)
		return false
	}

	if role != "client" {
		jsonBadRoleResponse(c, role)
		return false
	}

	return true
}

func errorHandlerDoctorAuth(c *gin.Context, err error, role string) bool {

	if err != nil {
		jsonUnauthorizedResponse(c, err)
		return false
	}

	if role != "doctor" {
		jsonBadRoleResponse(c, role)
		return false
	}

	return true
}

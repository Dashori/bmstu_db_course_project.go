package api

import (
	registry "backend/cmd/registry"

	"backend/cmd/modes/api/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

type services struct {
	Services *registry.AppServiceFields
}

func SetupServer(a *registry.App) *gin.Engine {
	t := services{a.Services}

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/setRole", t.setRole)

		api.GET("/doctors", t.getAllDoctors)
		api.POST("/doctor/create", t.createDoctor)
		api.POST("/doctor/login", t.loginDoctor)

		doctor := api.Group("/doctor")
		doctor.Use(middlewares.JwtAuthMiddleware())
		doctor.GET("/info", t.doctorInfo)
		doctor.GET("/records", t.doctorRecords)
		doctor.PATCH("/shedule", t.doctorShedule)

		api.POST("/client/create", t.createClient)
		api.POST("/client/login", t.loginClient)

		client := api.Group("/client")
		client.Use(middlewares.JwtAuthMiddleware())
		client.GET("/info", t.infoClient)
		client.GET("/records", t.ClientRecords)
		client.GET("/pets", t.ClientPets)
		client.POST("/record", t.NewRecord)
		client.POST("/pet", t.NewPet)
		client.DELETE("/pet", t.DeletePet)
	}

	port := a.Config.Port
	adress := a.Config.Address
	err := router.Run(adress + port)

	if err != nil {
		panic(err)
	}

	return router
}

type Role struct {
	Role string
}

func (t *services) setRole(c *gin.Context) {
	var role *Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	if role.Role == "doctor" {
		err = t.Services.DoctorService.SetRole()
	} else if role.Role == "client" {
		err = t.Services.ClientService.SetRole()
	} else {
		jsonBadRequestResponse(c, fmt.Errorf("Такой роли не существует!"))
	}

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	jsonStatusOkResponse(c)

}

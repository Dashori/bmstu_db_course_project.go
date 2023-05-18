package api

import (
	registry "backend/cmd/registry"

	"backend/cmd/modes/api/middlewares"
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
	router.Run(adress + port)

	return router
}

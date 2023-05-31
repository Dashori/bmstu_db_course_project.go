package main

import (
	"backend/internal/models"
	servicesImplementation "backend/internal/services/implementation"
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"time"
)

const N = 20

func main() {

	err := researchCreateRecordWithTrigger()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All is ok!")
	}

	err = researchCreateRecordWithoutTrigger()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All is ok!")
	}
}

func researchCreateRecordWithTrigger() error {

	dbContainer, db := SetupTestDatabase("../db/postgreSQL/init.sql")
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := servicesImplementation.СreateRecordServiceFieldsPostgres(db)
	records := servicesImplementation.CreateRecordServicePostgres(fields)

	clients := fields.ClientRepository
	doctors := fields.DoctorRepository
	pets := fields.PetRepository

	// это с триггером
	err := (*clients).Create(&models.Client{Login: "ChicagoTest", Password: "12345"})
	if err != nil {
		return err
	}

	client, err := (*clients).GetClientByLogin("ChicagoTest")
	if err != nil {
		return err
	}

	err = (*doctors).Create(&models.Doctor{Login: "ChicagoTest", Password: "12345", StartTime: 10, EndTime: 23})
	if err != nil {
		return err
	}

	doctor, err := (*doctors).GetDoctorByLogin("ChicagoTest")
	if err != nil {
		return err
	}

	var result int64

	for i := 0; i < N; i++ {

		err = (*pets).Create(&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: client.ClientId})
		if err != nil {
			return err
		}

		// трюк чтоб узнать id питомца Havrosha
		clientPets, err := (*pets).GetAllByClient(client.ClientId)
		if err != nil {
			return err
		}

		petId := clientPets[0].PetId

		err, duration := records.CreateRecordResearchTrigger(&models.Record{
			PetId: petId, ClientId: client.ClientId, DoctorId: doctor.DoctorId,
			DatetimeStart: time.Date(2024, 7, 7, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(2024, 7, 7, 16, 00, 00, 00, time.UTC)})
	
		if err != nil {
			return err
		}

		result += duration.Nanoseconds()
		fmt.Println("время!!!!! ", duration.Nanoseconds())

		err = (*pets).Delete(petId) // при удалении pet удалится и запись в records

		if err != nil {
			return err
		}
	}

	// fmt.Println("время!!!!! ", result.Nanoseconds())
	fmt.Println("итог время!!!!! ", result/N)
	return err
}


func researchCreateRecordWithoutTrigger() error {

	dbContainer, db := SetupTestDatabase("../db/postgreSQL/research.sql")
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := servicesImplementation.СreateRecordServiceFieldsPostgres(db)
	records := servicesImplementation.CreateRecordServicePostgres(fields)

	clients := fields.ClientRepository
	doctors := fields.DoctorRepository
	pets := fields.PetRepository

	// это с триггером
	err := (*clients).Create(&models.Client{Login: "ChicagoTest", Password: "12345"})
	if err != nil {
		return err
	}

	client, err := (*clients).GetClientByLogin("ChicagoTest")
	if err != nil {
		return err
	}

	err = (*doctors).Create(&models.Doctor{Login: "ChicagoTest", Password: "12345", StartTime: 10, EndTime: 23})
	if err != nil {
		return err
	}

	doctor, err := (*doctors).GetDoctorByLogin("ChicagoTest")
	if err != nil {
		return err
	}

	var result int64

	for i := 0; i < N; i++ {

		err = (*pets).Create(&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: client.ClientId})
		if err != nil {
			return err
		}

		// трюк чтоб узнать id питомца Havrosha
		clientPets, err := (*pets).GetAllByClient(client.ClientId)
		if err != nil {
			return err
		}

		petId := clientPets[0].PetId

		err, duration := records.CreateRecordResearch(&models.Record{
			PetId: petId, ClientId: client.ClientId, DoctorId: doctor.DoctorId,
			DatetimeStart: time.Date(2024, 7, 7, 15, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(2024, 7, 7, 16, 00, 00, 00, time.UTC)})
	
		if err != nil {
			return err
		}

		result += duration.Nanoseconds()
		fmt.Println("время!!!!! ", duration.Nanoseconds())

		err = (*pets).Delete(petId) // при удалении pet удалится и запись в records

		if err != nil {
			return err
		}
	}

	fmt.Println("итог время!!!!! ", result/N)
	return err
}

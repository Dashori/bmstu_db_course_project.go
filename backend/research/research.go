package main

import (
	"backend/internal/models"
	servicesImplementation "backend/internal/services/implementation"
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"time"
	"os"
	"math/rand"
)

const REPEAT = 20
const N_start = 10
const N_end = 10

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
	// doctors := fields.DoctorRepository
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


	var result int64

	// for i := N_start; i <= N_end; i+= 10 {

		doctorId := uint64(rand.Intn(10))

			
		text, err := os.ReadFile("../db/postgreSQL/10.sql")
		if err != nil {
			return err
		}

		if _, err := db.Exec(string(text)); err != nil {
			return err
		}

		for i := 0; i < REPEAT; i++ {

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
				PetId: petId, ClientId: client.ClientId, DoctorId: doctorId,
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
	// }

	// fmt.Println("время!!!!! ", result.Nanoseconds())
	fmt.Println("итог время!!!!! ", result/REPEAT)
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

	if db == nil {
		return nil
	}

	fields := servicesImplementation.СreateRecordServiceFieldsPostgres(db)
	records := servicesImplementation.CreateRecordServicePostgres(fields)

	clients := fields.ClientRepository
	// doctors := fields.DoctorRepository
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

	doctorId := uint64(rand.Intn(10))

	var result int64
	
				
	text, err := os.ReadFile("../db/postgreSQL/10.sql")
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(text)); err != nil {
		return err
	}

	for i := 0; i < REPEAT; i++ {

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
			PetId: petId, ClientId: client.ClientId, DoctorId: doctorId,
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

	fmt.Println("итог время!!!!! ", result/REPEAT)
	return err
}

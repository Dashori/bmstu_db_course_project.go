package main

import (
	"backend/internal/models"
	servicesImplementation "backend/internal/services/implementation"
	"context"
	"database/sql"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const N = 200

func main() {
	step := 0

	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 10; i <= 1000; i += step {
		fmt.Println("Records count: ", i)

		resultTimeTr, errorCountTr, err := researchCreateRecordWithTrigger(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Research with trigger end ok!")
		}

		resultTime, errorCount, err := researchCreateRecordWithoutTrigger(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Research without trigger end ok!")
		}

		_, err = file.WriteString(strconv.Itoa(i) + " " + strconv.Itoa(resultTimeTr) + " " + strconv.Itoa(errorCountTr) + " ")
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = file.WriteString(strconv.Itoa(resultTime) + " " + strconv.Itoa(errorCount) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}

		if i < 100 {
			i += 10
		} else if i == 100 {
			step = 100
		}
	}
}

func setupData(count int) (error, *sql.DB, testcontainers.Container) {
	dbContainer, db := SetupTestDatabase("../db/postgreSQL/research.sql")

	var path string = "../db/postgreSQL/" + strconv.Itoa(count) + ".sql"

	text, err := os.ReadFile(path)
	if err != nil {
		return err, nil, nil
	}

	if _, err := db.Exec(string(text)); err != nil {
		return err, nil, nil
	}

	return nil, db, dbContainer
}

func researchCreateRecordWithTrigger(count int) (int, int, error) {

	err, db, dbContainer := setupData(count)
	if err != nil {
		return 0, 0, err
	}
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	text, err := os.ReadFile("../db/postgreSQL/tr.sql")
	if err != nil {
		return 0, 0, err
	}

	if _, err := db.Exec(string(text)); err != nil {
		return 0, 0, err
	}

	fields := servicesImplementation.СreateRecordServiceFieldsPostgres(db)
	records := servicesImplementation.CreateRecordServicePostgres(fields)
	clients := fields.ClientRepository
	pets := fields.PetRepository

	err = (*clients).Create(&models.Client{Login: "ChicagoTest", Password: "12345"})
	if err != nil {
		return 0, 0, err
	}

	client, err := (*clients).GetClientByLogin("ChicagoTest")
	if err != nil {
		return 0, 0, err
	}

	var result int64
	var errorCount int64
	var successCount int64

	// for i := 0; i < N; i++ {
	for successCount != N{
	
		var doctorId = uint64(rand.Intn(30) + 1) // [0,n)
		var month = time.Month(rand.Intn(12) + 1)
		var day = rand.Intn(28) + 1
		var hour = rand.Intn(22)

		err = (*pets).Create(&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: client.ClientId})
		if err != nil {
			return 0, 0, err
		}

		// трюк чтоб узнать id питомца Havrosha
		clientPets, err := (*pets).GetAllByClient(client.ClientId)
		if err != nil {
			return 0, 0, err
		}

		petId := clientPets[0].PetId

		err, duration := records.CreateRecordResearchTrigger(&models.Record{
			PetId: petId, ClientId: client.ClientId, DoctorId: doctorId,
			DatetimeStart: time.Date(2030, month, day, hour, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(2030, month, day, hour+1, 00, 00, 00, time.UTC)})

		if err != nil {
			errorCount += 1
		} else {
			successCount += 1
		}

		result += duration.Nanoseconds()

		err = (*pets).Delete(petId) // при удалении pet удалится и запись в records

		if err != nil {
			return 0, 0, err
		}
	}

	fmt.Println("итог время!!!!! ", result/(N-errorCount))
	fmt.Println("итого ошибок!!!!", errorCount)
	return int(result), int(errorCount), err
}

func researchCreateRecordWithoutTrigger(count int) (int, int, error) {
	err, db, dbContainer := setupData(count)
	if err != nil {
		return 0, 0, err
	}
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}(dbContainer, context.Background())

	fields := servicesImplementation.СreateRecordServiceFieldsPostgres(db)
	records := servicesImplementation.CreateRecordServicePostgres(fields)
	clients := fields.ClientRepository
	pets := fields.PetRepository

	err = (*clients).Create(&models.Client{Login: "ChicagoTest", Password: "12345"})
	if err != nil {
		return 0, 0, err
	}

	client, err := (*clients).GetClientByLogin("ChicagoTest")
	if err != nil {
		return 0, 0, err
	}
	var result int64
	var errorCount int64
	var successCount int64

	for successCount != N {

	// for i := 0; i < N; i++ {

		var doctorId = uint64(rand.Intn(30) + 1)
		var month = time.Month(rand.Intn(12) + 1)
		var day = rand.Intn(28) + 1
		var hour = rand.Intn(22)

		err = (*pets).Create(&models.Pet{Name: "Havrosha", Type: "cat", Age: 1, Health: 10, ClientId: client.ClientId})
		if err != nil {
			return 0, 0, err
		}

		// трюк чтоб узнать id питомца Havrosha
		clientPets, err := (*pets).GetAllByClient(client.ClientId)
		if err != nil {
			return 0, 0, err
		}

		petId := clientPets[0].PetId

		err, duration := records.CreateRecordResearch(&models.Record{
			PetId: petId, ClientId: client.ClientId, DoctorId: doctorId,
			DatetimeStart: time.Date(2030, month, day, hour, 00, 00, 00, time.UTC),
			DatetimeEnd:   time.Date(2030, month, day, hour+1, 00, 00, 00, time.UTC)})

		if err != nil {
			errorCount += 1
		} else {
			successCount++
		}

		result += duration.Nanoseconds()

		err = (*pets).Delete(petId) // при удалении pet удалится и запись в records
		if err != nil {
			return 0, 0, err
		}
	}

	fmt.Println("итог время!!!!! ", result/(N-errorCount))
	fmt.Println("ошибок!!!!! ", errorCount)
	return int(result), int(errorCount), err
}

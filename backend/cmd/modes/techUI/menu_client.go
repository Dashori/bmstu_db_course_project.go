package techUI

import (
	"backend/cmd/registry"
	"backend/internal/models"
	"backend/internal/pkg/errors/cliErrors"
	"fmt"
	"time"
)

func loginClient(a *registry.AppServiceFields) (*models.Client, error) {
	var login string
	fmt.Printf("Введите логин: ")
	_, err := fmt.Scanf("%s", &login)
	if err != nil {
		return nil, cliErrors.ErrorInput
	}

	var password string
	fmt.Printf("Введите пароль: ")
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		return nil, cliErrors.ErrorInput
	}

	client, err := a.ClientService.Login(login, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func createClient(a *registry.AppServiceFields) (*models.Client, error) {
	var login string
	fmt.Printf("Введите логин: ")
	_, err := fmt.Scanf("%s", &login)
	if err != nil {
		return nil, cliErrors.ErrorInput
	}

	var password string
	fmt.Printf("Введите пароль: ")
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		return nil, cliErrors.ErrorInput
	}

	client, err := a.ClientService.Create(&models.Client{Login: login, Password: password}, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getPets(a *registry.AppServiceFields, client *models.Client) error {
	pets, err := a.PetService.GetAllByClient(client.ClientId)
	if err != nil {
		return err
	}

	printPets(pets)

	return nil
}

func addPet(a *registry.AppServiceFields, client *models.Client) error {
	fmt.Printf("Введите кличку питомца: ")
	var name string
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		return cliErrors.ErrorInput
	}

	fmt.Printf("Введите тип питомца: ")
	var typePet string
	_, err = fmt.Scanf("%s", &typePet)
	if err != nil {
		return cliErrors.ErrorInput
	}

	fmt.Printf("Введите возраст питомца: ")
	var age uint64
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		return cliErrors.ErrorInput
	}

	fmt.Printf("Введите уровень здоровья питомца: ")
	var health uint64
	_, err = fmt.Scanf("%d", &health)
	if err != nil {
		return cliErrors.ErrorInput
	}

	pet := models.Pet{Name: name, Type: typePet, Age: age, Health: health}

	return a.PetService.Create(&pet, client.Login)
}

func deletePet(a *registry.AppServiceFields, client *models.Client) error {
	fmt.Printf("\nВведите id питомца для удаления: ")
	var id uint64
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return cliErrors.ErrorInput
	}

	err = a.PetService.Delete(id, client.ClientId)
	return err
}

func addRecord(a *registry.AppServiceFields, client *models.Client) error {
	fmt.Printf("\nНапоминаем, что время 1 приема к врачу длится 1 час и начинается в 00.\nВы можете записаться к доктору только в его свободное время.\n")
	fmt.Printf("Введите id питомца: ")
	var petId uint64
	_, err := fmt.Scanf("%d", &petId)
	if err != nil {
		return cliErrors.ErrorInput
	}

	fmt.Printf("Введите id доктора: ")
	var doctorId uint64
	_, err = fmt.Scanf("%d", &doctorId)
	if err != nil {
		return cliErrors.ErrorInput
	}

	fmt.Printf("Введите дату для приема (год, месяц, день через пробел цифрами): ")
	var year, month, day int
	_, err = fmt.Scanf("%d %d %d", &year, &month, &day)
	if err != nil {
		return cliErrors.ErrorInput
	}

	fmt.Printf("Введите время начала приема: ")
	var start, end int
	_, err = fmt.Scanf("%d", &start)
	if err != nil {
		return cliErrors.ErrorInput
	}

	end = start + 1

	datetimeStart := time.Date(year, time.Month(month), day, start, 00, 00, 00, time.UTC)
	datetimeEnd := time.Date(year, time.Month(month), day, end, 00, 00, 00, time.UTC)

	record := models.Record{PetId: petId, ClientId: client.ClientId, DoctorId: doctorId,
		DatetimeStart: datetimeStart, DatetimeEnd: datetimeEnd}

	return a.RecordService.CreateRecord(&record)
}

func printInfoClient(client *models.Client) {
	fmt.Printf("\nЛогин: %s\nId: %d\n\n", client.Login, client.ClientId)
}

func clientLoop(a *registry.AppServiceFields, client *models.Client) error {
	var num int = 1

	for num != 0 {
		fmt.Printf(`Меню клиента: 
0 -- выйти
1 -- посмотреть свои записи на прием
2 -- посмотреть список своих питомцев
3 -- посмотреть список врачей и их расписание
4 -- добавить питомца
5 -- удалить питомца
6 -- записаться на прием
7 -- узнать информацию о себе 
Выберите действие: `)

		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			return cliErrors.ErrorInput
		}

		if num == 0 {
			return nil
		}

		client, err = a.ClientService.GetClientByLogin(client.Login)
		if err != nil {
			return err
		}

		switch num {
		case 1:
			records, err := a.RecordService.GetAllRecords(0, client.ClientId)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				printRecords(records)
			}
		case 2:
			err = getPets(a, client)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			doctors, err := getDoctors(a)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				printDoctors(doctors)
			}
		case 4:
			err = addPet(a, client)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nПитомец успешно добавлен!\n\n")
			}
		case 5:
			err = deletePet(a, client)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nПитомец успешно удален!\n\n")
			}
		case 6:
			err = addRecord(a, client)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nЗапись успешно добавлена!\n\n")
			}
		case 7:
			printInfoClient(client)
		default:
			return cliErrors.ErrorCase
		}
	}

	return nil
}

func clientMenu(a *registry.AppServiceFields) error {
	startMenu :=
		`Меню: 
0 -- выйти
1 -- войти
2 -- зарегестрироваться
Выберите действие: `

	fmt.Printf("%s", startMenu)
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		return cliErrors.ErrorInput
	}

	var client *models.Client

	err = a.ClientService.SetRole()
	if err != nil {
		return err
	}

	switch num {
	case 0:
		return nil
	case 1:
		client, err = loginClient(a)
		if err != nil {
			return err
		}
		fmt.Printf("\nАвторизация успешна!\n\n")
	case 2:
		client, err = createClient(a)
		if err != nil {
			return err
		}
		fmt.Printf("\nКлиент успешно добавлен!\n\n")
	default:
		return cliErrors.ErrorCase
	}

	return clientLoop(a, client)
}

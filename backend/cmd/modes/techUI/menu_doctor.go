package techUI

import (
	registry "backend/cmd/registry"
	"backend/internal/models"
	"backend/internal/pkg/errors/cliErrors"
	"fmt"
)

func createDoctor(a *registry.AppServiceFields) (*models.Doctor, error) {
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

	var start, end uint64
	fmt.Printf("Введите через пробел начало и конец рабочего дня: ")
	_, err = fmt.Scanf("%d %d", &start, &end)
	if err != nil {
		return nil, cliErrors.ErrorInput
	}

	doctor, err := a.DoctorService.Create(&models.Doctor{Login: login, Password: password,
		StartTime: start, EndTime: end}, password)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nДоктор %s успешно зарегестрирован!\n\n", login)

	return doctor, nil
}

func loginDoctor(a *registry.AppServiceFields) (*models.Doctor, error) {
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

	doctor, err := a.DoctorService.Login(login, password)
	if err != nil {
		return nil, err
	}

	return doctor, nil
}

func printInfo(doctor *models.Doctor) {
	fmt.Printf("\nЛогин: %s\nЕжедневное время начала приема: %d\nЕжедневное время конца приема: %d\n\n",
		doctor.Login, doctor.StartTime, doctor.EndTime)
}

func updateShedule(a *registry.AppServiceFields, id uint64) error {
	var start, end uint64

	fmt.Printf("Введите новое время начала и конца приема через пробел: ")
	_, err := fmt.Scanf("%d %d", &start, &end)
	if err != nil {
		return cliErrors.ErrorInput
	}

	return a.DoctorService.UpdateShedule(id, uint64(start), uint64(end))
}

func doctorLoop(a *registry.AppServiceFields, doctor *models.Doctor) error {
	var num int = 1

	for num != 0 {
		fmt.Printf(`Меню доктора: 
0 -- выйти
1 -- посмотреть свои записи
2 -- изменить расписание
3 -- узнать информацию о себе 
Выберите действие: `)

		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			return cliErrors.ErrorInput
		}

		if num == 0 {
			return nil
		}

		doctor, err = a.DoctorService.GetDoctorByLogin(doctor.Login)
		if err != nil {
			return err
		}

		switch num {
		case 1:
			records, err := a.RecordService.GetAllRecords(doctor.DoctorId, 0)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				printRecords(records)
			}
		case 2:
			err = updateShedule(a, doctor.DoctorId)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nРасписание успешно изменено!\n\n")
			}
		case 3:
			printInfo(doctor)
		default:
			return cliErrors.ErrorCase
		}
	}

	return nil
}

func doctorMenu(a *registry.AppServiceFields) error {
	startMenu :=
		`Меню: 
0 -- выйти
1 -- войти
2 -- зарегестрироваться
Выберите действие: `

	fmt.Printf(startMenu)
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		return cliErrors.ErrorInput
	}

	var doctor *models.Doctor
	switch num {
	case 0:
		return nil
	case 1:
		doctor, err = loginDoctor(a)
		if err != nil {
			return err
		} else if err == nil {
			fmt.Printf("\nАвторизация успешна!\n\n")
		}
	case 2:
		doctor, err = createDoctor(a)
		if err != nil {
			return err
		}
	default:
		return cliErrors.ErrorCase
	}

	return doctorLoop(a, doctor)
}

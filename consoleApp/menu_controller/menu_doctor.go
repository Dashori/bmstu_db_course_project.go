package menu_controller

import (
	errors "consoleApp/errors"
	handlers "consoleApp/handlers"
	models "consoleApp/models"
	utils "consoleApp/utils"
	view "consoleApp/view"
	"fmt"
	"net/http"
)

func doctorMenu(client *http.Client) error {

	view.PrintDoctorMenu()
	var num int
	fmt.Scanf("%d", &num)

	var token string
	var err error

	switch num {
	case 0:
		return nil
	case 1:
		token, err = loginDoctor(client)
		if err != nil {
			return err
		} else if err == nil {
			fmt.Printf("\nДоктор успешно авторизован!\n\n")
		}
	default:
		return errors.ErrorCase
	}

	return doctorLoop(client, token)
}

func doctorLoop(client *http.Client, token string) error {
	var num int = 1
	var err error

	for num != 0 {
		view.PrintDoctorLoop()

		fmt.Scanf("%d", &num)

		if num == 0 {
			return nil
		}

		switch num {
		case 1:
			err = getRecordsDoctor(client, token)
			if err != nil {
				return err
			}
		case 2:
			err = updateShedule(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nРасписание успешно изменено!\n\n")
			}
		case 3:
			printInfo(client, token)
		default:
			return errors.ErrorCase
		}
	}

	return nil
}

func loginDoctor(client *http.Client) (string, error) {
	login, password, err := view.InputCred()
	if err != nil {
		return "", err
	}

	newDoctor := models.Doctor{Login: login, Password: password}

	response, err := handlers.LoginDoctor(client, &newDoctor)
	if err == errors.ErrorResponseStatus {
		return "", utils.CheckErrorInBody(response)
	} else if err != nil {
		return "", err
	}

	result, err := utils.ParseDoctorBody(response)
	if err != nil {
		return "", err
	}

	return result.Token, nil
}

func printInfo(client *http.Client, token string) error {
	response, err := handlers.GetInfo(client, token)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	result, err := utils.ParseDoctorBody(response)
	if err != nil {
		return err
	}

	view.PrintDoctorInfo(result)

	return nil
}

func updateShedule(client *http.Client, token string) error {
	newShedule, err := view.InputNewShedule()
	if err != nil {
		return err
	}

	response, err := handlers.UpdateShedule(client, token, newShedule)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	return nil
}

func getRecordsDoctor(client *http.Client, token string) error {
	response, err := handlers.GetRecordsDoctor(client, token)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	records, err := utils.ParseRecordsBody(response)
	if err != nil {
		return err
	}

	view.PrintRecords(records)

	return err
}

func getDoctors(client *http.Client) error {
	response, err := handlers.GetDoctors(client)

	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	result, err := utils.ParseDoctorsBody(response)
	if err != nil {
		return err
	}

	view.PrintDoctors(result)

	return nil
}

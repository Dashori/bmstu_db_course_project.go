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

func setRole(client *http.Client, role string) error {

	response, err := handlers.SetRole(client, role)

	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	return nil
}

func clientMenu(client *http.Client) error {

	view.PrintClientMenu()
	var num int
	fmt.Scanf("%d", &num)

	var token string
	var err error

	switch num {
	case 0:
		return nil
	case 1:
		token, err = loginClient(client)
		if err != nil {
			return err
		}
		fmt.Printf("\nКлиент успешно авторизован!\n\n")
	case 2:
		token, err = createClient(client)
		if err != nil {
			return err
		}
		fmt.Printf("\nКлиент успешно добавлен!\n\n")
	default:
		return errors.ErrorCase
	}

	return clientLoop(client, token)
}

func clientLoop(client *http.Client, token string) error {
	var num int = 1
	var err error

	for num != 0 {
		view.PrintClientLoop()

		fmt.Scanf("%d", &num)

		if num == 0 {
			return nil
		}

		switch num {
		case 1:
			err := getRecords(client, token)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err = getPets(client, token)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			err := getDoctors(client)
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			err = addPet(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nПитомец успешно добавлен!\n\n")
			}
		case 5:
			err = deletePet(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nПитомец успешно удален!\n\n")
			}
		case 6:
			err = addRecord(client, token)
			if err != nil {
				fmt.Println(err)
			} else if err == nil {
				fmt.Printf("\nЗапись успешно добавлена!\n\n")
			}
		case 7:
			err = getInfo(client, token)
			if err != nil {
				fmt.Println(err)
			}
		default:
			return errors.ErrorCase
		}
	}

	return nil
}

func loginClient(client *http.Client) (string, error) {
	login, password, err := view.InputCred()
	if err != nil {
		return "", err
	}
	newClient := models.Client{Login: login, Password: password}

	response, err := handlers.LoginClient(client, &newClient)
	if err == errors.ErrorResponseStatus {
		return "", utils.CheckErrorInBody(response)
	} else if err != nil {
		return "", err
	}

	result, err := utils.ParseClientBody(response)
	if err != nil {
		return "", err
	}

	return result.Token, nil
}

func createClient(client *http.Client) (string, error) {
	login, password, err := view.InputCred()
	if err != nil {
		return "", err
	}
	newClient := models.Client{Login: login, Password: password}

	response, err := handlers.CreateClient(client, &newClient)
	if err == errors.ErrorResponseStatus {
		return "", utils.CheckErrorInBody(response)
	} else if err != nil {
		return "", err
	}

	result, err := utils.ParseClientBody(response)
	if err != nil {
		return "", err
	}

	return result.Token, nil
}

func getPets(client *http.Client, token string) error {
	response, err := handlers.GetClientPets(client, token)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	pets, err := utils.ParsePetsBody(response)
	if err != nil {
		return err
	}

	view.PrintPets(pets)

	return nil
}

func getInfo(client *http.Client, token string) error {
	response, err := handlers.GetClientInfo(client, token)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	result, err := utils.ParseClientBody(response)
	if err != nil {
		return err
	}

	view.PrintClientInfo(result)

	return err
}

func getRecords(client *http.Client, token string) error {
	response, err := handlers.GetClientRecords(client, token)
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

func addPet(client *http.Client, token string) error {
	pet, err := view.InputPetData()
	if err != nil {
		return err
	}

	response, err := handlers.AddPet(client, token, pet)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	return nil
}

func deletePet(client *http.Client, token string) error {
	id, err := view.InputPetId()
	if err != nil {
		return err
	}

	response, err := handlers.DeletePet(client, token, id)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	return nil
}

func addRecord(client *http.Client, token string) error {
	record, err := view.InputRecordData()
	if err != nil {
		return nil
	}

	response, err := handlers.AddRecord(client, token, record)
	if err == errors.ErrorResponseStatus {
		return utils.CheckErrorInBody(response)
	} else if err != nil {
		return err
	}

	return nil
}

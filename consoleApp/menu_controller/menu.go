package menu_controller

import (
	errors "consoleApp/errors"
	view "consoleApp/view"
	"fmt"
	"net/http"
)

func RunMenu(client *http.Client) error {
	view.PrintRunMenu()

	var who int
	fmt.Scanf("%d", &who)

	switch who {
	case 0:
		err := setRole(client, "client")
		if err != nil {
			return err
		}

		err = clientMenu(client)
		if err != nil {
			return err
		}
	case 1:
		err := setRole(client, "doctor")
		if err != nil {
			return err
		}

		err = doctorMenu(client)
		if err != nil {
			return err
		}
	case 2:
		err := getDoctors(client)
		if err != nil {
			return err
		}
	default:
		return errors.ErrorCase
	}

	return nil
}

package menu_controller

import (
	errors "consoleApp/errors"
	view "consoleApp/view"
	"fmt"
	"net/http"
)

const port = "8080"
const adress = "localhost"

func RunMenu(client *http.Client) error {
	view.PrintRunMenu()

	var who int
	fmt.Scanf("%d", &who)

	switch who {
	case 0:
		err := clientMenu(client)
		if err != nil {
			return err
		}
	case 1:
		err := doctorMenu(client)
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

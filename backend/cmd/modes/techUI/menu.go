package techUI

import (
	registry "backend/cmd/registry"
	"backend/internal/pkg/errors/cliErrors"
	"fmt"
)

func RunMenu(a *registry.AppServiceFields) error {
	startPosition :=
		`Кто вы?
0 -- клиент
1 -- работник клиники
2 -- я гость, мне просто врачей посмотреть
Выберите роль: `
	fmt.Printf(startPosition)

	var who int
	fmt.Scanf("%d", &who)

	switch who {
	case 0:
		err := clientMenu(a)
		if err != nil {
			fmt.Println(err)
		}
	case 1:
		err := doctorMenu(a)
		if err != nil {
			fmt.Println(err)
		}
	case 2:
		doctors, err := getDoctors(a)
		if err != nil {
			fmt.Println(err)
		} else if err == nil {
			printDoctors(doctors)
		}
	default:
		fmt.Printf("Неверная роль")
		return cliErrors.ErrorCase
	}

	return nil
}

package main

import (
	menu "consoleApp/menu_controller"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	err := menu.RunMenu(client)
	if err != nil {
		fmt.Println(err)
	}
}

package cliErrors

import "errors"

var (
	ErrorCase  = errors.New("Ошибка в меню! Выбрано неверное действие!")
	ErrorInput = errors.New("Ошибка ввода!")
)

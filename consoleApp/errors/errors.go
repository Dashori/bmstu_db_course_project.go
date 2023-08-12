package errors

import "errors"

var (
	ErrorCase  = errors.New("Ошибка в меню! Выбрано неверное действие!")
	ErrorInput = errors.New("Ошибка ввода!")

	ErrorJson      = errors.New("Ошибка при парсинге json!")
	ErrorHTTP      = errors.New("Неверный http запрос!")
	ErrorReadBody  = errors.New("Ошибка при чтении ответа!")
	ErrorParseBody = errors.New("Ошибка при обработке ответа!")

	ErrorUnknown = errors.New("Неизветсная ошибка при обращении к сервису!")

	ErrorNewRequest     = errors.New("Ошибка при формировании нового запроса!")
	ErrorExecuteRequest = errors.New("Ошибка при выполнении запроса!")
	ErrorResponseStatus = errors.New("Запрос завершился не успешно!")

	ErrorResponse = errors.New("Ошибка при обращении к сервису!")
)

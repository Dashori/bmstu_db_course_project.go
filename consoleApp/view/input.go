package view

import (
	errors "consoleApp/errors"
	models "consoleApp/models"
	"fmt"
)

func InputCred() (string, string, error) {
	var login string
	fmt.Printf("Введите логин: ")
	_, err := fmt.Scanf("%s", &login)
	if err != nil {
		return "", "", errors.ErrorInput
	}

	var password string
	fmt.Printf("Введите пароль: ")
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		return "", "", errors.ErrorInput
	}

	return login, password, nil
}

func InputNewShedule() (models.NewShedule, error) {
	var start, end uint64

	fmt.Printf("Введите новое время начала и конца приема через пробел: ")
	_, err := fmt.Scanf("%d %d", &start, &end)
	if err != nil {
		return models.NewShedule{}, errors.ErrorInput
	}

	return models.NewShedule{Start: start, End: end}, nil
}

func InputPetData() (models.Pet, error) {
	fmt.Printf("Введите кличку питомца: ")
	var name string
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		return models.Pet{}, errors.ErrorInput
	}

	fmt.Printf("Введите тип питомца: ")
	var typePet string
	_, err = fmt.Scanf("%s", &typePet)
	if err != nil {
		return models.Pet{}, errors.ErrorInput
	}

	fmt.Printf("Введите возраст питомца: ")
	var age uint64
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		return models.Pet{}, errors.ErrorInput
	}

	fmt.Printf("Введите уровень здоровья питомца: ")
	var health uint64
	_, err = fmt.Scanf("%d", &health)
	if err != nil {
		return models.Pet{}, errors.ErrorInput
	}

	return models.Pet{Name: name, Type: typePet, Age: age, Health: health}, nil
}

func InputRecordData() (models.Record, error) {
	fmt.Printf("\nНапоминаем, что время 1 приема к врачу длится 1 час и начинается в 00.\nВы можете записаться к доктору только в его свободное время.\n")
	fmt.Printf("Введите id питомца: ")
	var petId uint64
	_, err := fmt.Scanf("%d", &petId)
	if err != nil {
		return models.Record{}, errors.ErrorInput
	}

	fmt.Printf("Введите id доктора: ")
	var doctorId uint64
	_, err = fmt.Scanf("%d", &doctorId)
	if err != nil {
		return models.Record{}, errors.ErrorInput
	}

	fmt.Printf("Введите дату для приема (год, месяц, день через пробел цифрами): ")
	var year, month, day int
	_, err = fmt.Scanf("%d %d %d", &year, &month, &day)
	if err != nil {
		return models.Record{}, errors.ErrorInput
	}

	fmt.Printf("Введите время начала приема: ")
	var start int
	_, err = fmt.Scanf("%d", &start)
	if err != nil {
		return models.Record{}, errors.ErrorInput
	}

	record := models.Record{PetId: petId, DoctorId: doctorId, Year: year, Month: month, Day: day, DatetimeStart: start}
	return record, nil
}

func InputPetId() (uint64, error) {
	fmt.Printf("\nВведите id питомца для удаления: ")
	var id uint64
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		return 0, errors.ErrorInput
	}

	return id, nil
}

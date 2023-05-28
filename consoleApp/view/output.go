package view

import (
	models "consoleApp/models"
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintRunMenu() {
	startPosition :=
		`Кто вы?
0 -- клиент
1 -- работник клиники
2 -- я гость, мне просто врачей посмотреть
Выберите роль: `

	fmt.Printf("%s", startPosition)
}

func PrintClientMenu() {
	startMenu :=
		`Меню: 
0 -- выйти
1 -- войти
2 -- зарегестрироваться
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintClientLoop() {
	startMenu := `Меню клиента: 
0 -- выйти
1 -- посмотреть свои записи на прием
2 -- посмотреть список своих питомцев
3 -- посмотреть список врачей и их расписание
4 -- добавить питомца
5 -- удалить питомца
6 -- записаться на прием
7 -- узнать информацию о себе 
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintDoctorMenu() {
	startMenu :=
		`Меню: 
0 -- выйти
1 -- войти
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintDoctorLoop() {
	startMenu :=
		`Меню доктора: 
0 -- выйти
1 -- посмотреть свои записи
2 -- изменить расписание
3 -- узнать информацию о себе 
Выберите действие: `

	fmt.Printf("%s", startMenu)
}

func PrintRecords(records models.Records) {
	fmt.Printf("\n Ваши записи:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"№", "Id записи", "Питомец", "Клиент", "Доктор", "Год", "Месяц", "День", "Начало", "Конец")

	for i, r := range records.Records {
		fmt.Fprintf(w, "\n %d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n",
			i+1, r.RecordId, r.PetId, r.ClientId, r.DoctorId, r.DatetimeStart.Year(), r.DatetimeStart.Month(),
			r.DatetimeStart.Day(), r.DatetimeStart.Hour(),
			r.DatetimeEnd.Hour())
	}
	w.Flush()

	fmt.Printf("\nКонец записей!\n\n")
}

func PrintPets(pets models.Pets) {
	fmt.Printf("\nВаши питомцы:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\n",
		"№", "Id питомца", "Кличка", "Тип", "Возраст", "Уровень здоровья")

	for i, p := range pets.Pets {
		fmt.Fprintf(w, "\n %d\t%d\t%s\t%s\t%d\t%d\n",
			i+1, p.PetId, p.Name, p.Type, p.Age, p.Health)
	}
	w.Flush()

	fmt.Printf("\nКонец!\n\n")
}

func PrintDoctors(doctors models.Doctors) {
	fmt.Printf("\nДоктора клиники:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\n",
		"№", "Id доктора", "Логин доктора", "Начало приема", "Конец приема")

	for i, d := range doctors.Doctors {
		fmt.Fprintf(w, "\n %d\t%d\t%s\t%d\t%d\n",
			i+1, d.DoctorId, d.Login, d.StartTime, d.EndTime)
	}
	w.Flush()

	fmt.Printf("\nКонец!\n\n")
}

func PrintClientInfo(client models.Client) {
	fmt.Printf("\nВаш логин: %s\nВаш Id: %d\n\n", client.Login, client.ClientId)
}

func PrintDoctorInfo(doctor models.Doctor) {
	fmt.Printf("\nВаш логин: %s\nВаш Id: %d\nЕжедневное время начала приема: %d\nЕжедневное время конца приема: %d\n\n",
		doctor.Login, doctor.DoctorId, doctor.StartTime, doctor.EndTime)
}

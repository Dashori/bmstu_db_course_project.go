package techUI

import (
	registry "backend/cmd/registry"
	"backend/internal/models"
	"backend/internal/pkg/errors/repoErrors"
	"fmt"
	"os"
	"text/tabwriter"
)

func printRecords(records []models.Record) {
	fmt.Printf("\n Ваши записи:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t\t%s\t%s\n",
		"№", "Id записи", "ID питомец", "Кличка", "ID клиент", "Логин", "Год", "Месяц", "День", "Начало", "Конец")

	for i, r := range records {
		fmt.Fprintf(w, "\n %d\t%d\t%d\t%s\t%d\t%s\t%d\t%d\t%d\t%d\t%d\n",
			i+1, r.RecordId, r.PetId, r.PetName, r.ClientId, r.ClientLogin, r.DatetimeStart.Year(), r.DatetimeStart.Month(),
			r.DatetimeStart.Day(), r.DatetimeStart.Hour(),
			r.DatetimeEnd.Hour())
	}
	w.Flush()

	fmt.Printf("\nКонец записей!\n\n")
}

func printPets(pets []models.Pet) {
	fmt.Printf("\nВаши питомцы:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\n",
		"№", "Id питомца", "Кличка", "Тип", "Возраст", "Уровень здоровья")

	for i, p := range pets {
		fmt.Fprintf(w, "\n %d\t%d\t%s\t%s\t%d\t%d\n",
			i+1, p.PetId, p.Name, p.Type, p.Age, p.Health)
	}
	w.Flush()

	fmt.Printf("\nКонец!\n\n")
}

func getDoctors(a *registry.AppServiceFields) ([]models.Doctor, error) {
	doctors, err := a.DoctorService.GetAllDoctors()

	if err == repoErrors.EntityDoesNotExists {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return doctors, nil
}

func printDoctors(doctors []models.Doctor) {
	fmt.Printf("\nДоктора клиники:\n")
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 9, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\n",
		"№", "Id доктора", "Логин доктора", "Начало приема", "Конец приема")

	for i, d := range doctors {
		fmt.Fprintf(w, "\n %d\t%d\t%s\t%d\t%d\n",
			i+1, d.DoctorId, d.Login, d.StartTime, d.EndTime)
	}
	w.Flush()

	fmt.Printf("\nКонец!\n\n")
}

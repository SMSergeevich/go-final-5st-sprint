package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает поведение парсера данных об активности.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info перебирает строки данных, парсит их и выводит информацию.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// парсим строку
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}

		// формируем строку с информацией
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		// выводим результат в stdout
		fmt.Println(info)
	}
}

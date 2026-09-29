package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит данные о дневной активности пользователя.
type DaySteps struct {
	Steps                 int           // количество шагов
	Duration              time.Duration // длительность прогулки
	personaldata.Personal               // встроенные персональные данные (имя, вес, рост)
}

// Parse парсит строку формата "678,0h50m" и записывает данные в структуру.
func (ds *DaySteps) Parse(datastring string) (err error) {
	// разбиваем строку на части по запятой
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid data format")
	}

	// преобразуем количество шагов в int
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	// шаги должны быть положительными
	if steps <= 0 {
		return errors.New("steps must be greater than zero")
	}

	// преобразуем строку в time.Duration
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	// длительность должна быть положительной
	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	}

	// сохраняем полученные значения в структуру
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	// проверяем корректность всех полей структуры
	if ds.Steps <= 0 {
		return "", errors.New("steps must be greater than zero")
	}
	if ds.Duration <= 0 {
		return "", errors.New("duration must be greater than zero")
	}
	if ds.Weight <= 0 {
		return "", errors.New("weight must be greater than zero")
	}
	if ds.Height <= 0 {
		return "", errors.New("height must be greater than zero")
	}

	// вычисляем дистанцию
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	// вычисляем потраченные калории при ходьбе
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// формируем итоговую строку с информацией
	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, dist, calories,
	)
	return info, nil
}

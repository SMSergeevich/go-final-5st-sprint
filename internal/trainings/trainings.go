package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training хранит данные о тренировке.
type Training struct {
	Steps                 int           // количество шагов
	TrainingType          string        // тип тренировки (бег или ходьба)
	Duration              time.Duration // длительность тренировки
	personaldata.Personal               // встроенные персональные данные (имя, вес, рост)
}

// Parse парсит строку формата "3456,Ходьба,3h00m" в структуру Training.
func (t *Training) Parse(datastring string) (err error) {
	// разбиваем строку на части по запятой
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
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
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	// длительность должна быть положительной
	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	}

	// сохраняем полученные значения в структуру
	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration
	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	// проверяем корректность всех полей структуры
	if t.Steps <= 0 {
		return "", errors.New("steps must be greater than zero")
	}
	if t.Duration <= 0 {
		return "", errors.New("duration must be greater than zero")
	}
	if t.Weight <= 0 {
		return "", errors.New("weight must be greater than zero")
	}
	if t.Height <= 0 {
		return "", errors.New("height must be greater than zero")
	}

	// вычисляем дистанцию и среднюю скорость
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	// в зависимости от типа тренировки считаем калории
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("unknown training type")
	}
	if err != nil {
		return "", err
	}

	// формируем итоговую строку с информацией
	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), dist, speed, calories,
	)
	return info, nil
}

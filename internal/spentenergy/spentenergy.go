package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories считает калории, потраченные при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверяем корректность входных параметров
	if steps <= 0 {
		return 0, errors.New("steps must be greater than zero")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than zero")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than zero")
	}

	// средняя скорость и продолжительность в минутах
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	// формула: вес * скорость * минуты / минут_в_часе, затем коэффициент ходьбы
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories считает калории, потраченные при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверяем корректность входных параметров
	if steps <= 0 {
		return 0, errors.New("steps must be greater than zero")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than zero")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than zero")
	}

	// средняя скорость и продолжительность в минутах
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	// формула: вес * скорость * минуты / минут_в_часе
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

// MeanSpeed вычисляет среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// при нулевой или отрицательной продолжительности, а также при нулевых шагах — скорость нулевая
	if duration <= 0 || steps <= 0 || height <= 0 {
		return 0
	}
	// дистанция / время в часах
	dist := Distance(steps, height)
	return dist / duration.Hours()
}

// Distance вычисляет пройденную дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	// если шагов нет или рост некорректен — дистанция нулевая
	if steps <= 0 || height <= 0 {
		return 0
	}
	// длина шага = рост * коэффициент
	stepLength := height * stepLengthCoefficient
	// дистанция в км = шаги * длина шага / метров в км
	return (float64(steps) * stepLength) / mInKm
}

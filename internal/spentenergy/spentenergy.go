package spentenergy

import (
	"time"

	errors "github.com/DanilaNova/go-5-sprint-final/pkg"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories расчитывает потраченные калории при ходьбе
//
// Вес передаётся в килограммах
// Высота передаётся в метрах
//
// # Типы ошибок
//
// ошибки [RunningSpentCalories]
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	return calories * walkingCaloriesCoefficient, err
}

// RunningSpentCalories расчитывает потраченные калории при беге
//
// Вес передаётся в килограммах
// Высота передаётся в метрах
//
// # Типы ошибок
//
// [errors.ErrStepsIsZeroOrLess] |
// [errors.ErrWeightIsZeroOrLess] |
// [errors.ErrHeightIsZeroOrLess] |
// [errors.ErrDurationIsZeroOrLess]
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.ErrStepsIsZeroOrLess
	}
	if weight <= 0 {
		return 0, errors.ErrWeightIsZeroOrLess
	}
	if height <= 0 {
		return 0, errors.ErrHeightIsZeroOrLess
	}
	if duration <= 0 {
		return 0, errors.ErrDurationIsZeroOrLess
	}

	speed := MeanSpeed(steps, height, duration)

	return weight * speed * duration.Hours(), nil
}

// meanSpeed расчитывает среднюю арифметическую скорость в км/ч
//
// # Высота передаётся в метрах
//
// Если продолжительность равна нулю, возвращается 0
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// Distance расчитывает пройденную дистанцию
//
// Высота передаётся в метрах
func Distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / mInKm
}

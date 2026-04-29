package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DanilaNova/go-5-sprint-final/internal/personaldata"
	"github.com/DanilaNova/go-5-sprint-final/internal/spentenergy"

	errors "github.com/DanilaNova/go-5-sprint-final/pkg"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

// Parse достаёт из строки количество шагов, тип и время тренировки.
//
// Строка должна передаваться в формате "123,Тип,3h4m".
//
// # Типы ошибок:
//
// [errors.ErrIncorrectFormat] |
// [errors.ErrStepsIsZeroOrLess] |
// [errors.ErrDurationIsZeroOrLess] |
// [strconv.NumError] |
// [time.parseDurationError]
func (t *Training) Parse(datastring string) (err error) {
	split := strings.Split(datastring, ",")
	if len(split) != 3 {
		return fmt.Errorf("%w, ожидалось: 123,Тип,3h4m", errors.ErrIncorrectFormat)
	}

	t.Steps, err = strconv.Atoi(split[0])
	if err != nil {
		return err
	}
	if t.Steps <= 0 {
		return errors.ErrStepsIsZeroOrLess
	}

	t.TrainingType = split[1]

	t.Duration, err = time.ParseDuration(split[2])
	if t.Duration <= 0 {
		return errors.ErrDurationIsZeroOrLess
	}

	return err
}

// ActionInfo выводит информацию о тренировке
//
// # Типы ошибок:
//
// [errors.ErrUnknownTraining] |
// ошибки [errors.WalkingSpentCalories] |
// ошибки [errors.RunningSpentCalories]
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var (
		calories float64
		err      error
	)
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.ErrUnknownTraining
	}

	return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}

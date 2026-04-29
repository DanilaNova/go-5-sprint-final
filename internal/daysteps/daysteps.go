package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DanilaNova/go-5-sprint-final/internal/personaldata"
	"github.com/DanilaNova/go-5-sprint-final/internal/spentenergy"

	errors "github.com/DanilaNova/go-5-sprint-final/pkg"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse достаёт из строки количество шагов и время прогулки.
//
// Строка должна передаваться в формате "123,3h4m".
//
// # Типы ошибок:
//
// [errors.ErrIncorrectFormat] |
// [errors.ErrStepsIsZeroOrLess] |
// [errors.ErrDurationIsZeroOrLess] |
// [strconv.NumError] |
// [time.parseDurationError]
func (ds *DaySteps) Parse(datastring string) (err error) {
	split := strings.Split(datastring, ",")
	if len(split) != 2 {
		return fmt.Errorf("%w, ожидалось 123,3h4m", errors.ErrIncorrectFormat)
	}

	ds.Steps, err = strconv.Atoi(split[0])
	if err != nil {
		return err
	}
	if ds.Steps <= 0 {
		return errors.ErrStepsIsZeroOrLess
	}

	ds.Duration, err = time.ParseDuration(split[1])
	if ds.Duration <= 0 {
		return errors.ErrDurationIsZeroOrLess
	}

	return err
}

// ActionInfo выводит информацию о тренировке
//
// # Типы ошибок:
//
// ошибки [errors.WalkingSpentCalories]
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, ds.Steps, distance, calories), nil
}

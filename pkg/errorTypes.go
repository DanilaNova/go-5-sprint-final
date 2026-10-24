package errorTypes

import (
	"errors"
	"fmt"
)

var (
	ErrIncorrectFormat = errors.New("wrong string format")
	ErrUnknownTraining = errors.New("unknown training type")

	ErrValueIsZeroOrLess    = errors.New("the value is less than or equal to zero")
	ErrStepsIsZeroOrLess    = fmt.Errorf("%w: steps", ErrValueIsZeroOrLess)
	ErrWeightIsZeroOrLess   = fmt.Errorf("%w: weight", ErrValueIsZeroOrLess)
	ErrHeightIsZeroOrLess   = fmt.Errorf("%w: height", ErrValueIsZeroOrLess)
	ErrDurationIsZeroOrLess = fmt.Errorf("%w: duration", ErrValueIsZeroOrLess)
)

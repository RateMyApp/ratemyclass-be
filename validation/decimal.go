package validation

import (
	"errors"
	"fmt"
	"math"

	validation "github.com/go-ozzo/ozzo-validation"
)

func MaxDecimalPlaces(max int) validation.RuleFunc {
	return func(value interface{}) error {
		switch val := value.(type) {
		case float32:
			valuef := val * float32(math.Pow(10.0, float64(max)))
			extra := valuef - float32(int(valuef))
			if extra > 0 {
				return fmt.Errorf(fmt.Sprintf("must be at most %v decimal places but got %v", max, val))
			}

		default:
			return errors.New("value must be a float")
		}
		return nil
	}
}

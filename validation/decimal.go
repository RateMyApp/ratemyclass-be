package validation

import (
	"errors"
	"fmt"
	"math"

	validation "github.com/go-ozzo/ozzo-validation"
)

func MaxDecimalPlaces(max uint) validation.RuleFunc {
	return func(value interface{}) error {
		switch val := value.(type) {
		case float32:
			v := float64(val)
			return MaxDecimalPlaces(max)(v)
		case float64:
			valuef := val * float64(math.Pow(10.0, float64(max)))
			println(valuef)
			extra := valuef - float64(int(valuef))
			if !(extra <= 0) {
				return errors.New(fmt.Sprintf("Must be at most %v decimal places", max))
			}

		default:
			return errors.New("Value must be a float")
		}
		return nil
	}
}

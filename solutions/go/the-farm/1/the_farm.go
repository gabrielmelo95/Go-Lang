package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, cowsQty int) (foodPerCow float64, err error) {
	fodderAmount, err := fc.FodderAmount(cowsQty)
	if err != nil {
		return 0.0, err
	}
	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	foodPerCow = fodderAmount * fatteningFactor / float64(cowsQty)
	err = nil
	return
}

func ValidateInputAndDivideFood(fc FodderCalculator, cowsQty int) (foodPerCow float64, err error) {
	if cowsQty > 0 {
		return DivideFood(fc, cowsQty)
	} else {
		return 0.0, errors.New("invalid number of cows")
	}
}

type InvalidCowsError struct {
	cowsQty int
	details string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %v", e.cowsQty, e.details)
}
func ValidateNumberOfCows(cowsQty int) error {
	if cowsQty < 0 {
		return &InvalidCowsError{
			cowsQty: cowsQty,
			details: "there are no negative cows",
		}
	} else if cowsQty == 0 {
		return &InvalidCowsError{
			cowsQty: 0,
			details: "no cows don't need food",
		}
	}
	return nil
}

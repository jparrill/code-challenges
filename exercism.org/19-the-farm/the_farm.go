package thefarm

import (
    "errors"
    "fmt"
)

type SillyNephewError struct { message string }
func (e *SillyNephewError) Error() string { return e.message }

var NegativeFodder = errors.New("negative fodder")
var NonScaleError = errors.New("non-scale error")
var DivisionByZeroError = errors.New("division by zero")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	sne := &SillyNephewError{
        message: fmt.Sprintf("silly nephew, there cannot be %d cows", cows),
    }
    
    if cows < 0 {
        return 0.0, sne
    } else if cows == 0 {
    	return 0.0, DivisionByZeroError
    }

    fodder, err := weightFodder.FodderAmount()
    if (err != ErrScaleMalfunction) && (err != nil) {
        return 0, NonScaleError
    } else if (fodder < 0) { 
    	return 0, NegativeFodder
    } else {
    	if (err == ErrScaleMalfunction) {
        	fodder *= 2
        }
    }
    
	return (fodder / float64(cows)), nil
}

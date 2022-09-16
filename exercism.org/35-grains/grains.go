package grains

import (
    "errors"
    "math"
)

func Square(number int) (uint64, error) {
	var result uint64
    var err error
    
    switch {
    case number < 0:
    	err = errors.New("Number could not be negative")
        result = 0
    case number == 0:
    	err = errors.New("Number could not be 0")
        result = 0
    case number > 64:
    	err = errors.New("Number could not be greater than 64")
        result = 0
    case number > 0 && number < 65:
    	err = nil
        result = uint64(math.Pow(float64(2),float64(number)-1))
    }

    return result, err
}

func Total() uint64 {
    var total, current uint64
    var err error
    
	for pos := 1; pos <= 64; pos++ {
        // Nexts iterations
        current, err = Square(pos)
        if err != nil {
            panic(err)
        }
        total += current
    }

    return total
}

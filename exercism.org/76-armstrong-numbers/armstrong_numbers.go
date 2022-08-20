package armstrong

import "math"

func IsNumber(n int) bool {
    var count, num int
    var result float64
    num = n
	for num > 0{
    	count++
        num /= 10
    }
	
    num = n

    if n < 10 {
        return true
    }

    for i:=0; i<=count;i++{
        value := num%10
    	result += math.Pow(float64(value), float64(count))
        num /= 10
    }

    if result == float64(n) {
        return true
    }

    return false
    
}

package diffsquares

import "math"

func SquareOfSum(n int) int {
	var result float64

    for a:=1;a<=n;a++ { result = result + float64(a) }

    return int(math.Pow(result,2))
}

func SumOfSquares(n int) int {
	var result float64

    for a:=1;a<=n;a++ { result = result + math.Pow(float64(a),2) }
    
    return int(result)
}

func Difference(n int) int {
	var r_sqos, r_sosq float64

    for a:=1;a<=n;a++ { r_sqos = r_sqos + float64(a) }
    r_sqos = math.Pow(r_sqos,2)

    for a:=1;a<=n;a++ { r_sosq = r_sosq + math.Pow(float64(a),2) }

    return int(math.Max(r_sosq, r_sqos)) - int(math.Min(r_sosq, r_sqos))   
}

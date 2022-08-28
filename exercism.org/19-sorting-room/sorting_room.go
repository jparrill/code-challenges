package sorting
import (
    "fmt"
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    v, _ := strconv.Atoi(fnb.Value())
    
	switch t := fnb.(type) {
    case FancyNumber:
    	return v
    default:
    	fmt.Println("Number is not a FancyNumber but : ", t)
        fmt.Println("Value: ", v)
    	return 0
    }
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    v, _ := strconv.Atoi(fnb.Value())

    switch t := fnb.(type) {
    case FancyNumber:
    	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(v))
    default:
    	fmt.Println(t)
    	return fmt.Sprintf("This is a fancy box containing the number 0.0")
    }
    
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	fmt.Println("Entry Value: ",i)

    switch t := i.(type) {
    case int:
        return DescribeNumber(float64(i.(int)))
    case float64:
    	return DescribeNumber(i.(float64))
    case NumberBox:
    	return DescribeNumberBox(i.(NumberBox))
    case FancyNumberBox:
    	return DescribeFancyNumberBox(i.(FancyNumberBox))
    default:
    	fmt.Println("Value: ",t)
    	return fmt.Sprintf("Return to sender")
    }
}

package luhn

import (
    "fmt"
    "strings"
    )

func Valid(id string) bool {
    var sum, dob, conv, count int

    // Strip spaces
    word := strings.Replace(id," ","",-1)

    for i := len(word)-1 ; i >= 0; i-- {
        val := rune(word[i])
        if (val <= 9) || (val >= 48 && val <= 57) {
            conv = int(val)
            if val >= 48 && val <= 57 {
                conv = conv - 48 
            }

            if count%2 != 0 {
                fmt.Println("Odd in reverse: ", count)
                fmt.Println("	digit: ", conv)
        		fmt.Println("	dob: ", conv+conv)
                if dob = conv + conv; dob > 9 {
                    fmt.Println("	Doubling digit: ", dob)
               		dob = dob - 9
            	}
            	sum += dob
            } else {
            	fmt.Println("Even, Not Doubling: ", count, conv)
            	sum += conv
        	}
        count++
        } else {
        	return false
        }
    }// for

    if count <= 1 {
        return false
    }
	
	fmt.Println("Total: ", sum)
    if sum % 10 != 0  && sum > 9{
        return false
    }

    return true
}

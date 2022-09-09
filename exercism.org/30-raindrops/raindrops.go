package raindrops

import (
    "strconv"
    "strings"
    )

func Convert(number int) string {
	factors := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	result := make([]string,0,0)

    for k, v := range factors {
        if number % k == 0 {
            result = append(result, v)
        }
    }

    if len(result) == 0 {
        return strconv.Itoa(number)
    }

    return strings.Join(result, "")
}

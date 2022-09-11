package isogram

import "strings"

func IsIsogram(word string) bool {
	result := true
    mr := make(map[rune]bool)
    w := strings.ToLower(word)
    

    for _,k := range w {
        if (rune(k) != 45) && (rune(k) != 32) {
        	if _, ok := mr[k]; ok {
                result = false
                break
            } else {
            	mr[k] = true
            }
        }
    }

    return result
}

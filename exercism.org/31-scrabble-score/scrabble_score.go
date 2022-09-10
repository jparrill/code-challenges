package scrabble

import (
    "fmt"
    "strings"
)

func Score(word string) int {
	var score int

    final_word := strings.ToUpper(word)
    
    for _,v := range final_word {
        switch v {
        case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
        	score+=1
        case 'D', 'G':
        	score+=2
        case 'B', 'C', 'M', 'P':
        	score+=3
        case 'F', 'H', 'V', 'W', 'Y':
        	score+=4
        case 'K':
        	score+=5
        case 'J', 'X':
        	score+=8
        case 'Q','Z':
        	score+=10
        default:
        	fmt.Println("Rune not found")
        }
    }
    
    return score
}

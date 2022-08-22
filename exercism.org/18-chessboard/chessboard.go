package chessboard
import (
    "fmt"
)

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    var count int
	if _, exists := cb[file]; !exists {
        return 0
    }

    for _,v := range cb[file] {
        if v {
        	count++
        }
    }

    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var count int
    var char rune
    
	if (rank <= 0) || (rank >= 9) {
        return 0
    }

    // Iterate over ASCII values of A to H
    for f := 65; f <= 72; f++ {
        char = rune(f)
        fmt.Printf("File: %v Rank: %d Value: %v\n", string(char), rank, cb[string(char)][rank - 1])
        if (cb[string(char)][rank - 1] == true) {
            count++
        }
    }

    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int
    var char rune
    
	// Iterate over ASCII values of A to H
    for f := 65; f <= 72; f++ {
        char = rune(f)
        count += len(cb[string(char)])
    }

    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count int
    var char rune
    
	// Iterate over ASCII values of A to H
    for f := 65; f <= 72; f++ {
        for r := 0; r <= 7; r++{
        	char = rune(f)
        	fmt.Printf("File: %v Rank: %d Value: %v\n", string(char), r, cb[string(char)][r])
        	if (cb[string(char)][r] == true) {
            	count++
        	}
        }
    }

    return count
}

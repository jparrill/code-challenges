package logs

import (
    "fmt"
    "strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var app string

    fmt.Println(log)
    for _, char := range log {
        fmt.Println(string(char))
        switch {
        case string(char) == "❗":
        	app = "recommendation"
            fmt.Println("Found: ",app)
            return app
        case string(char) == "🔍":
        	app = "search"
            fmt.Println("Found: ",app)
            return app
        case string(char) == "☀":
        	app = "weather"
            fmt.Println("Found: ",app)
            return app
        default:
        	app = "default"
        }
    }
    
    return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    var xs, newXs []string
    xs = strings.Split(log, " ")

    fmt.Println(log)
    for _, char := range xs {
        if char == string(oldRune) {
            newXs = append(newXs, string(newRune))
        } else {
        	newXs = append(newXs, string(char))
        }
    }
	fmt.Println(strings.Join(newXs, " "))
    
    return strings.Join(newXs, " ")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    var lenght int
    for _,_ = range log {
        lenght += 1
        if lenght > limit {
        	return false
    	}
    }

    return true
}

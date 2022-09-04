package parsinglogfiles
import (
    "regexp"
    "fmt"
    )

func IsValidLine(text string) bool {
    var valid bool
	re := regexp.MustCompile(`^\[[A-Z]?[A-Z]?[A-Z]?\]`)

    if !(re.MatchString(text)) {
    	return false
    }
    
    switch {
    case re.FindString(text) == "[TRC]":
    	valid = true
    case re.FindString(text) == "[DBG]":
    	valid = true
    case re.FindString(text) == "[INF]":
    	valid = true
    case re.FindString(text) == "[WRN]":
    	valid = true
    case re.FindString(text) == "[ERR]":
    	valid = true
    case re.FindString(text) == "[FTL]":
    	valid = true
    }

    return valid
}

func SplitLogLine(text string) []string {
    var sections []string
	re := regexp.MustCompile(`\<\W*\>`)

    sections = re.Split(text, -1)

    return sections
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re := regexp.MustCompile(`\"(\w*? )*(p|P)(a|A)(s|S)(s|S)(w|W)(o|O)(r|R)(d|D)(\w*? )*\"`)
    
    for _,k := range lines{
        fmt.Println("1: ",k)
        if re.MatchString(k) {
            fmt.Println("2: ",k)
            count++   
        }
    }

    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+\s*?`)
    txt := re.ReplaceAllString(text,"")
    return txt
}

func TagWithUserName(lines []string) []string {
    var xs []string
	re := regexp.MustCompile(`User\s+(\w+)`)
    for _,k := range lines {
    	if user := re.FindStringSubmatch(k); user != nil {
        	xs = append(xs, fmt.Sprintf("[USR] %s %s", user[1], k))    
        } else {
        	xs = append(xs,k)
        }
        
    }
    return xs
}

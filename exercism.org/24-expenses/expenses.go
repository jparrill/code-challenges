package expenses
import "fmt"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
    var filteredRecs []Record
    
	for _,rec := range in {
        if predicate(rec) {
            filteredRecs = append(filteredRecs, rec)
        }
    }

    return filteredRecs
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func (rec Record) bool {
        if (rec.Day >= p.From) && (rec.Day <= p.To) {
            return true
        }

        return false
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func (rec Record) bool {
        if rec.Category == c {
            return true
        }
    	return false
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    var amount float64
	for _, rec := range in {
        if (rec.Day >= p.From) && (rec.Day <= p.To) {
            amount += rec.Amount
        }
    }

    return amount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
    var amount float64
    var counterCat int
    
    for _, rec := range in {
        if (rec.Category == c) {
            if (rec.Day >= p.From) && (rec.Day <= p.To) {
            	amount += rec.Amount
            }
        	counterCat++
        }
    }

    if counterCat <= 0 {
        return 0, fmt.Errorf("unknown category %s", c)
    }

    return amount, nil
}

package booking

import (
    "time"
    "fmt"
    )

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/02/2006 15:04:05"
	t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	layout := "January 2, 2006 15:04:05"
    appointment, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }

    return now.After(appointment)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    appointment, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    if (appointment.Hour() >= 12) && (appointment.Hour() < 18) {
        return true
    }

    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    a, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    
	return fmt.Sprintf("You have an appointment on %v, %v %d, %d, at %d:%d.",a.Weekday(), a.Month(), a.Day(), a.Year(), a.Hour(), a.Minute() )
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

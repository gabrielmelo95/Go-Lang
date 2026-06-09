package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	formatedTime, _ := time.Parse("1/2/2006 15:04:05", date)
	return formatedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	start := time.Now()
	end, _ := time.Parse("January 2, 2006 15:04:05", date)
	difference := end.Sub(start)
	if difference < 0 {
		return true
	}
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	timeBasedTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	afternoonStart, _ := time.Parse("15:04", "12:00")
	afternoonEnd, _ := time.Parse("15:04", "18:00")
	if timeBasedTime.Hour() >= afternoonStart.Hour() && timeBasedTime.Hour() < afternoonEnd.Hour() {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointmentTime, _ := time.Parse("1/2/2006 15:04:05", date)

	return "You have an appointment on " + appointmentTime.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(2026, time.September, 15, 0, 0, 0, 0, time.UTC)
}

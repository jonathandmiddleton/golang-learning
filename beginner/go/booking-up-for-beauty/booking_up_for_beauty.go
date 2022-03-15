package booking

import "time"

const longForm = "7/13/2020 15:04:05 -0700 MST"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse(longForm, "2021/11/03 13:45:00")
	return t

}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	t, _:= time.Parse(longForm, "7/13/2020 15:04:05 -0700 MST")
		return time.Now().After(t)
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}

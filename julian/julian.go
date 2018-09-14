package julian

import (
	"github.com/ryanfaerman/calendar"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

//go:generate stringer -type=Month
type Month int

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var StandardMonthLengths = [12]int{
	31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

// IsLeap determines if the year of the given date is a leap year
func (d Date) IsLeap() bool {
	return d.Year%4 == 0
}

// daysInMonth returns the number of days for the given month and year
func daysInMonth(d Date) int {
	if d.IsLeap() && Month(d.Month) == February {
		return 29
	}
	return StandardMonthLengths[d.Month-1]
}

// dayOfYear calculates the elapsed number of days for the given year
func dayOfYear(d Date) int {
	dOY := d.Day + 31*(d.Month-1)
	if Month(d.Month) > February {
		dOY -= (4*d.Month + 23) / 10
		if d.IsLeap() {
			dOY++
		}
	}

	return dOY
}

// ToAbsolute implements the Calendar interface. It returns the number of days
// since the Epoch.
func (d Date) ToAbsolute() int {
	days := dayOfYear(d)       // days so far this year
	days += 365 * (d.Year - 1) // days in prior years
	days += (d.Year - 1) / 4   // Julian leap years
	days -= 2                  // days elapsed before absolute date 1

	return days
}

// FromAbsolute implements the Calendar interface, wrapping the constructor
func (d Date) FromAbsolute(days int) calendar.Calendar {
	return NewFromAbsolute(days)
}

// NewFromAbsolute converts the absolute days since Epoch into its equivalent
// Julian Date.
func NewFromAbsolute(days int) Date {
	d0 := days + 2  // adjust for epoch
	n4 := d0 / 1461 // number of Julian leap years
	d1 := d0 % 1461 // number of Julian leap days
	n1 := d1 / 365

	day := (d1 % 365)
	year := 4*n4 + n1 + 1

	month := 1
	mlen := daysInMonth(Date{Month: month, Year: year})
	for mlen < day {
		day -= mlen
		month++
		mlen = daysInMonth(Date{Month: month, Year: year})
	}

	return Date{year, month, day}
}

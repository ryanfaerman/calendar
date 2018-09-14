package gregorian

import "testing"

func TestIsLeap(t *testing.T) {
	examples := map[int]bool{
		2004: true,
		2020: true,
		2000: true,
		2018: false,
		2017: false,
		1982: false,
	}

	for year, expected := range examples {
		if (Date{Year: year}).IsLeap() != expected {
			t.Errorf("Incorrect leap year result for %d", year)
		}
	}
}

func TestDaysInMonth(t *testing.T) {
	examples := []struct {
		Year  int
		Month Month
		Days  int
	}{
		{2018, February, 28},
		{2018, September, 30},
		{2004, February, 29},
	}

	for _, example := range examples {
		result := daysInMonth(Date{Month: int(example.Month), Year: example.Year})
		if result != example.Days {
			t.Errorf(
				"Incorrect number of days (%d) for %s in %d; expected %d",
				result,
				example.Month.String(),
				example.Year,
				example.Days,
			)
		}
	}
}

func TestDayOfYear(t *testing.T) {
	examples := map[Date]int{
		Date{1987, 1, 1}:   1,
		Date{1980, 12, 31}: 366,
	}

	for date, days := range examples {
		if dayOfYear(date) != days {
			t.Errorf("Incorrect number of days (%d) for date %v", days, date)
		}
	}
}

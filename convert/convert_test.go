package convert

import (
	"testing"

	"github.com/ryanfaerman/calendar"
	"github.com/ryanfaerman/calendar/gregorian"
	"github.com/ryanfaerman/calendar/julian"
)

func TestConversion(t *testing.T) {
	examples := map[calendar.Calendar]calendar.Calendar{
		gregorian.Date{2018, 9, 15}: julian.Date{2018, 9, 2},
	}

	for from, to := range examples {
		forward := Convert(from, to)
		if forward != to {
			t.Errorf(
				"Expected %v and %v to be equivalent; got %v",
				forward, to,
				forward,
			)
		}
		reverse := Convert(to, from)
		if reverse != from {
			t.Errorf(
				"Expected %v and %v to be equivalent; got %v",
				reverse, to,
				reverse,
			)
		}
	}
}

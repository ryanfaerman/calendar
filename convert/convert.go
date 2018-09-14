package convert

import "github.com/ryanfaerman/calendar"

func Convert(from, to calendar.Calendar) calendar.Calendar {
	return to.FromAbsolute(from.ToAbsolute())
}

package calendar

type Calendar interface {
	IsLeap() bool
	ToAbsolute() int
	FromAbsolute(int) Calendar
}

func Convert(from, to Calendar) Calendar {
	return to.FromAbsolute(from.ToAbsolute())
}

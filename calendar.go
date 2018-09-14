package calendar

type Calendar interface {
	IsLeap() bool
	ToAbsolute() int
	FromAbsolute(int) Calendar
}

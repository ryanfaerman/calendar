package main

import (
	"fmt"

	"github.com/ryanfaerman/calendar/gregorian"
)

type Calendar interface {
	IsLeap() bool
	ToAbsolute() int
	FromAbsolute(int) Calendar
}

func Convert(from, to Calendar) Calendar {
	return to.FromAbsolute(from.ToAbsolute())
}

func main() {

	g := gregorian.Date{1945, 11, 12}
	abs := g.ToAbsolute()
	fmt.Println(abs)
	fmt.Println(gregorian.NewFromAbsolute(abs))

	fmt.Println(Convert(g, gregorian.Date{}))

}

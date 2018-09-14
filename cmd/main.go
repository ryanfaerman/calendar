package main

import (
	"fmt"

	"github.com/ryanfaerman/calendar"
	"github.com/ryanfaerman/calendar/gregorian"
)

func main() {

	g := gregorian.Date{1945, 11, 12}
	abs := g.ToAbsolute()
	fmt.Println(abs)
	fmt.Println(gregorian.NewFromAbsolute(abs))

	fmt.Println(calendar.Convert(g, gregorian.Date{}))

}

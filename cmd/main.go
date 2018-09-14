package main

import (
	"fmt"

	"github.com/ryanfaerman/calendar"
	"github.com/ryanfaerman/calendar/gregorian"
	"github.com/ryanfaerman/calendar/julian"
)

func main() {

	g := gregorian.Date{2018, 9, 15}
	j := julian.Date{2018, 9, 2}

	fmt.Println("gregorian -> julian", calendar.Convert(g, julian.Date{}))
	fmt.Println("julian -> gregorian", calendar.Convert(j, gregorian.Date{}))

}

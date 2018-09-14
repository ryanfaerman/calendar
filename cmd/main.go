package main

import (
	"fmt"

	"github.com/ryanfaerman/calendar/convert"
	"github.com/ryanfaerman/calendar/gregorian"
	"github.com/ryanfaerman/calendar/julian"
)

func main() {

	g := gregorian.Date{2018, 9, 15}
	j := julian.Date{2018, 9, 2}

	fmt.Println("gregorian -> julian", convert.Convert(g, julian.Date{}))
	fmt.Println("julian -> gregorian", convert.Convert(j, gregorian.Date{}))

}

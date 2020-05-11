package main

import (
	"flag"
	"fmt"

	"github.com/mgirouard/ww"
)

func main() {
	a := flag.Bool("a", false, "Include Saturdays")
	u := flag.Bool("u", false, "Include Sundays")
	r := flag.Bool("r", false, "Format as a date range")
	flag.Parse()

	w := ww.NewWeek()
	w.Saturday = *a
	w.Sunday = *u
	w.Range = *r
	fmt.Println(w.String())
}

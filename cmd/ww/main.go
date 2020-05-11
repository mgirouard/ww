package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mgirouard/ww"
)

func main() {
	a := flag.Bool("a", false, "Include Saturday")
	l := flag.Bool("l", false, "Last week")
	n := flag.Bool("n", false, "Next week")
	r := flag.Bool("r", false, "Format as a date range")
	u := flag.Bool("u", false, "Include Sunday")
	flag.Parse()

	start := time.Now()
	if *l {
		start = start.AddDate(0, 0, -7)
	} else if *n {
		start = start.AddDate(0, 0, 7)
	}

	w := ww.NewWeek(start)
	w.Saturday = *a
	w.Sunday = *u
	w.Range = *r

	args := flag.Args()
	if len(args) > 0 {
		w.Format = args[0]
	}
	fmt.Println(w.String())
}

package main

import (
	"fmt"
	"os"
	"time"
)

const DefaultFormat = "2006-01-02"

type Command struct {
	start  time.Time
	format string
}

func NewCommand(f string) Command {
	return Command{time.Now(), f}
}

func (c Command) String() string {
	out := ""
	dow := int(c.start.Weekday())
	if dow != 0 {
		c.start = c.start.AddDate(0, 0, -dow)
	}
	for i := 1; i <= 5; i++ {
		c.start = c.start.AddDate(0, 0, 1)
		out = out + fmt.Sprintf("%s\n", c.start.Format(c.format))
	}
	return out
}

func main() {
	var f string
	if len(os.Args) > 2 {
		fmt.Println("Expected either zero or one arguments")
		os.Exit(1)
	} else if len(os.Args) == 1 {
		f = DefaultFormat
	} else {
		f = os.Args[1]
	}
	c := NewCommand(f)
	fmt.Println(c.String())
}

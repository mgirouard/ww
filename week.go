package main

import (
	"strings"
	"time"
)

const DefaultFormat = "2006-01-02"
const DefaultSeparator = "\n"
const RangeSeparator = "--"

type Week struct {
	start     time.Time
	Format    string
	Separator string
	Days      []string
	Saturday  bool
	Sunday    bool
	Range     bool
}

func NewWeek(s time.Time) *Week {
	d := int(s.Weekday())
	if d != 0 {
		s = s.AddDate(0, 0, -d)
	}
	return &Week{
		start:     s,
		Format:    DefaultFormat,
		Separator: DefaultSeparator,
		Days:      []string{},
	}
}

func (w *Week) Render() {
	s := w.start
	days := []string{}
	for d := 0; d <= 6; d++ {
		days = append(days, s.Format(w.Format))
		s = s.AddDate(0, 0, 1)
	}
	if !w.Sunday {
		days = days[1:]
	}
	if !w.Saturday {
		days = days[:len(days)-1]
	}
	if w.Range {
		days = []string{days[0], days[len(days)-1]}
		w.Separator = RangeSeparator
	}
	w.Days = days
}

func (w Week) String() string {
	w.Render()
	return strings.Join(w.Days, w.Separator)
}

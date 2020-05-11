# ww

`ww` is a dumb little tool that tells you the dates for the work week.

	$ go get github.com/mgirouard/ww
	$ ww
	2020-04-20
	2020-04-21
	2020-04-22
	2020-04-23
	2020-04-24

It's most useful when you keep a work journal.

	$ vim -o $(ww)

On Mondays and Fridays it's useful to compare the current week with another to
help with planning. I keep three splits in my terminal to recall notes from
last week or plan ahead for the coming week.
	
	# Last week
	$ vim -o $(ww --last)

	# This week
	$ vim -o $(ww)

	# Next week
	$ vim -o $(ww --next)

A format can be specified to simplify working with specific file types or
arbitrary paths. It's format string is strange, but at least it's documented in
[Time.Format].

	$ vim -o $(ww Documents/2006-01-02.md)

I can never remember it, so I keep a dotfile to remember it for me:

	$ cat ~/.ww.yaml
	format: Documents/2006-01-02.md

[Time.Format]: https://golang.org/pkg/time/#Time.Format

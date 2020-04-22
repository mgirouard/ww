# ww

`ww` is a dumb little tool that tells you the dates for the work week.

	$ ww
	2020-04-20
	2020-04-21
	2020-04-22
	2020-04-23
	2020-04-24

This is most useful when you keep a work journal.

	$ vim -o $(ww)

It's possible to do the previous or next week as well. I keep three splits in
my terminal to recall notes from last week or plan ahead for the coming week.
	
	# Last week
	$ vim -o $(ww --last)

	# This week
	$ vim -o $(ww)

	# Next week
	$ vim -o $(ww --next)

I prefer to work in Markdown.

	$ vim -o $(ww --ext=md)

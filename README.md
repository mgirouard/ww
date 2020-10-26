# `cow` is a different kind of `cal`

    $ go get github.com/mgirouard/cow

Default command renders a cal-like calendar:

    $ cow
        October 2020
    Su Mo Tu We Th Fr Sa
                 1  2  3
     4  5  6  7  8  9 10
    11 12 13 14 15 16 17
    18 19 20 21 22 23 24
    25 26 27 28 29 30 31

Cow is more than cal. Cow outputs a list of dates and assumes you want to do
stuff with that list.

    $ cow -f %Y%m%d.md | touch

Run `cow help` for info and examples.

## Hacking

To build all the things:

	$ make

If you use ctags:
	
	$ make tags

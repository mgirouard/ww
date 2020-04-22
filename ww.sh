#!/bin/bash
set -eu

# determine the current day of the week
day=$(date +%w)

# there is only one True date format
fmt=+%Y-%m-%d

# set a prefix string, defaulting to the current week (eg "this fri")
week=this
for arg in "$@"; do
    case "$arg" in
        -n|--next)
        week=next
        fl_next=1
        ;;

        -l|--last)
        week=last
        fl_last=1
        ;;

        *)
        ;;
    esac
done

out() {
    # first param will be the day of the week
    # eg: "mon" or "tue" etc
    local dow=$1

    # determine the day index of $dow
    # 0 = Sunday; 6 = Saturday
    # this is needed below to adjust a prefix string for date's -d param
    local idx=$(date +%w -d $dow)

    # work around quirks with last/this/next week prefixes given the current
    # day of the week
    if [ $idx -lt $day ] && [ $week == "this" ]; then
        local week=last
    elif [ $idx -lt $day ] && [ $week == "last" ]; then
        local week="-2 weeks"
    elif [ $idx -gt $day ] && [ $week == "next" ]; then
        local week="+1 week"
    fi

    echo $(date -d "$week $dow" $fmt)
}

out mon
out tue
out wed
out thu
out fri

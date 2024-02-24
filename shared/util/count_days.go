package util

import (
	"strconv"
	"time"
)

func CountDate(year1, month1, day1, year2, month2, day2 string) float64 {
	// The leap year 2016 had 366 days.
	t1 := Date(year1, month1, day1)
	t2 := Date(year2, month2, day2)
	days := t2.Sub(t1).Hours() / 24
	return days // 366
}

func Date(year, month, day string) time.Time {
	intYear, _ := strconv.Atoi(year)
	intMonth, _ := strconv.Atoi(month)
	intDay, _ := strconv.Atoi(day)
	return time.Date(intYear, time.Month(intMonth), intDay, 0, 0, 0, 0, time.UTC)
}

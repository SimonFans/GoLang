package main

import (
	"fmt"
	"time"
)

func main() {
	// 2022-11-02 13:20:31.595302752 -0700 PDT m=+0.000036103
	t := time.Now()
	// year=2022 month=November day=2
	year, month, day := t.Date()
	fmt.Printf("year=%v month=%v day=%v\n", year, month, day)
	fmt.Println(month == time.November)
	loc, _ := time.LoadLocation("Local")
	// Local
	fmt.Println(loc)
	const shortForm = "2006-Jan-02"
	t1, _ := time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	// 2012-07-09 00:00:00 -0700 PDT
	fmt.Println(t1)
	// Wednesday
	weekday := time.Now().Weekday()
	fmt.Println(weekday)
	fmt.Printf("weekday in number format=%v\n", int(weekday))
}

package main

import "fmt"
import "time"

func main() {
	
	p := fmt.Println
	
	now := time.Now()
	// 2019-08-14 12:46:13.6297598 +0800 CST m=+0.003989001
	p("now:",now)
	
	then := time.Date(2009, 11, 17, 20 ,34, 58, 651387237, time.UTC)
	//2009-11-17 20:34:58.651387237 +0000 UTC
	p("then:",then)
	
	p("year:",then.Year())
	p("month:",then.Month())
	p("day:",then.Day())
	p("hour:",then.Hour())
	p("minute:",then.Minute())
	p("second:",then.Second())
	p("nanosecond:", then.Nanosecond())
	//UTC time zone.
	p("location:",then.Location())
	
	p("weekday:",then.Weekday())
	
	p("before:",then.Before(now))
	p("after:",then.After(now))
	p("equal:",then.Equal(now))
	
	//The Sub methods returns a Duration representing the interval between two times.
	diff := now.Sub(then)
	p("diff:", diff)
	
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	
	p(then.Add(diff))
	p(then.Add(-diff))
}
package time_demo

import (
	"fmt"
	"time"
)

const TimeFormat_MINE = "2006-01-02 15:04:05"
const YYYY_MM_DD = "2006-01-02"

func TimeFormat() {
	var t time.Time = time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(TimeFormat_MINE))
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	printTime(t)
}

func printTime(t time.Time) {
	fmt.Println(t.Format(TimeFormat_MINE))
	year, week := t.ISOWeek()
	fmt.Printf("%4d年 第%02d周 周%d\n", year, week, t.Weekday())
}

func Time2() {
	var t_init time.Time = time.Now()
	//fmt.Println(t_init.Unix())
	//fmt.Println(t_init.UnixNano())
	var t time.Time = time.Unix(t_init.Unix(), 0)
	printTime(t)
	t = time.Unix(0, t_init.UnixNano())
	printTime(t)
}

func Time3() {
	t, err := time.Parse(TimeFormat_MINE, "2021-05-01 10:14:25")
	if err != nil {
		fmt.Printf("An error occurred\n")
		return // exit the function on error
	}
	printTime(t)
}

func GetTime(timeString string) time.Time {
	t, err := time.Parse(YYYY_MM_DD, timeString)
	if err != nil {
		panic("时间格式有误")
	}
	//fmt.Println(t.Format(YYYY_MM_DD))
	return t
}

func GetTimeString(thisTime *time.Time) string {
	return thisTime.Format(YYYY_MM_DD)
}

func Time4() {
	t := time.Now()
	var one_week time.Duration = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec (1s = 1000000000 ns)
	one_week_after := t.Add(one_week)
	printTime(one_week_after)
}

// 休眠
func Time5() {
	printTime(time.Now())
	var second5 time.Duration = 5 * 1e9 // 5s
	time.Sleep(second5)
	printTime(time.Now())
}

func Time6() {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return
	}
	t := time.Date(2021, 5, 5, 10, 15, 0, 0, location)
	printTime(t)
}

package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	TimeFormat1 = "2006-01-02 15:04:05"
)

func ToTimeByEast7(timeStr string) (time.Time, error) {
	local, _ := time.LoadLocation("Asia/Shanghai") // 东八区
	return time.ParseInLocation(TimeFormat1, timeStr, local)
}

func FormatTime(t time.Time) string {
	return t.Format(TimeFormat1)
}

type Time struct {
	time.Time
}

const (
	zeroTime = "0000-00-00 00:00:00"
)

func (t Time) MarshalJSON() ([]byte, error) {
	// 如果是默认值的话，则需要格式化为: 0000-00-00 00:00:00
	tms := zeroTime
	if !t.Time.IsZero() {
		tms = FormatTime(t.Time)
	}
	ret := fmt.Sprintf("\"%s\"", tms)
	return []byte(ret), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	tmd := string(data)
	tmd = strings.Trim(tmd, "\"")
	if tmd == zeroTime {
		t.Time = time.Time{}
		return nil
	}
	tm, err := ToTimeByEast7(tmd)
	if err != nil {
		return err
	}
	t.Time = tm
	return nil
}

func test() {

	tm := time.Time{}
	t := tm.Format("2006-01-02 15:04:05")
	fmt.Printf("null=%d, is zero?=%t, value=%s\n", tm.Unix(), tm.IsZero(), t)

	tm = time.Now()
	t = tm.Format("2006-01-02 15:04:05")
	fmt.Printf("null=%d, is zero?=%t, value=%s\n", tm.Unix(), tm.IsZero(), t)

	type TimeTest struct {
		Created Time  `json:"created"`
		Updated *Time `json:"updated"`
	}

	tt := TimeTest{
		Created: Time{time.Now()},
		Updated: &Time{time.Time{}},
	}

	data, err := json.Marshal(tt)
	fmt.Printf("data=%s, value=%+v, err=%v\n", string(data), tt, err)

	tt2 := TimeTest{}
	err = json.Unmarshal(data, &tt2)
	fmt.Printf("data=%s, value=%+v, err=%v\n", string(data), tt2, err)

}

const (
	oneDaySeconds = 24 * 60 * 60
)

func test_floor() {

	now := time.Now().Unix()

	// floor date 2023-04-11 00:00:00
	// 今日开始，结束时间
	beg := int(now - (now % oneDaySeconds))
	end := beg + oneDaySeconds
	year, month, day := time.Unix(now, 0).Date()

	xBegin := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())

	fmt.Printf("x_begin=%v, %d\n", xBegin, xBegin.Unix())
	fmt.Printf("begin=%v, %d\n", time.Unix(int64(beg), 0), beg)

	// 昨日开始，结束时间
	yesBeg := beg - oneDaySeconds
	yesEnd := beg

	// 今日、昨日日期。
	day, yesDay := beg, yesBeg

	// for i := 30; i > 0; i-- {
	// 	start := beg - i*oneDaySeconds
	// 	stop := start + oneDaySeconds

	// 	fmt.Printf("i = %d, start=%d, stop=%d\n", i, start, stop)
	// }

	fmt.Printf("yes beg = %d, end = %d\n", yesBeg, yesEnd)
	fmt.Printf("now beg = %d, end = %d\n", beg, end)
	fmt.Printf("date now day = %d, yes day = %d\n", day, yesDay)
}

func main() {
	test_floor()
}

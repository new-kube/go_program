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

func main() {

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

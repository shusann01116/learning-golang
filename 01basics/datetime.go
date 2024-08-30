package basics

import (
	"fmt"
	"time"
	_ "time/tzdata" // Provide embedding tz data on binary
)

func DateTime() {
	now := time.Now()

	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)

	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))
}

func TZPreDefined() {
	now := time.Date(2024, time.August, 30, 0, 0, 0, 0, time.Local)
	past := time.Date(2024, time.August, 30, 0, 0, 0, 0, time.UTC)

	fmt.Println(now)
	fmt.Println(past)
}

func Duration() {
	// Duration の作り方
	// 1. pre-defined な Duration との積
	fiveMinute := 5 * time.Minute

	var seconds int = 10
	tenSecond := time.Duration(seconds) * time.Second

	fmt.Println(fiveMinute - tenSecond)

	// 2. time.Time 同士の Sub
	now := time.Date(2024, time.August, 30, 0, 0, 0, 0, time.Local)
	past := time.Date(2024, time.August, 30, 0, 0, 0, 0, time.UTC)
	fmt.Println(now.Sub(past))
}

func Truncate() {
	// 1 時間にまとめてバッチで読み込むファイル名を取得する
	filepath := time.Now().Truncate(time.Hour).Format("20060102050405.json")
	fmt.Println(filepath)

	// 5 分後と 5 分前の時刻
	fiveMinute := 5 * time.Minute
	fiveMinuteAfter := time.Now().Add(fiveMinute)
	fiveMinuteBefore := time.Now().Add(-fiveMinute)
	fmt.Println(fiveMinuteAfter)
	fmt.Println(fiveMinuteBefore)
}

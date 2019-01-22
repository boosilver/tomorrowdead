package main

import (
	"fmt"
	"strconv"
	"time"
)

func getServerTime() string {
	t := time.Now().Format("2006-01-02") //03:04:05
	return "Time now : " + t
}

func gettypetime(a string) string { //string ตัวแรก คือประกาศ type ของ a ตัวสองคือ type ของค่าที่ return กลับ
	var yy, mm, dd, aa = 0.0, 0.0, 0.0, 0.0
	aa, err := strconv.ParseFloat(a, 32)
	if err == nil {
		fmt.Println(aa)
	}

	y := time.Now().Format("2006")
	yy, err = strconv.ParseFloat(y, 32)
	if err == nil {
		fmt.Println(yy)
	}

	m := time.Now().Format("01")
	mm, err = strconv.ParseFloat(m, 32)
	if err == nil {
		fmt.Println(mm)
	}

	d := time.Now().Format("02")
	dd, err = strconv.ParseFloat(d, 32)
	if err == nil {
		fmt.Println(dd)
	}
	sum := yy + mm + dd
	suma := sum / aa
	r := fmt.Sprintf("%.4f", suma)

	return r
}

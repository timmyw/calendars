package calendars

import (
//	"fmt"
	"time"
	"errors"
	"strconv"
)

// Reference time (from time package) in the reference format.
const ReferenceTime = "20060102130405"

// Convert a "standard" date/time (without timezone) to a time.Time
// value
// 
// The "standard" format is: YYYYMMYYHHMMSS
func Convert_datetime(dt string) (time.Time, error) {

	// 01234567890123
	// YYYYMMYYHHMMSS
	
	if len(dt) < 14 {
		return *(new(time.Time)), errors.New("Date time too short")
	}

	Y, _ := strconv.Atoi(dt[0:4])
	M, _ := strconv.Atoi(dt[4:6])
	D, _ := strconv.Atoi(dt[6:8])
	h, _ := strconv.Atoi(dt[8:10])
	m, _ := strconv.Atoi(dt[10:12])
	s, _ := strconv.Atoi(dt[12:14])

	t := time.Date(Y, time.Month(M), D, h, m, s, 0, time.UTC)
	return t, nil
}

// Convert a time.Time to a universal (Julian) date
//
func Convert_to_universal(t time.Time) float64 {

	y := t.Year()
	m := int(t.Month())
	d := t.Day()

	my := (m - 14) / 12;
	iypmy := y + my

	var jdate float64

	jy := int64(1461 * (iypmy + 4800)) / 4
	jm := int64(367 * int64(m - 2 - 12 * my)) / 12
	jdate = float64(jy + jm - int64(3 * ((iypmy + 4900) / 100)) / 4 + int64(d) - 2432076)

	return jdate
}

func fsign(x float64) float64 {
	if x < 0.0 {
		return -1.0
	}
	if x == 0.0 {
		return 0.0
	}
	return 1.0
}

func ftrunc(x float64) float64 {
	return float64(int64(x))
}

func YMD_to_JD(Y int, M int, D int) float64 {

	var work float64
	y := float64(Y)
	m := float64(M)
	d := float64(D)
	work = 367.0 * y - ftrunc((7.0 * (y + ftrunc((m + 9.0) / 12.0))) / 4.0) + ftrunc((275.0 * m) / 9.0) + d + 1721013.5 - 0.5*fsign(100.0 * y + m - 190002.5) + 0.5
	return work
}

func HMS_to_UT(H int, M int, S int) float64 {
	m := float64(M)
	s := float64(S)
	h := float64(H) + (m / 60.0) + (s / 3600.0)

	return h / 24.0
}

func YMDHMS_to_JD(Y int, M int, D int, H int, Min int, S int) float64 {
	jd := YMD_to_JD(Y, M, D)
	ut := HMS_to_UT(H, Min, S)
	// fmt.Printf("%10.5f  %10.5f : %10.5f\n", jd, ut, jd+ut)

	return jd + ut
}

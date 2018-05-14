package calendars

import (
	"fmt"
	"math"
	"time"
	"errors"
	"strconv"
)

// Reference format
const ReferenceFormat = "%04d%02d%02d%02d%02d%02d"

// Reference time (from time package) in the reference format.
const ReferenceTime = "20060102130405"

type DateElements struct{
	Y int
	M int
	D int
	Hr int
	Min int
	Sec int
}

// Convert a "standard" date/time (without timezone) to a time.Time
// value
// 
// The "standard" format is: YYYYMMYYHHMMSS
func Convert_datetime(dt string) (time.Time, error) {

	// 01234567890123
	// YYYYMMYYHHMMSS
	
	if len(dt) < 12 {
		return *(new(time.Time)), errors.New("Date time too short")
	}

	Y, _ := strconv.Atoi(dt[0:4])
	M, _ := strconv.Atoi(dt[4:6])
	D, _ := strconv.Atoi(dt[6:8])
	h, _ := strconv.Atoi(dt[8:10])
	m, _ := strconv.Atoi(dt[10:12])
	s := 0
	if len(dt) >= 14 { s, _ = strconv.Atoi(dt[12:14]) }

	t := time.Date(Y, time.Month(M), D, h, m, s, 0, time.UTC)
	return t, nil
}

func Convert_to_datetime(t time.Time) string {
	Y, M, D := t.Year(), int(t.Month()), t.Day()
	Hr, Min, Sec := t.Hour(), t.Minute(), t.Second()

	return fmt.Sprintf(ReferenceFormat, Y, M, D, Hr, Min, Sec)
}

// Convert a time.Time to a universal (Julian) date
//
func Convert_to_universal(t time.Time) float64 {

	Y, M, D := t.Year(), int(t.Month()), t.Day()
	Hr, Min, Sec := t.Hour(), t.Minute(), t.Second()

	return YMDHMS_to_JD(Y, M, D, Hr, Min, Sec)
}

func Convert_from_universal(ut float64) time.Time {
	Y, M, D, Hr, Min, Sec := JD_to_YMDHMS(ut)
	return time.Date(Y, time.Month(M), D, Hr, Min, Sec, 0, time.UTC)
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

func YMDHMS_to_JD(Y int, M int, D int, Hr int, Min int, Sec int) float64 {
	jd := YMD_to_JD(Y, M, D)
	ut := HMS_to_UT(Hr, Min, Sec)
	// fmt.Printf("%10.5f  %10.5f : %10.5f\n", jd, ut, jd+ut)

	return jd + ut
}

func frac(x float64) float64 {
	return x - float64(int64(x))
}

// Convert a universal time (0 <= ut <= 24) to hours, minutes and
// seconds.  The non-fractional portion of the time is ignored.
func UT_to_HMS(ut float64) (int, int, int) {

	ut = frac(ut) * 24
	h := int(ut)
	ut = frac(ut) * 60
	m := int(ut)
	ut = frac(ut) * 60

	return h, m, int(ut + 0.5)
}

// Fix a float to a specified number of decimal places
func fix(x float64, decs int) float64 {
	y := int64(math.Pow(10.0, float64(decs)))
	z := int64(x * float64(y))
	fmt.Printf("%d\n", z)
	x = float64(int64(x * float64(y))) / float64(y)
	fmt.Printf("%10.5f\n", x)
	return x
}

// Convert a Julian date to Y, M, D
// From an algorithm pulled from:
// http://aa.usno.navy.mil/faq/docs/JD_Formula.php
func JD_to_YMD(JD float64) (int, int, int) {

	jd := JD + 0.5
	// fmt.Printf("inp: %10.5f\n", jd)
	L := int(jd + 68569)
	N := int(4 * L / 146097)
	L = L - (146097 * N + 3) / 4
	I := int(4000 * (L + 1) / 1461001)
	L = L - 1461 * I / 4+31
	J := int(80 * L / 2447)
	K := int(L - 2447 * J / 80)
	L = J / 11
	J = J + 2 - 12 * L
	I = 100 * (N - 49) + I + L

	// fmt.Printf("Day:%d\n", K)
	return I, J, K
}

// Convert a full Julian datetime to H,M,D and Hr, Min and Sec
func JD_to_YMDHMS(jd float64) (int, int, int, int, int, int) {

	H, M, D := JD_to_YMD(jd)
	Hr, Min, Sec := UT_to_HMS(jd - 0.5)

	return H, M, D, Hr, Min, Sec
}

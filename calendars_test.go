package calendars_test

import (

	"fmt"
	"time"
	
	"github.com/timmyw/calendars"
	
)

const dts1 = "19270506153107"

func ExampleConvert_to_universal() {
	t, _ := calendars.Convert_datetime(dts1)

	j := calendars.Convert_to_universal(t)

	fmt.Printf("%10.5f", j)
	// Output: 2425007.14661
}

func ExampleConvert_to_datetime() {
	t := time.Date(1927, 5, 6, 15, 31, 7, 0, time.UTC)

	fmt.Println(calendars.Convert_to_datetime(t))

	// Output: 19270506153107
}

func ExampleConvert_datetime() {

	t, err := calendars.Convert_datetime(dts1)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(t)

	// Output: 1927-05-06 15:31:07 +0000 UTC
}

func ExampleYMDHMS_to_JD() {
	Y := 1927
	M := 5
	D := 6
	Hr := 15
	Min := 37
	Sec := 15

	jd := calendars.YMDHMS_to_JD(Y, M, D, Hr, Min, Sec)
	fmt.Printf("%12.5f\n", jd)

	// Output: 2425007.15087
}

func ExampleYMD_to_JD() {
	Y := 1927
	M := 5
	D := 6

	jd := calendars.YMD_to_JD(Y, M, D)
	fmt.Printf("%7.1f\n", jd)

	// Output: 2425006.5
}

func ExampleHMS_to_UT() {
	H := 15
	M := 37
	S := 15

	ut := calendars.HMS_to_UT(H, M, S)
	fmt.Printf("%10.5f", ut)

	// Output: 0.65087
}

func ExampleUT_to_HMS() {

	ut := float64(0.65087)
	h, m, s := calendars.UT_to_HMS(ut)
	fmt.Printf("%02d:%02d:%02d", h, m, s)

	// Output: 15:37:15
}

func ExampleJD_to_YMD() {

	jd := 2425006.5

	y, m, d := calendars.JD_to_YMD(jd)
	fmt.Printf("%04d%02d%02d", y, m, d)

	// Output: 19270506
}

func ExampleJD_to_YMDHMS() {
	jd := 2425007.15087
	Y, M, D, Hr, Min, Sec := calendars.JD_to_YMDHMS(jd)
	fmt.Printf("%04d/%02d/%02d %02d:%02d:%02d\n", Y, M, D, Hr, Min, Sec)

	// Output: 1927/05/06 15:37:15
}

func ExampleConvert_datetime_2() {
	var dates = []string{ "197010091234", "20011212090522" }

	for _, date := range dates {
		// fmt.Printf("%s\n", date)

		t, err := calendars.Convert_datetime(date)
		if err != nil {
			fmt.Println(err)
		} else {
			ut := calendars.Convert_to_universal(t)
			// fmt.Printf("%10.5f\n", ut)

			t2 := calendars.Convert_from_universal(ut)

			fmt.Println(calendars.Convert_to_datetime(t2))
		}
	}

	// Output: 19701009123400
	// 20011212090522
}

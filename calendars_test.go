package calendars_test

import (

	"fmt"
	
	"timmyw/calendars"
	
)

const dts1 = "19270506153107"

/*
func ExampleConvert_to_universal() {
	t, _ := calendars.Convert_datetime(dts1)

	j := calendars.Convert_to_universal(t)

	fmt.Print(j)
	// Output: 99.99
}

func ExampleConvert_datetime() {

	t, err := calendars.Convert_datetime(dts1)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(t)

	// Output: 1927-05-06 15:31:07 +0000 UTC
}
*/

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

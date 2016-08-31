/*
  Copyright (c) 2012-2013 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package to

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestToString(t *testing.T) {

	if String(0) != "0" {
		t.Fatalf("Test failed.")
	}
	if String(-0) != "0" {
		t.Fatalf("Test failed.")
	}
	if String(1) != "1" {
		t.Fatalf("Test failed.")
	}
	if String(-1) != "-1" {
		t.Fatalf("Test failed.")
	}
	if String(10) != "10" {
		t.Fatalf("Test failed.")
	}
	if String(-10) != "-10" {
		t.Fatalf("Test failed.")
	}
	if String(int64(9223372036854775807)) != "9223372036854775807" {
		t.Fatalf("Test failed.")
	}
	if String(int64(-9223372036854775807)) != "-9223372036854775807" {
		t.Fatalf("Test failed.")
	}
	if String(uint64(18446744073709551615)) != "18446744073709551615" {
		t.Fatalf("Test failed.")
	}

}

func TestToBytes(t *testing.T) {

	if String(Bytes(0)) != "0" {
		t.Fatalf("Test failed.")
	}
	if String(Bytes(-0)) != "0" {
		t.Fatalf("Test failed.")
	}
	if String(Bytes(1)) != "1" {
		t.Fatalf("Test failed.")
	}
	if String(Bytes(-1)) != "-1" {
		t.Fatalf("Test failed.")
	}
	if String(Bytes(10)) != "10" {
		t.Fatalf("Test failed.")
	}
	if String(Bytes(-10)) != "-10" {
		t.Fatalf("Test failed.")
	}
	if String(Bytes(int64(9223372036854775807))) != "9223372036854775807" {
		t.Fatalf("Test failed.")
	}
	if String(Bytes(int64(-9223372036854775807))) != "-9223372036854775807" {
		t.Fatalf("Test failed.")
	}
	if String(Bytes(uint64(18446744073709551615))) != "18446744073709551615" {
		t.Fatalf("Test failed.")
	}

}

func TestFloating(t *testing.T) {
	if String(float32(1.1)) != "1.1" {
		t.Fatalf("Test failed.")
	}
	if String(float32(-1.1)) != "-1.1" {
		t.Fatalf("Test failed.")
	}
	if String(12345.12345) != "12345.12345" {
		t.Fatalf("Test failed.")
	}
	if String(-12345.12345) != "-12345.12345" {
		t.Fatalf("Test failed.")
	}
	if String(float64(-12345.12345)) != "-12345.12345" {
		t.Fatalf("Test failed.")
	}
}

func TestComplex(t *testing.T) {
	if String(complex(1, 1)) != "(1+1i)" {
		t.Fatalf("Test failed.")
	}
	if String(complex(1, -1)) != "(1-1i)" {
		t.Fatalf("Test failed.")
	}
	if String(complex(-1, -1)) != "(-1-1i)" {
		t.Fatalf("Test failed.")
	}
	if String(complex(-1, 1)) != "(-1+1i)" {
		t.Fatalf("Test failed.")
	}
	if String(true) != "true" {
		t.Fatalf("Test failed.")
	}
	if String(false) != "false" {
		t.Fatalf("Test failed.")
	}
	if String("hello") != "hello" {
		t.Fatalf("Test failed.")
	}
}

func TestNil(t *testing.T) {
	if String(nil) != "" {
		t.Fatalf("Test failed.")
	}
}

func TestBytes(t *testing.T) {
	if String([]byte{'h', 'e', 'l', 'l', 'o'}) != "hello" {
		t.Fatalf("Test failed.")
	}
}

func TestIntegers(t *testing.T) {
	if Int64(1) != int64(1) {
		t.Fatalf("Test failed.")
	}
	if Int64(-1) != int64(-1) {
		t.Fatalf("Test failed.")
	}
	if int32(Int64(true)) != int32(1) {
		t.Fatalf("Test failed.")
	}
	if int32(Int64(false)) != int32(0) {
		t.Fatalf("Test failed.")
	}
	if int32(Int64("123")) != int32(123) {
		t.Fatalf("Test failed.")
	}
	if uint32(Uint64("123")) != uint32(123) {
		t.Fatalf("Test failed.")
	}
	if uint32(Uint64("0")) != uint32(0) {
		t.Fatalf("Test failed.")
	}
	if uint(Uint64(5)) != uint(5) {
		t.Fatalf("Test failed.")
	}
	if uint(Uint64(5.1)) != uint(5) {
		t.Fatalf("Test failed.")
	}
	if int8(Uint64(6.1)) != int8(6) {
		t.Fatalf("Test failed.")
	}
}

func TestFloat(t *testing.T) {
	if float32(Float64(1)) != float32(1) {
		t.Fatalf("Test failed.")
	}
	if float32(Float64(1.2)) != float32(1.2) {
		t.Fatalf("Test failed.")
	}
	if Float64(-11.2) != float64(-11.2) {
		t.Fatalf("Test failed.")
	}
	if Float64("-11.2") != float64(-11.2) {
		t.Fatalf("Test failed.")
	}
}

func TestBool(t *testing.T) {
	if Bool("t") != true {
		t.Fatalf("Test failed.")
	}
	if Bool("FALSE") != false {
		t.Fatalf("Test failed.")
	}
	if Bool("0") != false {
		t.Fatalf("Test failed.")
	}
	if Bool("1") != true {
		t.Fatalf("Test failed.")
	}
	if Bool(1) != true {
		t.Fatalf("Test failed.")
	}
}

/*
// Delayed until Go 1.1
func TestList(t *testing.T) {
	mylist := []string{
		"a", "b", "c", "d", "e",
	}

	res := List(mylist)
	if res[2].(string) != "c" {
		t.Fatalf("Test failed.")
	}

	mylist1 := []int{
		1, 2, 3, 4, 5,
	}
	res1 := List(mylist1)
	if res1[2].(int) != 3 {
		t.Fatalf("Test failed.")
	}

	mylist2 := []interface{}{
		1, 2, 3, 4, 5,
	}
	res2 := List(mylist2)
	if res2[2].(int) != 3 {
		t.Fatalf("Test failed.")
	}
}

// Delayed until Go 1.1
func TestMap(t *testing.T) {
	mymap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
	}

	res := Map(mymap)

	if res["3"].(string) != "c" {
		t.Fatalf("Test failed.")
	}
}
*/

func TestConvert(t *testing.T) {
	b, _ := Convert(1, reflect.Bool)

	if b.(bool) != true {
		t.Fatalf("Test failed.")
	}

	i, _ := Convert("456", reflect.Int64)

	if i.(int64) != int64(456) {
		t.Fatalf("Test failed.")
	}

	i, _ = Convert("0", reflect.Int64)

	if i.(int64) != int64(0) {
		t.Fatalf("Test failed.")
	}

	ui, _ := Convert("456", reflect.Uint64)

	if ui.(uint64) != uint64(456) {
		t.Fatalf("Test failed.")
	}

	ui, _ = Convert("0", reflect.Uint64)

	if ui.(uint64) != uint64(0) {
		t.Fatalf("Test failed.")
	}

	var err error
	var bs interface{}

	if bs, err = Convert("string", reflect.Slice); err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(bs.([]byte), []byte("string")) != true {
		t.Fatalf("Test failed.")
	}

}

func TestTimeDuration(t *testing.T) {
	if Duration(123) != time.Duration(123) {
		t.Fatalf("Test failed.")
	}
	if Duration("12s37ms") != time.Second*12+time.Millisecond*37 {
		t.Fatalf("Test failed.")
	}
	if Duration("13:37") != time.Hour*13+time.Minute*37 {
		t.Fatalf("Test failed.")
	}
	if Duration("-13:37") != -(time.Hour*13 + time.Minute*37) {
		t.Fatalf("Test failed.")
	}
	if Duration("13:37:21") != time.Hour*13+time.Minute*37+time.Second*21 {
		t.Fatalf("Test failed.")
	}
	if Duration("13:37:21.456123") != time.Hour*13+time.Minute*37+time.Second*21+time.Microsecond*456123 {
		t.Fatalf("Test failed.")
	}
	if Duration("13:37:21.4561231") != time.Hour*13+time.Minute*37+time.Second*21+456123100 {
		t.Fatalf("Test failed.")
	}
	if Duration("13:37:21.456123789") != time.Hour*13+time.Minute*37+time.Second*21+time.Nanosecond*456123789 {
		t.Fatalf("Test failed.")
	}
	if Duration("13:37:21.456123789999") != time.Hour*13+time.Minute*37+time.Second*21+time.Nanosecond*456123789 {
		t.Fatalf("Test failed.")
	}
	if Duration("-13:37:21.456123789999") != -(time.Hour*13 + time.Minute*37 + time.Second*21 + time.Nanosecond*456123789) {
		t.Fatalf("Test failed.")
	}
	if Duration("abc") != time.Duration(0) {
		t.Fatalf("Test failed.")
	}
}

func TestDate(t *testing.T) {
	if time.Date(2012, 3, 24, 0, 0, 0, 0, time.Local).Equal(Time("2012-03-24")) != true {
		t.Fatalf("Test failed.")
	}
	if time.Date(2012, 3, 24, 0, 0, 0, 0, time.Local).Equal(Time("2012/03/24")) != true {
		t.Fatalf("Test failed.")
	}
	if time.Date(2012, 3, 24, 23, 13, 37, 0, time.Local).Equal(Time("2012-03-24 23:13:37")) != true {
		t.Fatalf("Test failed.")
	}
	if time.Date(2012, 3, 24, 23, 13, 37, 123, time.Local).Equal(Time("2012-03-24 23:13:37.000000123")) != true {
		t.Fatalf("Test failed.")
	}
	if time.Date(2012, 3, 24, 23, 13, 37, 0, time.Local).Equal(Time("03/24/2012 23:13:37")) != true {
		t.Fatalf("Test failed.")
	}
	if time.Date(2012, 3, 24, 23, 13, 37, 123, time.Local).Equal(Time("03/24/12 23:13:37.000000123")) != true {
		t.Fatalf("Test failed.")
	}
	if time.Date(2012, 3, 24, 23, 13, 37, 0, time.Local).Equal(Time("24/Mar/2012 23:13:37")) != true {
		t.Fatalf("Test failed.")
	}
	if time.Date(2012, 3, 24, 0, 0, 0, 0, time.Local).Equal(Time("Mar 24, 2012")) != true {
		t.Fatalf("Test failed.")
	}
	if time.Date(2012, 3, 24, 23, 13, 37, 123000000, time.UTC).Equal(Time("2012-03-24T23:13:37.123Z")) != true {
		t.Fatalf("Test failed.")
	}
	if time.Date(2012, 3, 24, 23, 13, 37, 123456789, time.UTC).Equal(Time("2012-03-24T23:13:37.123456789Z")) != true {
		t.Fatalf("Test failed.")
	}
}

func BenchmarkFmtIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s", 1)
	}
}

func BenchmarkFmtFloatToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s", 1.1)
	}
}

func BenchmarkStrconvIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(1)
	}
}

func BenchmarkStrconvFloatToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatFloat(1.1, 'f', 2, 64)
	}
}

func BenchmarkIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(1)
	}
}

func BenchmarkFloatToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(1.1)
	}
}

func BenchmarkIntToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(1)
	}
}

func BenchmarkBoolToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(true)
	}
}

func BenchmarkFloatToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(1.1)
	}
}

func BenchmarkIntToBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool(1)
	}
}

func BenchmarkStringToTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Time("2012-03-24")
	}
}

/*
func BenchmarkMap(b *testing.B) {
	mymap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
	}
	for i := 0; i < b.N; i++ {
		Map(mymap)
	}
}

func BenchmarkList(b *testing.B) {
	mylist := []string{
		"a", "b", "c", "d", "e",
	}
	for i := 0; i < b.N; i++ {
		List(mylist)
	}
}
*/

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Convert("567", reflect.Int64)
		if err != nil {
			b.Fatalf("Test failed.")
			return
		}
	}
}

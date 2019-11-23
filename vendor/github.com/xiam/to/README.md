# to

The `to` package provides quick-and-dirty conversions between built-in Go data
types.  When any conversion is unreasonable a [zero value][3] is used as
fallback.

If you're not working with human provided data, fuzzy input or if you'd rather
not ignore any error in your program, you should better use the standard Go
packages for conversion, such as [strconv][4], [fmt][5] or even [standard
conversion][6] they may be better suited for the task.

[![Build Status](https://travis-ci.org/xiam/to.svg?branch=master)](https://travis-ci.org/xiam/to)

## Installation

```sh
go get -u github.com/xiam/to
```

## Usage

Import the package

```go
import "github.com/xiam/to"
```

Use the available `to` functions to convert a `float64` into a `string`:

```go
// "1.23"
s := to.String(1.23)
```

Or a `bool` into `string`:

```go
// "true"
s := to.String(true)
```

What about the other way around? `string` to `float64` and `string` to `bool`.

```go
// 1.23
f := to.Float64("1.23")

// true
b := to.Bool("true")
```

Note that this package only provides `to.Uint64()`, `to.Int64()`,`to.Int()` and
`to.Float64()` but no `to.Uint8()`, `to.Uint()` or `to.Float32()` functions, if
you'd like to produce a `float32` instead of a `float64` you'd first use
`to.Float64()` and then cast the output using `float32()`.

```go
f32 := float32(to.Float64("12.34"))
```

There is another important function, `to.Convert()` that receives `interface{}`
as first argument and `reflect.Kind` as second:

```go
// val.(int64) = 12345
val, err := to.Convert("12345", reflect.Int64)
```

Date formats and durations are matched against common patterns and converted:

```go
timeVal := to.Time("2012-03-24")

timeVal := to.Time("Mar 24, 2012")

durationVal := to.Duration("12s37ms")
```

## Benchmarks

```
go test -bench=.

goos: linux
goarch: amd64
pkg: github.com/xiam/to
BenchmarkFmtIntToString-4         	11385524	       105 ns/op
BenchmarkFmtFloatToString-4       	 2266578	       531 ns/op
BenchmarkStrconvIntToString-4     	202011031	         5.88 ns/op
BenchmarkStrconvFloatToString-4   	 2515765	       474 ns/op
BenchmarkIntToString-4            	16711124	        77.0 ns/op
BenchmarkFloatToString-4          	 5648437	       214 ns/op
BenchmarkIntToBytes-4             	19158598	        70.9 ns/op
BenchmarkBoolToString-4           	167781417	         7.11 ns/op
BenchmarkFloatToBytes-4           	 6133180	       196 ns/op
BenchmarkIntToBool-4              	11871574	       106 ns/op
BenchmarkStringToTime-4           	  211500	      4997 ns/op
BenchmarkConvert-4                	13155015	        92.9 ns/op
PASS
ok  	github.com/xiam/to	17.887s
```

See the [docs][1] for a full reference of all the available `to` methods.

## License

This is Open Source released under the terms of the MIT License:

> Copyright (c) 2013-today José Nieto, https://xiam.dev
>
> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:
>
> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

[1]: http://godoc.org/github.com/xiam/to
[3]: http://golang.org/ref/spec#The_zero_value
[4]: http://golang.org/pkg/strconv/
[5]: http://golang.org/pkg/fmt/
[6]: http://golang.org/ref/spec#Conversions

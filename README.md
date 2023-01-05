[![GoDoc](https://godoc.org/github.com/qsymmachus/ranger?status.svg)](https://godoc.org/github.com/qsymmachus/ranger)

ranger 🧝‍♂️
=========

The missing range functions for Go. Generating a range of numbers or characters ([intervals](https://en.wikipedia.org/wiki/Interval_(mathematics))) is an inane but useful task – perfect for a computer.

The package includes functions for generating ranges of the following types:

* __Integers__: `int`, `int8`, `int16`, `int64`, `uint`, `uint8`, `uint16`, `uint64`
* __Floats__: `float32`, `float64`
* __Characters__: `rune`, `string`

The `Range[N]` function is parameterized by the type `N` which can be any integer or float type. This requires Go v1.18 to support generics; for older versions of Go use v0.0.1 of this package.

Examples
--------

```go
// Integer intervals
// ----------------
numbers := ranger.Range[int](1, 10)
fmt.Printf("%v", numbers) // [1 2 3 4 5 6 7 8 9 10]

odds := ranger.Range[int](1, 10, ranger.Step(2))
fmt.Printf("%v", odds) // [1 3 5 7 9]

littleNumbers := ranger.Range[int8](1, 127)
fmt.Printf("%v", littleNumbers) // [1 2 3 4...for a little while]

bigNumbers := ranger.Range[int64](1, 1000000)
fmt.Printf("%v", bigNumbers) // [1 2 3 4...for a long while]

// Float intervals
// ---------------
floats := ranger.Range[float32](1, 2.5, ranger.Step(0.5))
fmt.Printf("%v", floats) // [1.0 1.5 2.0 2.5]

// Character intervals
// -------------------
arrows := ranger.Rune('←', '↓')
fmt.Printf("%v", arrows) // [8592 8593 8594 8595] Runes are just numbers

alphabet, err := ranger.String("a", "z")
if err != nil {
  fmt.Println("Start or end string wasn't exactly one character")
}
for _, letter := range alphabet {
  fmt.Println(letter) // a b c d e f g...!
}
fmt.Println("Now I know my ABCs!")
```

Development
-----------

To run all tests:

```
go test ./...
```

To run benchmarks:

```
go test -bench ./...
```

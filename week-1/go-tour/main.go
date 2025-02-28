package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"time"
)

// ? by convention the package name is same as the last element of import path for eg. math/rand then I can use rand in order to access the func

// ? can declare multiple variable of same type in a row type at end
func Add(t, s string, x, y int) int {
	fmt.Println(s, t)
	return x + y
}

// ? can return muliple results
func Swap(x, y string) (string, string) {
	return y, x
}

// ? naked return where you I can name my return vars in function; good to use in short functions not for bigger functions since it hinders readability
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// ? Types
/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte  alias for uint8

rune  alias for int32   represents a Unicode code point

float32 float64

complex64 complex128 */

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	Z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// ? const can be char, string, bool, number it cannot be declare using := similar to var

const PI = 3.13

func Summation(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func FooBar(n int) {
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Foobar")
	} else if n%3 == 0 {
		fmt.Println("Foo")
	} else if n%5 == 0 {
		fmt.Println("Bar")
	} else {
		fmt.Println("You are far away from mordor")
	}
}

func Pow(x, n, lim float64) float64 {

	// ? here the scope of v is for that if block only can't access it outside the if world

	if v := math.Pow(x, n); v <= lim {
		return v
	}
	return lim
}

func OSguess() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func SaturdayFinder() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today 🥳")
	case today + 1:
		fmt.Println("tomorrow 😁")
	case today + 3:
		fmt.Println("day after tomorrow 😒")
	default:
		fmt.Println("dreams away my friend 🥹")
	}
}

// ? in order to write long if else chain we can write switch {} directly which will be marked as truth

func A() {
	i := 0
	// ! defer evalutation begins when
	defer fmt.Println("after defer", i)
	i++
	fmt.Println("before defer", i)
}

func B() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func C() (i int) {
	defer func() { i++ }()
	return 1
}

func D(i int) {
	defer fmt.Println("defered in D")
	if i == 0 {
		panic("something went wrong")
	} else {

		fmt.Println("This won't run")
	}
}

func E() {
	defer fmt.Println("defered in E")
	D(0)
}

func main() {


	// fmt.Println("Hello World ", rand.Uint32())

	// ? inside the function i can use :=

	// a, b := Swap("bob", "anon")
	// fmt.Printf("bob and anon swapped and became %v and %v\n", a, b)

	// fmt.Println(
	// 	Pow(3, 2, 10),
	// 	Pow(3, 3, 20),
	// )

	// SaturdayFinder()

	// OSguess()

	// fmt.Println(Summation(10))

	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", Z, Z)

	// var c, python, java bool
	// var i, j int = 1, 2
	// fmt.Println(c, python, java, i, j)
}

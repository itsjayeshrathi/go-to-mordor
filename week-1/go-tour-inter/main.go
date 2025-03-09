package main

import (
	"fmt"
	"math"

	"itsjayeshrathi.dev/go/go-tour-inter/interfaces"
)

type Vertex struct {
	X float64
	Y float64
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

// ? slices are just a view and not full fledged container they are like magic mirror which can show the array elements they are storing as well as you can modify those thngs

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ? I declared a type over here in order to add method to any type it should be in same PACKAGE

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	i, j := 42, 130
	ptr := &i
	*ptr = 23
	fmt.Println(i, j, *ptr)
	v := Vertex{2, 3}
	fmt.Println(v)
	fmt.Println(v1, v2, v3, *p)

	primes := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes, len(primes), cap(primes))
	s1 := primes[0:3]
	s2 := primes[4:]
	s2[0] = 0
	s2 = append(s2, 30)
	fmt.Println(primes, s1, s2)

	a := make([]int, 5)
	fmt.Println(a)

	fmt.Printf("%.2f", v.Abs())

	fmt.Println("===============================")

	var inter interfaces.I
	var t *interfaces.T
	inter = t
	interfaces.Describe(inter)
	inter.M()
	inter = &interfaces.T{S: "hello"}

	interfaces.Describe(inter)
	inter.M()

}

// ? range is only for slices and maps
func Add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// package main

// import "golang.org/x/tour/pic"

// func Pic(dx, dy int) [][]uint8 {
//     // Allocate outer slice
//     img := make([][]uint8, dy)

//     for i := range dy {
//         // Allocate inner slice
//         img[i] = make([]uint8, dx)
//         for j := range dx {
// 			img[i][j] = uint8((i * j)) // Example pattern
//         }
//     }
//     return img
// }

// func main() {
// 	pic.Show(Pic)
// }

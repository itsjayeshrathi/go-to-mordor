package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice, cap(slice))
	clear(slice)
	slice = append(slice, 10, 11, 12)
	fmt.Println(slice, cap(slice))
	fmt.Println(slices.Max(slice))
	temp := "name"
	for i:= range len(temp){
		fmt.Println(string(temp[i]))
	}
}

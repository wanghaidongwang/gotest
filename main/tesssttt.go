package main

import "fmt"

func main() {
//	var numbers = make([]int,3,5)
    var numbers  []int
   // numbers =
   numbers[0] = 1
   numbers[1] = 2
	printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
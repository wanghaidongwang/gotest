package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 11
	ch <- 12
	ch <- 13
	//close(ch)
	b := <-ch
	//b:=  <- ch
	//<- ch
	//for x := range ch {
	//	fmt.Println(x)
	//}
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(b)
	//fmt.Println(a)
	//x, ok := <- ch
	//fmt.Println(x, ok)
}

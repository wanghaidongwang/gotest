package main

import "fmt"

func main ()  {
	ch := make(chan int, 10)
	ch <- 11
	ch <- 12
	ch <- 13
	//close(ch)
	b,a:=  <- ch,<-ch
	//b:=  <- ch
	y:=  <- ch
	//for x := range ch {
	//	fmt.Println(x)
	//}
    fmt.Println(y)
	fmt.Println(b)
	fmt.Println(a)
	//x, ok := <- ch
	//fmt.Println(x, ok)
}

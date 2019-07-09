
package main

import "fmt"

func main() {
	arr := []string{"hello", "world"}
	fmt.Println(test(arr))
}

func test(arr []string) string {
	return arr[0]
}




package main

import "fmt"

func main()  {
	type test struct {
 		ip  string
 		potr int
	}

	type v struct{
		name string
		password  string
        a []test
	}
	var (
		//var j v
		//var a [2]test
		j = v{
			name:     "hahaha",
			password: "ieuhr",
			a:        []test{
				{"adada",43},
				{"bbbbb",56},
			},
		} //j.a[1].ip="gagaga"
	)
	//j.a[1].potr=32

    fmt.Println(j)

}

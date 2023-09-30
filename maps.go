package main

import (
	"fmt"
	"maps"
)

func main(){
	m := make(map[string]int)

	m["k1"] = 8

	fmt.Println("map:", m)

	v := m["k1"]
	fmt.Println("v:", v)

	fmt.Println("len:", len(m))

	delete(m, "k1")
	fmr.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k1"]
	fmt.Println("prs:", prs)
	

	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)


    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }

}
package main

import (
	"fmt"
	"time"

	"github.com/gbeltramo/go-23d/internal/load23d"
)

func main() {
	fmt.Println("Decimate")

	t0 := time.Now()
	triang, err := load23d.LoadSTL("./data/bunny.stl")
	fmt.Printf("Time bunny.stl: %v\n", time.Since(t0))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("The length of the triangulation in bunny is %d\n", len(triang.Tri))

	for idx := 0; idx < 3; idx++ {
		fmt.Printf("triang.Items[idx]: %v\n---\n", triang.Tri[idx])
	}
}

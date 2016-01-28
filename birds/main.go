package main

import (
	"github.com/rahart/learning-go/birds/bird"
)


func main() {
	d := bird.Duck{ Name: "Mallard"}
	d2 := bird.Duck{ Name: "Domestic"}
	p := bird.Penguin{ Name: "Emperor"}
	h := bird.Hawk{ Name: "Red-tailed"}

	bird.Info(d, d2, p, h)
	bird.Version()

}

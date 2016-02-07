package main

import (
	"fmt"

	"github.com/rahart/learning-go/structs/config"
)

func main() {
	m := map[string]string{
		"bind": "0.0.0.0",
		"port": "8080",
	}

	c, err := config.Build(m)
	if err != nil {
		fmt.Println(err)
	}
	// c := new(config.Config)
	// c.P = config.Parameters(m)

	fmt.Printf("%T, %v \n", c.P, c)
	for _, v := range c.P {
		fmt.Println(v.Key, " ", v.Value)
	}
}

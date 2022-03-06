package main

import (
	"aproxiable"
	"fmt"
)

func main() {

	a, err := aproxiable.NewAproxiableFromConfigPath("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Got an error back when creating a new aproxiable %v", err))
	}

	a.Start()
}

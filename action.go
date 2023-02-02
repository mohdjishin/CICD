package main

import (
	"fmt"
	"runtime"
)

func Demo() {
	fmt.Println("GO version ", runtime.Version())

	fmt.Println("GOOS ", runtime.GOOS)
	fmt.Println("GOARCH ", runtime.GOARCH)

}

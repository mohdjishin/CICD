package main

import (
	"fmt"
	"runtime"

	"rsc.io/quote"
)

func Demo() {
	fmt.Println("GO version ", runtime.Version())

	fmt.Println("GOOS ", runtime.GOOS)
	fmt.Println("GOARCH ", runtime.GOARCH)
	fmt.Println(quote.Go())
}

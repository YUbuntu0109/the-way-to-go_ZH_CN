package main

import (
	"fmt"
	"runtime"
	"os"
)

func main(){
	// var goos string = runtime.GOOS
	goos := runtime.GOOS
	// Printf(format string, list of variables to be printed)
	fmt.Printf("the opertating system is : %s\n", goos)
	var path string = os.Getenv("PATH")
	// path := os.Getenv("PATH")
	fmt.Printf("Path is : %s\n", path)
}

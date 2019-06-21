package main

import (
	"fmt"
	"runtime"
)

func main() {

	//returns the same number, do they do the same thing ?  test2.exe added last line
	fmt.Printf("number of threads: %v\n", runtime.GOMAXPROCS(-1))
	fmt.Printf("number of CPU's: %v\n", runtime.NumCPU())
	fmt.Printf("number of threads: %v\n", runtime.GOMAXPROCS(runtime.NumCPU()))
}

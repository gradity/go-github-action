package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("hello from docker action")
	flag.Parse()
	first := flag.Arg(0)
	// second := flag.Arg(1)
	msg := fmt.Sprintf("hello %v", first)
	fmt.Printf("::set-output name=message::%v", msg)
}

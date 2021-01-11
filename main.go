package main

import (
	"fmt"
	"flag"
)

const VERSION = "1.0"

func main() {
	var version = flag.Bool("version", true, "gives program version")
	flag.Parse()
	fmt.Println("Hello, World!")
	if *version == true {
	fmt.Println(VERSION)
	}
}

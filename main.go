package main

import (
	"fmt"
)

func print(a ...any) {
	fmt.Print(a...)
}

func main() {
	print("hoge", 10, "\n")
}

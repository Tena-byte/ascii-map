package main

import (
	"ascii/ascii"
	"fmt"
	"os"
)

func main() {


	font := ascii.LoadBanner("standard")

	art := ascii.Render(os.Args[1], font)

	for _, line := range art{

		fmt.Println(line)
	}
}

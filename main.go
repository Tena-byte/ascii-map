package main

import (
	"ascii/ascii"
	"fmt"
	"os"
)

func main() {

	if !ascii.ValidateInput(os.Args[1]){
		fmt.Println("invalid input")
		return
	}

	font := ascii.LoadBanner("standard")

	art := ascii.Render(os.Args[1], font)
	
	art = ascii.AddBorder(art)

	for _, line := range art{

		fmt.Println(line)
	}
}

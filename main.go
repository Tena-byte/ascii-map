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

	if len(os.Args) < 3{
		fmt.Println("Usage : go run . <text> <color>")
		return
	}

	font := ascii.LoadBanner("standard")

	art := ascii.Render(os.Args[1], font)

	art = ascii.AddBorder(art)
	
	art = ascii.AsciiColor(art, os.Args[2])

	for _, line := range art{

		fmt.Println(line)
	}
}

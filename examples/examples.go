package main

import (
	"fmt"

	"github.com/engmtcdrm/go-prettyprint"
)

// TODO: Add centering for text
// TODO: Fix the loop for outputting colors

func main() {
	fmt.Print("Standard: ")

	for i := 0; i < 8; i++ {
		fmt.Print(prettyprint.Bg8Bit(i, fmt.Sprintf("%-5v", i)) + "  ")
	}

	fmt.Println()
	fmt.Print("Intense:  ")

	for i := 8; i < 16; i++ {
		fmt.Print(prettyprint.Bg8Bit(i, fmt.Sprintf("%-5v", i)) + "  ")
	}

	fmt.Println()
	fmt.Println()

	for i := 1; i < 16; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(prettyprint.Bg8Bit((16*i)+j, fmt.Sprintf("%-5v", (16*i)+j)) + "  ")
		}

		fmt.Println()
	}
}

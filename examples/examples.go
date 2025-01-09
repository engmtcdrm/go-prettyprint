package main

import (
	"fmt"
	"strconv"
	"strings"

	pp "github.com/engmtcdrm/go-prettyprint"
)

// centerStr centers a string within a given width by padding with spaces
// on the left and right. If the string is already longer than the width,
// it is returned as-is.
func centerStr(s string, width int) string {
	if len(s) >= width {
		return s
	}

	// Pad left and right with spaces, then trim to desired width
	return fmt.Sprint(strings.Repeat(" ", (width-len(s))/2), s, strings.Repeat(" ", width))[0:width]
}

func main() {
	padding := 8

	fmt.Print("Standard: ")

	for i := 0; i < 8; i++ {
		fmt.Print(pp.Bg8Bit(i, centerStr(strconv.Itoa(i), padding)))
	}

	fmt.Println()
	fmt.Print("Intense:  ")

	for i := 8; i < 16; i++ {
		fmt.Print(pp.Bg8Bit(i, centerStr(strconv.Itoa(i), padding)))
	}

	fmt.Println()
	fmt.Println()

	d := 0
	d2 := 0

	for i := 0; i < 160; i += 72 {
		for j := 16 + i; j < 16+i+36; j++ {
			fmt.Print(pp.Bg8Bit(j, centerStr(strconv.Itoa(j), padding)))

			d++

			if d == 6 {
				fmt.Print("  ")
				for k := 31; k < 37; k++ {
					fmt.Print(pp.Bg8Bit(j+k, centerStr(strconv.Itoa(j+k), padding)))
				}

				fmt.Println()
				d = 0
				d2++
			}

			if d2 == 6 {
				fmt.Println()
				d2 = 0
			}
		}
	}

	fmt.Print("Greyscale: ")

	for i := 232; i < 244; i++ {
		fmt.Print(pp.Bg8Bit(i, centerStr(strconv.Itoa(i), padding)))
	}

	fmt.Println()
	fmt.Print(strings.Repeat(" ", 11))

	for i := 244; i < 256; i++ {
		fmt.Print(pp.Bg8Bit(i, centerStr(strconv.Itoa(i), padding)))
	}

	fmt.Println()
}

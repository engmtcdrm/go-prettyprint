package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

func Print8BitColors() {
	padding := 8

	start := time.Now()

	cStr := "Standard: "

	for i := 0; i < 8; i++ {
		cStr += pp.Bg8Bit(i, centerStr(strconv.Itoa(i), padding))
	}

	cStr += "\nIntense:  "

	for i := 8; i < 16; i++ {
		cStr += pp.Bg8Bit(i, centerStr(strconv.Itoa(i), padding))
	}

	cStr += "\n\n"

	d := 0
	d2 := 0

	for i := 0; i < 160; i += 72 {
		for j := 16 + i; j < 16+i+36; j++ {
			cStr += pp.Bg8Bit(j, centerStr(strconv.Itoa(j), padding))

			d++

			if d == 6 {
				cStr += "  "
				// fmt.Print("  ")
				for k := 31; k < 37; k++ {
					cStr += pp.Bg8Bit(j+k, centerStr(strconv.Itoa(j+k), padding))
				}

				cStr += "\n"
				d = 0
				d2++
			}

			if d2 == 6 {
				cStr += "\n"
				d2 = 0
			}
		}
	}

	cStr += "Greyscale: "

	for i := 232; i < 244; i++ {
		cStr += pp.Bg8Bit(i, centerStr(strconv.Itoa(i), padding))
	}

	cStr += "\n           "

	for i := 244; i < 256; i++ {
		cStr += pp.Bg8Bit(i, centerStr(strconv.Itoa(i), padding))
	}

	fmt.Println(cStr)
	fmt.Println()

	end := time.Now()

	fmt.Println("Time taken:", end.Sub(start))
}

func determineColorSlice(n int) []int {
	if n < 0 || n > 255 {
		s := []int{}

		for i := 0; i < 256; i++ {
			s = append(s, i)
		}

		return s
	}

	return []int{n}
}

func Print24BitColors(r int, g int, b int) {
	rSlice := determineColorSlice(r)
	gSlice := determineColorSlice(g)
	bSlice := determineColorSlice(b)

	d := 0
	cStr := ""

	start := time.Now()

	for _, r := range rSlice {
		for _, g := range gSlice {
			for _, b := range bSlice {
				cStr += pp.Bg24Bit(r, g, b, centerStr(fmt.Sprintf("%d,%d,%d", r, g, b), 13))

				d++

				if d == 8 {
					cStr += "\n"
					d = 0
				}
			}

			cStr += "\n"
			fmt.Print(cStr)
			cStr = ""
		}
	}

	end := time.Now()

	fmt.Println("Time taken:", end.Sub(start))
}

func main() {
	Print8BitColors()
	Print24BitColors(-1, -1, -1)
}

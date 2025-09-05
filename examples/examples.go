package main

import (
	"fmt"

	pp "github.com/engmtcdrm/go-prettyprint"
)

func doColor(color string, colorFunc func(...any) string, colorFormatFunc func(string, ...any) string) {
	fmt.Println(colorFunc("This is the ", color, " func"))
	fmt.Println(colorFormatFunc("This is a %s text: %d", color, 123))
	fmt.Printf("Only %s is %s\n", colorFunc("this part"), color)
	fmt.Println()
}

func main() {
	doColor("Bold", pp.Bold, pp.Boldf)

	// Foreground colors
	doColor("Black", pp.Black, pp.Blackf)
	doColor("Green", pp.Green, pp.Greenf)
	doColor("Yellow", pp.Yellow, pp.Yellowf)
	doColor("Blue", pp.Blue, pp.Bluef)
	doColor("Red", pp.Red, pp.Redf)
	doColor("Magenta", pp.Magenta, pp.Magentaf)
	doColor("Cyan", pp.Cyan, pp.Cyanf)
	doColor("White", pp.White, pp.Whitef)

	// Bright Foreground Colors
	doColor("Bright Black", pp.IntenseBlack, pp.IntenseBlackf)

	// Background colors
	doColor("black bg", pp.BlackBg, pp.BlackBgf)
	doColor("green bg", pp.GreenBg, pp.GreenBgf)
	doColor("yellow bg", pp.YellowBg, pp.YellowBgf)
	doColor("blue bg", pp.BlueBg, pp.BlueBgf)
	doColor("red bg", pp.RedBg, pp.RedBgf)
	doColor("magenta bg", pp.MagentaBg, pp.MagentaBgf)
	doColor("cyan bg", pp.CyanBg, pp.CyanBgf)
	doColor("white bg", pp.WhiteBg, pp.WhiteBgf)

	fmt.Println(pp.Complete("All done!"))
	fmt.Println(pp.Completef("All done with %d tasks!", 5))
	fmt.Println(pp.Alert("This is an alert message!"))
	fmt.Println(pp.Alertf("This is an alert message with code %d!", 42))
	fmt.Println(pp.Info("This is an info message!"))
	fmt.Println(pp.Infof("This is an info message with code %d!", 100))
	fmt.Println(pp.RedAlert("This is a red alert message!"))
	fmt.Println(pp.RedAlertf("This is a red alert message with code %d!", 500))
	fmt.Println(pp.Fail("This is a fail message!"))
	fmt.Println(pp.Failf("This is a fail message with code %d!", 1))
	fmt.Println(pp.Var("variable", "value"))
	fmt.Println(pp.VarDoubleQuote("variable", "value"))
	fmt.Println(pp.VarSingleQuote("variable", "value"))

	fmt.Println(pp.Format("This is a bold and red text", pp.Bold, pp.Red))
	fmt.Println(pp.Format("This is a green text with yellow background", pp.Green, pp.YellowBg))
	fmt.Println(pp.Format("This is a bold, blue text with white background", pp.Bold, pp.Blue, pp.WhiteBg))
	fmt.Println(pp.Format("This is a bold, magenta text with cyan background", pp.Bold, pp.Magenta, pp.CyanBg))
	fmt.Println(pp.Foreground24Bit(255, 130, 80, "This is a normal text"))
}

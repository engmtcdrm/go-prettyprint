## Description

The `prettyprint` package provides utilities for formatting strings with ANSI color codes. It includes predefined icons for different message types and functions to format strings with foreground and background colors.

Key Features:

- **Standard Color Formatting**: 8 basic foreground colors (`Black`, `Red`, `Green`, `Yellow`, `Blue`, `Magenta`, `Cyan`, `White`) and background colors (`BlackBg`, `RedBg`, etc.) with both variadic (`Red`) and formatted (`Redf`) versions.
- **Intense/Bright Color Support**: High-intensity variants of all 8 colors for both foreground (`IntenseRed`, `IntenseGreen`, etc.) and background (`IntenseRedBg`, `IntenseGreenBg`, etc.) with formatted versions.
- **8-bit Color Support**: Access to the full 256-color palette with `Fg8Bit`/`Bg8Bit` functions for precise color control (0-255 range).
- **24-bit True Color Support**: Full RGB color support with `Fg24Bit`/`Bg24Bit` functions allowing any RGB combination (0-255 for each component).
- **Text Formatting**: Style text with `Bold`, `Dim`, `Italic`, `Underline`, `DoubleUnderline`, `Strike`, and `Reverse` formatting, each with formatted versions.
- **Advanced Formatting**: `Format` function to apply multiple formatting functions in sequence.
- **Predefined Message Icons**: Pre-styled icons for common message types (`Complete` ✓, `Alert` !, `Fail` ✗, `Info` i) with automatic color coding.
- **Variable Display Helpers**: Functions like `Var`, `VarSingleQuote`, and `VarDoubleQuote` for displaying variable names and values with consistent formatting.
- **Input Validation**: All color functions validate input ranges and gracefully fall back to plain text for invalid values.

## Installation

prettyprint is available using the standard `go get` command.

Install by running:

```bash
go get github.com/engmtcdrm/go-prettyprint
```

Run tests by running:

```bash
go test github.com/engmtcdrm/go-prettyprint
```

## Usage

Using the alias `pp` makes using the package less verbose.

```go
import pp "github.com/engmtcdrm/go-prettyprint"
```

## Index

- [Description](#description)
- [Installation](#installation)
- [Usage](#usage)
- [Index](#index)
- [Variables](#variables)
- [func Alert](#func-alert)
- [func Alertf](#func-alertf)
- [func Bg24Bit](#func-bg24bit)
- [func Bg24Bitf](#func-bg24bitf)
- [func Bg8Bit](#func-bg8bit)
- [func Bg8Bitf](#func-bg8bitf)
- [func Black](#func-black)
- [func BlackBg](#func-blackbg)
- [func BlackBgf](#func-blackbgf)
- [func Blackf](#func-blackf)
- [func Blue](#func-blue)
- [func BlueBg](#func-bluebg)
- [func BlueBgf](#func-bluebgf)
- [func Bluef](#func-bluef)
- [func Bold](#func-bold)
- [func Boldf](#func-boldf)
- [func Complete](#func-complete)
- [func Completef](#func-completef)
- [func Cyan](#func-cyan)
- [func CyanBg](#func-cyanbg)
- [func CyanBgf](#func-cyanbgf)
- [func Cyanf](#func-cyanf)
- [func Dim](#func-dim)
- [func Dimf](#func-dimf)
- [func DoubleUnderline](#func-doubleunderline)
- [func DoubleUnderlinef](#func-doubleunderlinef)
- [func Fail](#func-fail)
- [func Failf](#func-failf)
- [func Fg24Bit](#func-fg24bit)
- [func Fg24Bitf](#func-fg24bitf)
- [func Fg8Bit](#func-fg8bit)
- [func Fg8Bitf](#func-fg8bitf)
- [func Format](#func-format)
- [func Green](#func-green)
- [func GreenBg](#func-greenbg)
- [func GreenBgf](#func-greenbgf)
- [func Greenf](#func-greenf)
- [func Info](#func-info)
- [func Infof](#func-infof)
- [func IntenseBlack](#func-intenseblack)
- [func IntenseBlackBg](#func-intenseblackbg)
- [func IntenseBlackBgf](#func-intenseblackbgf)
- [func IntenseBlackf](#func-intenseblackf)
- [func IntenseBlue](#func-intenseblue)
- [func IntenseBlueBg](#func-intensebluebg)
- [func IntenseBlueBgf](#func-intensebluebgf)
- [func IntenseBluef](#func-intensebluef)
- [func IntenseCyan](#func-intensecyan)
- [func IntenseCyanBg](#func-intensecyanbg)
- [func IntenseCyanBgf](#func-intensecyanbgf)
- [func IntenseCyanf](#func-intensecyanf)
- [func IntenseGreen](#func-intensegreen)
- [func IntenseGreenBg](#func-intensegreenbg)
- [func IntenseGreenBgf](#func-intensegreenbgf)
- [func IntenseGreenf](#func-intensegreenf)
- [func IntenseMagenta](#func-intensemagenta)
- [func IntenseMagentaBg](#func-intensemagentabg)
- [func IntenseMagentaBgf](#func-intensemagentabgf)
- [func IntenseMagentaf](#func-intensemagentaf)
- [func IntenseRed](#func-intensered)
- [func IntenseRedBg](#func-intenseredbg)
- [func IntenseRedBgf](#func-intenseredbgf)
- [func IntenseRedf](#func-intenseredf)
- [func IntenseWhite](#func-intensewhite)
- [func IntenseWhiteBg](#func-intensewhitebg)
- [func IntenseWhiteBgf](#func-intensewhitebgf)
- [func IntenseWhitef](#func-intensewhitef)
- [func IntenseYellow](#func-intenseyellow)
- [func IntenseYellowBg](#func-intenseyellowbg)
- [func IntenseYellowBgf](#func-intenseyellowbgf)
- [func IntenseYellowf](#func-intenseyellowf)
- [func Italic](#func-italic)
- [func Italicf](#func-italicf)
- [func Magenta](#func-magenta)
- [func MagentaBg](#func-magentabg)
- [func MagentaBgf](#func-magentabgf)
- [func Magentaf](#func-magentaf)
- [func Red](#func-red)
- [func RedAlert](#func-redalert)
- [func RedAlertf](#func-redalertf)
- [func RedBg](#func-redbg)
- [func RedBgf](#func-redbgf)
- [func Redf](#func-redf)
- [func Reverse](#func-reverse)
- [func Reversef](#func-reversef)
- [func Strike](#func-strike)
- [func Strikef](#func-strikef)
- [func Underline](#func-underline)
- [func Underlinef](#func-underlinef)
- [func Var](#func-var)
- [func VarDoubleQuote](#func-vardoublequote)
- [func VarQuote](#func-varquote)
- [func VarSingleQuote](#func-varsinglequote)
- [func White](#func-white)
- [func WhiteBg](#func-whitebg)
- [func WhiteBgf](#func-whitebgf)
- [func Whitef](#func-whitef)
- [func Yellow](#func-yellow)
- [func YellowBg](#func-yellowbg)
- [func YellowBgf](#func-yellowbgf)
- [func Yellowf](#func-yellowf)


## Variables

<a name="IconComplete"></a>

```go
var (
    IconComplete = "[✓]" // IconComplete is the icon for a completed message.
    IconAlert    = "[!]" // IconAlert is the icon for an alerted message.
    IconFailed   = "[✗]" // IconFailed is the icon for a failed message.
    IconInfo     = "[i]" // IconInfo is the icon for an informational message.
)
```

<a name="Alert"></a>
## func [Alert](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L23>)

```go
func Alert(msg string) string
```

Alert returns a string with a yellow exclamation icon and the given message.

<a name="Alertf"></a>
## func [Alertf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L28>)

```go
func Alertf(msg string, a ...any) string
```

Alertf returns a string with a yellow exclamation icon and the given formatted message.

<a name="Bg24Bit"></a>
## func [Bg24Bit](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L251>)

```go
func Bg24Bit(r int, g int, b int, a ...any) string
```

Bg24Bit formats using the default formats for its operands and returns the resulting string with an 24\-bit background color. Spaces are added between operands when neither is a string.

r, g, and b must be between 0 and 255 otherwise no color will be applied.

<a name="Bg24Bitf"></a>
## func [Bg24Bitf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L265>)

```go
func Bg24Bitf(r int, g int, b int, format string, a ...any) string
```

Bg24Bitf formats according to a format specifier and returns the resulting string with an 24\-bit background color.

r, g, and b must be between 0 and 255 otherwise no color will be applied.

<a name="Bg8Bit"></a>
## func [Bg8Bit](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L222>)

```go
func Bg8Bit(color int, a ...any) string
```

Bg8Bit formats using the default formats for its operands and returns the resulting string with an 8\-bit background color. Spaces are added between operands when neither is a string.

color must be between 0 and 255 otherwise no color will be applied.

<a name="Bg8Bitf"></a>
## func [Bg8Bitf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L236>)

```go
func Bg8Bitf(color int, format string, a ...any) string
```

Bg8Bitf formats according to a format specifier and returns the resulting string with an 8\-bit background color.

color must be between 0 and 255 otherwise no color will be applied.

<a name="Black"></a>
## func [Black](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L12>)

```go
func Black(a ...any) string
```

Black formats using the default formats for its operands and returns the resulting string with a black foreground. Spaces are added between operands when neither is a string.

<a name="BlackBg"></a>
## func [BlackBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L12>)

```go
func BlackBg(a ...any) string
```

BlackBg formats using the default formats for its operands and returns the resulting string with a black background. Spaces are added between operands when neither is a string.

<a name="BlackBgf"></a>
## func [BlackBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L18>)

```go
func BlackBgf(format string, a ...any) string
```

BlackBgf formats according to a format specifier and returns the resulting string with a black background.

<a name="Blackf"></a>
## func [Blackf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L18>)

```go
func Blackf(format string, a ...any) string
```

Blackf formats according to a format specifier and returns the resulting string with a black foreground.

<a name="Blue"></a>
## func [Blue](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L64>)

```go
func Blue(a ...any) string
```

Blue formats using the default formats for its operands and returns the resulting string with a blue foreground. Spaces are added between operands when neither is a string.

<a name="BlueBg"></a>
## func [BlueBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L64>)

```go
func BlueBg(a ...any) string
```

BlueBg formats using the default formats for its operands and returns the resulting string with a blue background. Spaces are added between operands when neither is a string.

<a name="BlueBgf"></a>
## func [BlueBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L70>)

```go
func BlueBgf(format string, a ...any) string
```

BlueBgf formats according to a format specifier and returns the resulting string with a blue background.

<a name="Bluef"></a>
## func [Bluef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L70>)

```go
func Bluef(format string, a ...any) string
```

Bluef formats according to a format specifier and returns the resulting string with a blue foreground.

<a name="Bold"></a>
## func [Bold](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L11>)

```go
func Bold(a ...any) string
```

Bold formats using the default formats for its operands and returns the resulting string in bold. Spaces are added between operands when neither is a string.

<a name="Boldf"></a>
## func [Boldf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L17>)

```go
func Boldf(format string, a ...any) string
```

Boldf formats according to a format specifier and returns the resulting string in bold.

<a name="Complete"></a>
## func [Complete](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L13>)

```go
func Complete(msg string) string
```

Complete returns a string with a green checkmark icon and the given message.

<a name="Completef"></a>
## func [Completef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L18>)

```go
func Completef(msg string, a ...any) string
```

Completef returns a string with a green checkmark icon and the given formatted message.

<a name="Cyan"></a>
## func [Cyan](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L90>)

```go
func Cyan(a ...any) string
```

Cyan formats using the default formats for its operands and returns the resulting string with a cyan foreground. Spaces are added between operands when neither is a string.

<a name="CyanBg"></a>
## func [CyanBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L90>)

```go
func CyanBg(a ...any) string
```

CyanBg formats using the default formats for its operands and returns the resulting string with a cyan background. Spaces are added between operands when neither is a string.

<a name="CyanBgf"></a>
## func [CyanBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L96>)

```go
func CyanBgf(format string, a ...any) string
```

CyanBgf formats according to a format specifier and returns the resulting string with a cyan background.

<a name="Cyanf"></a>
## func [Cyanf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L96>)

```go
func Cyanf(format string, a ...any) string
```

Cyanf formats according to a format specifier and returns the resulting string with a cyan foreground.

<a name="Dim"></a>
## func [Dim](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L23>)

```go
func Dim(a ...any) string
```

Dim formats using the default formats for its operands and returns the resulting string in dim. Spaces are added between operands when neither is a string.

<a name="Dimf"></a>
## func [Dimf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L29>)

```go
func Dimf(format string, a ...any) string
```

Dimf formats according to a format specifier and returns the resulting string in dim.

<a name="DoubleUnderline"></a>
## func [DoubleUnderline](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L86>)

```go
func DoubleUnderline(a ...any) string
```

DoubleUnderline formats using the default formats for its operands and returns the resulting string with a double underline. Spaces are added between operands when neither is a string.

<a name="DoubleUnderlinef"></a>
## func [DoubleUnderlinef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L92>)

```go
func DoubleUnderlinef(format string, a ...any) string
```

DoubleUnderlinef formats according to a format specifier and returns the resulting string with a double underline.

<a name="Fail"></a>
## func [Fail](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L43>)

```go
func Fail(msg string) string
```

Fail returns a string with a red X icon and the given message.

<a name="Failf"></a>
## func [Failf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L48>)

```go
func Failf(msg string, a ...any) string
```

Failf returns a string with a red X icon and the given formatted message.

<a name="Fg24Bit"></a>
## func [Fg24Bit](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L251>)

```go
func Fg24Bit(r int, g int, b int, a ...any) string
```

Fg24Bit formats using the default formats for its operands and returns the resulting string with an 24\-bit foreground color. Spaces are added between operands when neither is a string.

r, g, and b must be between 0 and 255 otherwise no color will be applied.

<a name="Fg24Bitf"></a>
## func [Fg24Bitf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L265>)

```go
func Fg24Bitf(r int, g int, b int, format string, a ...any) string
```

Fg24Bitf formats according to a format specifier and returns the resulting string with an 24\-bit foreground color.

r, g, and b must be between 0 and 255 otherwise no color will be applied.

<a name="Fg8Bit"></a>
## func [Fg8Bit](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L222>)

```go
func Fg8Bit(color int, a ...any) string
```

Fg8Bit formats using the default formats for its operands and returns the resulting string with an 8\-bit foreground color. Spaces are added between operands when neither is a string.

color must be between 0 and 255 otherwise no color will be applied.

<a name="Fg8Bitf"></a>
## func [Fg8Bitf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L236>)

```go
func Fg8Bitf(color int, format string, a ...any) string
```

Fg8Bitf formats according to a format specifier and returns the resulting string with an 8\-bit foreground color.

color must be between 0 and 255 otherwise no color will be applied.

<a name="Format"></a>
## func [Format](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L98>)

```go
func Format(msg string, a ...func(...any) string) string
```

Format applies the given formatting functions to the message in order and returns the resulting string. If no formatting functions are provided, the original message is returned.

<a name="Green"></a>
## func [Green](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L38>)

```go
func Green(a ...any) string
```

Green formats using the default formats for its operands and returns the resulting string with a green foreground. Spaces are added between operands when neither is a string.

<a name="GreenBg"></a>
## func [GreenBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L38>)

```go
func GreenBg(a ...any) string
```

GreenBg formats using the default formats for its operands and returns the resulting string with a green background. Spaces are added between operands when neither is a string.

<a name="GreenBgf"></a>
## func [GreenBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L44>)

```go
func GreenBgf(format string, a ...any) string
```

GreenBgf formats according to a format specifier and returns the resulting string with a green background.

<a name="Greenf"></a>
## func [Greenf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L44>)

```go
func Greenf(format string, a ...any) string
```

Greenf formats according to a format specifier and returns the resulting string with a green foreground.

<a name="Info"></a>
## func [Info](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L53>)

```go
func Info(msg string) string
```

Info returns a string with a cyan info icon and the given message.

<a name="Infof"></a>
## func [Infof](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L58>)

```go
func Infof(msg string, a ...any) string
```

Infof returns a string with a cyan info icon and the given formatted message.

<a name="IntenseBlack"></a>
## func [IntenseBlack](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L116>)

```go
func IntenseBlack(a ...any) string
```

IntenseBlack formats using the default formats for its operands and returns the resulting string with a intense black \(gray\) foreground. Spaces are added between operands when neither is a string.

<a name="IntenseBlackBg"></a>
## func [IntenseBlackBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L116>)

```go
func IntenseBlackBg(a ...any) string
```

IntenseBlackBg formats using the default formats for its operands and returns the resulting string with a intense black background. Spaces are added between operands when neither is a string.

<a name="IntenseBlackBgf"></a>
## func [IntenseBlackBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L122>)

```go
func IntenseBlackBgf(format string, a ...any) string
```

IntenseBlackBgf formats according to a format specifier and returns the resulting string with a intense black background.

<a name="IntenseBlackf"></a>
## func [IntenseBlackf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L122>)

```go
func IntenseBlackf(format string, a ...any) string
```

IntenseBlackf formats according to a format specifier and returns the resulting string with a intense black \(gray\) foreground.

<a name="IntenseBlue"></a>
## func [IntenseBlue](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L168>)

```go
func IntenseBlue(a ...any) string
```

IntenseBlue formats using the default formats for its operands and returns the resulting string with a intense blue foreground. Spaces are added between operands when neither is a string.

<a name="IntenseBlueBg"></a>
## func [IntenseBlueBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L168>)

```go
func IntenseBlueBg(a ...any) string
```

IntenseBlueBg formats using the default formats for its operands and returns the resulting string with a intense blue background. Spaces are added between operands when neither is a string.

<a name="IntenseBlueBgf"></a>
## func [IntenseBlueBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L174>)

```go
func IntenseBlueBgf(format string, a ...any) string
```

IntenseBlueBgf formats according to a format specifier and returns the resulting string with a intense blue background.

<a name="IntenseBluef"></a>
## func [IntenseBluef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L174>)

```go
func IntenseBluef(format string, a ...any) string
```

IntenseBluef formats according to a format specifier and returns the resulting string with a intense blue foreground.

<a name="IntenseCyan"></a>
## func [IntenseCyan](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L194>)

```go
func IntenseCyan(a ...any) string
```

IntenseCyan formats using the default formats for its operands and returns the resulting string with a intense cyan foreground. Spaces are added between operands when neither is a string.

<a name="IntenseCyanBg"></a>
## func [IntenseCyanBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L194>)

```go
func IntenseCyanBg(a ...any) string
```

IntenseCyanBg formats using the default formats for its operands and returns the resulting string with a intense cyan background. Spaces are added between operands when neither is a string.

<a name="IntenseCyanBgf"></a>
## func [IntenseCyanBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L200>)

```go
func IntenseCyanBgf(format string, a ...any) string
```

IntenseCyanBgf formats according to a format specifier and returns the resulting string with a intense cyan background.

<a name="IntenseCyanf"></a>
## func [IntenseCyanf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L200>)

```go
func IntenseCyanf(format string, a ...any) string
```

IntenseCyanf formats according to a format specifier and returns the resulting string with a intense cyan foreground.

<a name="IntenseGreen"></a>
## func [IntenseGreen](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L142>)

```go
func IntenseGreen(a ...any) string
```

IntenseGreen formats using the default formats for its operands and returns the resulting string with a intense green foreground. Spaces are added between operands when neither is a string.

<a name="IntenseGreenBg"></a>
## func [IntenseGreenBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L142>)

```go
func IntenseGreenBg(a ...any) string
```

IntenseGreenBg formats using the default formats for its operands and returns the resulting string with a intense green background. Spaces are added between operands when neither is a string.

<a name="IntenseGreenBgf"></a>
## func [IntenseGreenBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L148>)

```go
func IntenseGreenBgf(format string, a ...any) string
```

IntenseGreenBgf formats according to a format specifier and returns the resulting string with a intense green background.

<a name="IntenseGreenf"></a>
## func [IntenseGreenf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L148>)

```go
func IntenseGreenf(format string, a ...any) string
```

IntenseGreenf formats according to a format specifier and returns the resulting string with a intense green foreground.

<a name="IntenseMagenta"></a>
## func [IntenseMagenta](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L181>)

```go
func IntenseMagenta(a ...any) string
```

IntenseMagenta formats using the default formats for its operands and returns the resulting string with a intense magenta foreground. Spaces are added between operands when neither is a string.

<a name="IntenseMagentaBg"></a>
## func [IntenseMagentaBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L181>)

```go
func IntenseMagentaBg(a ...any) string
```

IntenseMagentaBg formats using the default formats for its operands and returns the resulting string with a intense magenta background. Spaces are added between operands when neither is a string.

<a name="IntenseMagentaBgf"></a>
## func [IntenseMagentaBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L187>)

```go
func IntenseMagentaBgf(format string, a ...any) string
```

IntenseMagentaBgf formats according to a format specifier and returns the resulting string with a intense magenta background.

<a name="IntenseMagentaf"></a>
## func [IntenseMagentaf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L187>)

```go
func IntenseMagentaf(format string, a ...any) string
```

IntenseMagentaf formats according to a format specifier and returns the resulting string with a intense magenta foreground.

<a name="IntenseRed"></a>
## func [IntenseRed](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L129>)

```go
func IntenseRed(a ...any) string
```

IntenseRed formats using the default formats for its operands and returns the resulting string with a intense red foreground. Spaces are added between operands when neither is a string.

<a name="IntenseRedBg"></a>
## func [IntenseRedBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L129>)

```go
func IntenseRedBg(a ...any) string
```

IntenseRedBg formats using the default formats for its operands and returns the resulting string with a intense red background. Spaces are added between operands when neither is a string.

<a name="IntenseRedBgf"></a>
## func [IntenseRedBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L135>)

```go
func IntenseRedBgf(format string, a ...any) string
```

IntenseRedBgf formats according to a format specifier and returns the resulting string with a intense red background.

<a name="IntenseRedf"></a>
## func [IntenseRedf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L135>)

```go
func IntenseRedf(format string, a ...any) string
```

IntenseRedf formats according to a format specifier and returns the resulting string with a intense red foreground.

<a name="IntenseWhite"></a>
## func [IntenseWhite](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L207>)

```go
func IntenseWhite(a ...any) string
```

IntenseWhite formats using the default formats for its operands and returns the resulting string with a intense white foreground. Spaces are added between operands when neither is a string.

<a name="IntenseWhiteBg"></a>
## func [IntenseWhiteBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L207>)

```go
func IntenseWhiteBg(a ...any) string
```

IntenseWhiteBg formats using the default formats for its operands and returns the resulting string with a intense white background. Spaces are added between operands when neither is a string.

<a name="IntenseWhiteBgf"></a>
## func [IntenseWhiteBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L213>)

```go
func IntenseWhiteBgf(format string, a ...any) string
```

IntenseWhiteBgf formats according to a format specifier and returns the resulting string with a intense white background.

<a name="IntenseWhitef"></a>
## func [IntenseWhitef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L213>)

```go
func IntenseWhitef(format string, a ...any) string
```

IntenseWhitef formats according to a format specifier and returns the resulting string with a intense white foreground.

<a name="IntenseYellow"></a>
## func [IntenseYellow](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L155>)

```go
func IntenseYellow(a ...any) string
```

IntenseYellow formats using the default formats for its operands and returns the resulting string with a intense yellow foreground. Spaces are added between operands when neither is a string.

<a name="IntenseYellowBg"></a>
## func [IntenseYellowBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L155>)

```go
func IntenseYellowBg(a ...any) string
```

IntenseYellowBg formats using the default formats for its operands and returns the resulting string with a intense yellow background. Spaces are added between operands when neither is a string.

<a name="IntenseYellowBgf"></a>
## func [IntenseYellowBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L161>)

```go
func IntenseYellowBgf(format string, a ...any) string
```

IntenseYellowBgf formats according to a format specifier and returns the resulting string with a intense yellow background.

<a name="IntenseYellowf"></a>
## func [IntenseYellowf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L161>)

```go
func IntenseYellowf(format string, a ...any) string
```

IntenseYellowf formats according to a format specifier and returns the resulting string with a intense yellow foreground.

<a name="Italic"></a>
## func [Italic](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L35>)

```go
func Italic(a ...any) string
```

Italic formats using the default formats for its operands and returns the resulting string in italic. Spaces are added between operands when neither is a string.

<a name="Italicf"></a>
## func [Italicf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L41>)

```go
func Italicf(format string, a ...any) string
```

Italicf formats according to a format specifier and returns the resulting string in italic.

<a name="Magenta"></a>
## func [Magenta](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L77>)

```go
func Magenta(a ...any) string
```

Magenta formats using the default formats for its operands and returns the resulting string with a magenta foreground. Spaces are added between operands when neither is a string.

<a name="MagentaBg"></a>
## func [MagentaBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L77>)

```go
func MagentaBg(a ...any) string
```

MagentaBg formats using the default formats for its operands and returns the resulting string with a magenta background. Spaces are added between operands when neither is a string.

<a name="MagentaBgf"></a>
## func [MagentaBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L83>)

```go
func MagentaBgf(format string, a ...any) string
```

MagentaBgf formats according to a format specifier and returns the resulting string with a magenta background.

<a name="Magentaf"></a>
## func [Magentaf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L83>)

```go
func Magentaf(format string, a ...any) string
```

Magentaf formats according to a format specifier and returns the resulting string with a magenta foreground.

<a name="Red"></a>
## func [Red](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L25>)

```go
func Red(a ...any) string
```

Red formats using the default formats for its operands and returns the resulting string with a red foreground. Spaces are added between operands when neither is a string.

<a name="RedAlert"></a>
## func [RedAlert](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L33>)

```go
func RedAlert(msg string) string
```

RedAlert returns a string with a red exclamation icon and the given message.

<a name="RedAlertf"></a>
## func [RedAlertf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L38>)

```go
func RedAlertf(msg string, a ...any) string
```

RedAlertf returns a string with a red exclamation icon and the given formatted message.

<a name="RedBg"></a>
## func [RedBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L25>)

```go
func RedBg(a ...any) string
```

RedBg formats using the default formats for its operands and returns the resulting string with a red background. Spaces are added between operands when neither is a string.

<a name="RedBgf"></a>
## func [RedBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L31>)

```go
func RedBgf(format string, a ...any) string
```

RedBgf formats according to a format specifier and returns the resulting string with a red background.

<a name="Redf"></a>
## func [Redf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L31>)

```go
func Redf(format string, a ...any) string
```

Redf formats according to a format specifier and returns the resulting string with a red foreground.

<a name="Reverse"></a>
## func [Reverse](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L60>)

```go
func Reverse(a ...any) string
```

Reverse formats using the default formats for its operands and returns the resulting string with reversed foreground and background. Spaces are added between operands when neither is a string.

<a name="Reversef"></a>
## func [Reversef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L66>)

```go
func Reversef(format string, a ...any) string
```

Reversef formats according to a format specifier and returns the resulting string with reversed foreground and background.

<a name="Strike"></a>
## func [Strike](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L73>)

```go
func Strike(a ...any) string
```

Strike formats using the default formats for its operands and returns the resulting string with a strikethrough. Spaces are added between operands when neither is a string.

<a name="Strikef"></a>
## func [Strikef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L79>)

```go
func Strikef(format string, a ...any) string
```

Strikef formats according to a format specifier and returns the resulting string with a strikethrough.

<a name="Underline"></a>
## func [Underline](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L47>)

```go
func Underline(a ...any) string
```

Underline formats using the default formats for its operands and returns the resulting string underlined. Spaces are added between operands when neither is a string.

<a name="Underlinef"></a>
## func [Underlinef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/formats.go#L53>)

```go
func Underlinef(format string, a ...any) string
```

Underlinef formats according to a format specifier and returns the resulting string underlined.

<a name="Var"></a>
## func [Var](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L63>)

```go
func Var(variable string, value string) string
```

Var returns a string with the given variable and value.

<a name="VarDoubleQuote"></a>
## func [VarDoubleQuote](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L73>)

```go
func VarDoubleQuote(variable string, value string) string
```

VarDoubleQuote returns a string with the given variable and value, double quoted.

<a name="VarQuote"></a>
## func [VarQuote](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L79>)

```go
func VarQuote(variable string, value string) string
```

VarQuote returns a string with the given variable and value, double quoted. Deprecated: use VarDoubleQuote instead.

<a name="VarSingleQuote"></a>
## func [VarSingleQuote](<https://github.com/engmtcdrm/go-prettyprint/blob/main/icons.go#L68>)

```go
func VarSingleQuote(variable string, value string) string
```

VarSingleQuote returns a string with the given variable and value, single quoted.

<a name="White"></a>
## func [White](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L103>)

```go
func White(a ...any) string
```

White formats using the default formats for its operands and returns the resulting string with a white foreground. Spaces are added between operands when neither is a string.

<a name="WhiteBg"></a>
## func [WhiteBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L103>)

```go
func WhiteBg(a ...any) string
```

WhiteBg formats using the default formats for its operands and returns the resulting string with a white background. Spaces are added between operands when neither is a string.

<a name="WhiteBgf"></a>
## func [WhiteBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L109>)

```go
func WhiteBgf(format string, a ...any) string
```

WhiteBgf formats according to a format specifier and returns the resulting string with a white background.

<a name="Whitef"></a>
## func [Whitef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L109>)

```go
func Whitef(format string, a ...any) string
```

Whitef formats according to a format specifier and returns the resulting string with a white foreground.

<a name="Yellow"></a>
## func [Yellow](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L51>)

```go
func Yellow(a ...any) string
```

Yellow formats using the default formats for its operands and returns the resulting string with a yellow foreground. Spaces are added between operands when neither is a string.

<a name="YellowBg"></a>
## func [YellowBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L51>)

```go
func YellowBg(a ...any) string
```

YellowBg formats using the default formats for its operands and returns the resulting string with a yellow background. Spaces are added between operands when neither is a string.

<a name="YellowBgf"></a>
## func [YellowBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/bg.go#L57>)

```go
func YellowBgf(format string, a ...any) string
```

YellowBgf formats according to a format specifier and returns the resulting string with a yellow background.

<a name="Yellowf"></a>
## func [Yellowf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/fg.go#L57>)

```go
func Yellowf(format string, a ...any) string
```

Yellowf formats according to a format specifier and returns the resulting string with a yellow foreground.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)

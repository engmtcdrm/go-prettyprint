## Description

The `prettyprint` package provides utilities for formatting strings with ANSI color codes. It includes predefined icons for different message types and functions to format strings with foreground and background colors.

Key Features:

- Foreground Color Formatting: Wraps a string in the ANSI color and ANSI reset codes, i.e. `Cyan` and `Cyanf`.
- Background Color Formatting: Wraps a string in the ANSI color and ANSI reset codes, i.e. `CyanBg` and `CyanBgf`.
- Predefined Icons: These icons can be used in front of text to provide better readability when writing to console or a log file.
- Helper functions for using predefined icons and color coding them to common colors, i.e. `Complete`, `Alert`, and `Fail`.

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

- [Variables](<#variables>)
- [func Alert\(msg string\) string](<#func-alert>)
- [func Black\(a ...any\) string](<#func-black>)
- [func BlackBg\(a ...any\) string](<#func-blackbg>)
- [func BlackBgf\(format string, a ...any\) string](<#func-blackbgf>)
- [func Blackf\(format string, a ...any\) string](<#func-blackf>)
- [func Blue\(a ...any\) string](<#func-blue>)
- [func BlueBg\(a ...any\) string](<#func-bluebg>)
- [func BlueBgf\(format string, a ...any\) string](<#func-bluebgf>)
- [func Bluef\(format string, a ...any\) string](<#func-bluef>)
- [func Complete\(msg string\) string](<#func-complete>)
- [func Cyan\(a ...any\) string](<#func-cyan>)
- [func CyanBg\(a ...any\) string](<#func-cyanbg>)
- [func CyanBgf\(format string, a ...any\) string](<#func-cyanbgf>)
- [func Cyanf\(format string, a ...any\) string](<#func-cyanf>)
- [func Fail\(msg string\) string](<#func-fail>)
- [func Green\(a ...any\) string](<#func-green>)
- [func GreenBg\(a ...any\) string](<#func-greenbg>)
- [func GreenBgf\(format string, a ...any\) string](<#func-greenbgf>)
- [func Greenf\(format string, a ...any\) string](<#func-greenf>)
- [func Info\(msg string\) string](<#func-info>)
- [func Magenta\(a ...any\) string](<#func-magenta>)
- [func MagentaBg\(a ...any\) string](<#func-magentabg>)
- [func MagentaBgf\(format string, a ...any\) string](<#func-magentabgf>)
- [func Magentaf\(format string, a ...any\) string](<#func-magentaf>)
- [func Red\(a ...any\) string](<#func-red>)
- [func RedAlert\(msg string\) string](<#func-redalert>)
- [func RedBg\(a ...any\) string](<#func-redbg>)
- [func RedBgf\(format string, a ...any\) string](<#func-redbgf>)
- [func Redf\(format string, a ...any\) string](<#func-redf>)
- [func Var\(variable string, value string\) string](<#func-var>)
- [func VarQuote\(variable string, value string\) string](<#func-varquote>)
- [func White\(a ...any\) string](<#func-white>)
- [func WhiteBg\(a ...any\) string](<#func-whitebg>)
- [func WhiteBgf\(format string, a ...any\) string](<#func-whitebgf>)
- [func Whitef\(format string, a ...any\) string](<#func-whitef>)
- [func Yellow\(a ...any\) string](<#func-yellow>)
- [func YellowBg\(a ...any\) string](<#func-yellowbg>)
- [func YellowBgf\(format string, a ...any\) string](<#func-yellowbgf>)
- [func Yellowf\(format string, a ...any\) string](<#func-yellowf>)

## Variables

```go
var (
    // IconComplete is the icon for a completed message.
    IconComplete = "[✓]"

    // IconAlert is the icon for an alerted message.
    IconAlert = "[!]"

    // IconFailed is the icon for a failed message.
    IconFailed = "[✗]"

    // IconInfo is the icon for an informational message.
    IconInfo = "[i]"
)
```

## func [Alert](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L237>)

```go
func Alert(msg string) string
```

Alert returns a string with a yellow exclamation icon and the given message.

## func [Black](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L26>)

```go
func Black(a ...any) string
```

Black formats using the default formats for its operands and returns the resulting string with a black foreground. Spaces are added between operands when neither is a string.

## func [BlackBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L39>)

```go
func BlackBg(a ...any) string
```

BlackBg formats using the default formats for its operands and returns the resulting string with a black background. Spaces are added between operands when neither is a string.

## func [BlackBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L45>)

```go
func BlackBgf(format string, a ...any) string
```

BlackBgf formats according to a format specifier and returns the resulting string with a black background.

## func [Blackf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L32>)

```go
func Blackf(format string, a ...any) string
```

Blackf formats according to a format specifier and returns the resulting string with a black foreground.

## func [Blue](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L130>)

```go
func Blue(a ...any) string
```

Blue formats using the default formats for its operands and returns the resulting string with a blue foreground. Spaces are added between operands when neither is a string.

## func [BlueBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L143>)

```go
func BlueBg(a ...any) string
```

BlueBg formats using the default formats for its operands and returns the resulting string with a blue background. Spaces are added between operands when neither is a string.

## func [BlueBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L149>)

```go
func BlueBgf(format string, a ...any) string
```

BlueBgf formats according to a format specifier and returns the resulting string with a blue background.

## func [Bluef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L136>)

```go
func Bluef(format string, a ...any) string
```

Bluef formats according to a format specifier and returns the resulting string with a blue foreground.

## func [Complete](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L232>)

```go
func Complete(msg string) string
```

Complete returns a string with a green checkmark icon and the given message.

## func [Cyan](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L182>)

```go
func Cyan(a ...any) string
```

Cyan formats using the default formats for its operands and returns the resulting string with a cyan foreground. Spaces are added between operands when neither is a string.

## func [CyanBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L195>)

```go
func CyanBg(a ...any) string
```

CyanBg formats using the default formats for its operands and returns the resulting string with a cyan background. Spaces are added between operands when neither is a string.

## func [CyanBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L201>)

```go
func CyanBgf(format string, a ...any) string
```

CyanBgf formats according to a format specifier and returns the resulting string with a cyan background.

## func [Cyanf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L188>)

```go
func Cyanf(format string, a ...any) string
```

Cyanf formats according to a format specifier and returns the resulting string with a cyan foreground.

## func [Fail](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L247>)

```go
func Fail(msg string) string
```

Fail returns a string with a red X icon and the given message.

## func [Green](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L78>)

```go
func Green(a ...any) string
```

Green formats using the default formats for its operands and returns the resulting string with a green foreground. Spaces are added between operands when neither is a string.

## func [GreenBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L91>)

```go
func GreenBg(a ...any) string
```

GreenBg formats using the default formats for its operands and returns the resulting string with a green background. Spaces are added between operands when neither is a string.

## func [GreenBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L97>)

```go
func GreenBgf(format string, a ...any) string
```

GreenBgf formats according to a format specifier and returns the resulting string with a green background.

## func [Greenf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L84>)

```go
func Greenf(format string, a ...any) string
```

Greenf formats according to a format specifier and returns the resulting string with a green foreground.

## func [Info](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L252>)

```go
func Info(msg string) string
```

Info returns a string with a cyan info icon and the given message.

## func [Magenta](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L156>)

```go
func Magenta(a ...any) string
```

Magenta formats using the default formats for its operands and returns the resulting string with a magenta foreground. Spaces are added between operands when neither is a string.

## func [MagentaBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L169>)

```go
func MagentaBg(a ...any) string
```

MagentaBg formats using the default formats for its operands and returns the resulting string with a magenta background. Spaces are added between operands when neither is a string.

## func [MagentaBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L175>)

```go
func MagentaBgf(format string, a ...any) string
```

MagentaBgf formats according to a format specifier and returns the resulting string with a magenta background.

## func [Magentaf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L162>)

```go
func Magentaf(format string, a ...any) string
```

Magentaf formats according to a format specifier and returns the resulting string with a magenta foreground.

## func [Red](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L52>)

```go
func Red(a ...any) string
```

Red formats using the default formats for its operands and returns the resulting string with a red foreground. Spaces are added between operands when neither is a string.

## func [RedAlert](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L242>)

```go
func RedAlert(msg string) string
```

RedAlert returns a string with a red exclamation icon and the given message.

## func [RedBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L65>)

```go
func RedBg(a ...any) string
```

RedBg formats using the default formats for its operands and returns the resulting string with a red background. Spaces are added between operands when neither is a string.

## func [RedBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L71>)

```go
func RedBgf(format string, a ...any) string
```

RedBgf formats according to a format specifier and returns the resulting string with a red background.

## func [Redf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L58>)

```go
func Redf(format string, a ...any) string
```

Redf formats according to a format specifier and returns the resulting string with a red foreground.

## func [Var](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L257>)

```go
func Var(variable string, value string) string
```

Var returns a string with the given variable and value.

## func [VarQuote](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L262>)

```go
func VarQuote(variable string, value string) string
```

VarQuote returns a string with the given variable and value, quoted.

## func [White](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L208>)

```go
func White(a ...any) string
```

White formats using the default formats for its operands and returns the resulting string with a white foreground. Spaces are added between operands when neither is a string.

## func [WhiteBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L221>)

```go
func WhiteBg(a ...any) string
```

WhiteBg formats using the default formats for its operands and returns the resulting string with a white background. Spaces are added between operands when neither is a string.

## func [WhiteBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L227>)

```go
func WhiteBgf(format string, a ...any) string
```

WhiteBgf formats according to a format specifier and returns the resulting string with a white background.

## func [Whitef](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L214>)

```go
func Whitef(format string, a ...any) string
```

Whitef formats according to a format specifier and returns the resulting string with a white foreground.

## func [Yellow](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L104>)

```go
func Yellow(a ...any) string
```

Yellow formats using the default formats for its operands and returns the resulting string with a yellow foreground. Spaces are added between operands when neither is a string.

## func [YellowBg](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L117>)

```go
func YellowBg(a ...any) string
```

YellowBg formats using the default formats for its operands and returns the resulting string with a yellow background. Spaces are added between operands when neither is a string.

## func [YellowBgf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L123>)

```go
func YellowBgf(format string, a ...any) string
```

YellowBgf formats according to a format specifier and returns the resulting string with a yellow background.

## func [Yellowf](<https://github.com/engmtcdrm/go-prettyprint/blob/main/prettyprint.go#L110>)

```go
func Yellowf(format string, a ...any) string
```

Yellowf formats according to a format specifier and returns the resulting string with a yellow foreground.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)

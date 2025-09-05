# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v1.2.0] - 2025-09-04

### Added

#### Intense/Bright Color Support

- Added `IntenseBlack` and `IntenseBlackf` for intense black (gray) foreground
- Added `IntenseRed` and `IntenseRedf` for intense red foreground
- Added `IntenseGreen` and `IntenseGreenf` for intense green foreground
- Added `IntenseYellow` and `IntenseYellowf` for intense yellow foreground
- Added `IntenseBlue` and `IntenseBluef` for intense blue foreground
- Added `IntenseMagenta` and `IntenseMagentaf` for intense magenta foreground
- Added `IntenseCyan` and `IntenseCyanf` for intense cyan foreground
- Added `IntenseWhite` and `IntenseWhitef` for intense white foreground
- Added `IntenseBlackBg` and `IntenseBlackBgf` for intense black background
- Added `IntenseRedBg` and `IntenseRedBgf` for intense red background
- Added `IntenseGreenBg` and `IntenseGreenBgf` for intense green background
- Added `IntenseYellowBg` and `IntenseYellowBgf` for intense yellow background
- Added `IntenseBlueBg` and `IntenseBlueBgf` for intense blue background
- Added `IntenseMagentaBg` and `IntenseMagentaBgf` for intense magenta background
- Added `IntenseCyanBg` and `IntenseCyanBgf` for intense cyan background
- Added `IntenseWhiteBg` and `IntenseWhiteBgf` for intense white background

#### 8-bit and 24-bit Color Support

- Added `Fg8Bit` and `Fg8Bitf` for 8-bit foreground colors (0-255 range)
- Added `Bg8Bit` and `Bg8Bitf` for 8-bit background colors (0-255 range)
- Added `Fg24Bit` and `Fg24Bitf` for 24-bit RGB foreground colors
- Added `Bg24Bit` and `Bg24Bitf` for 24-bit RGB background colors

#### Text Formatting Functions

- Added `Bold` and `Boldf` for bold text formatting
- Added `Dim` and `Dimf` for dim/faint text formatting
- Added `Italic` and `Italicf` for italic text formatting
- Added `Underline` and `Underlinef` for underlined text formatting
- Added `DoubleUnderline` and `DoubleUnderlinef` for double underlined text formatting
- Added `Strike` and `Strikef` for strikethrough text formatting
- Added `Reverse` and `Reversef` for reversed foreground/background colors
- Added `Format` function to apply multiple formatting functions in sequence

#### Icon and Message Functions

- Added `Completef` for formatted completion messages with green checkmark
- Added `Alertf` for formatted alert messages with yellow exclamation
- Added `RedAlertf` for formatted red alert messages with red exclamation
- Added `Failf` for formatted failure messages with red X
- Added `Infof` for formatted info messages with cyan info icon
- Added `VarSingleQuote` for displaying variables with single quotes
- Added `VarDoubleQuote` for displaying variables with double quotes

### Deprecated

- Marked `VarQuote` as deprecated in favor of `VarDoubleQuote` for clarity

## [v1.1.0] - 2025-01-03

### Added

- Added funcs `BlackBg` and `BlackBgf`
- Added funcs `RedBg` and `RedBgf`
- Added funcs `GreenBg` and `GreenBgf`
- Added funcs `YellowBg` and `YellowBgf`
- Added funcs `BlueBg` and `BlueBgf`
- Added funcs `MagentaBg` and `MagentaBgf`
- Added funcs `CyanBg` and `CyanBgf`
- Added funcs `WhiteBg` and `WhiteBgf`

## [v1.0.0] - 2024-12-30

Initial release testing versioning in go. This version is identical to v0.1.0 and was removed.

## [v0.1.0] - 2024-12-30

Initial release

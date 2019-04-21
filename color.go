package color

import (
	"fmt"
	"strings"
)

// ColorStubStart is the text string stub which the shell
// can interpret as an ANSI color
const ColorStubStart = "\033["

// ColorStubEnd is the text string stub which the shell
// uses to end ColorStubstart
const ColorStubEnd = "m"

// Palette stores the color codes
var Palette = map[string]string{
	"default":   "0",
	"bold":      "1",
	"dim":       "2",
	"italics":   "3",
	"underline": "4",
	"black":     "30",
	"white":     "97",
	"gray":      "90",
	"grey":      "90",
	"red":       "31",
	"yellow":    "33",
	"green":     "32",
	"cyan":      "36",
	"blue":      "34",
	"violet":    "35",
	"bgwhite":   "47",
	"bgblack":   "40",
	"bgred":     "41",
	"bgyellow":  "43",
	"bggreen":   "42",
	"bgcyan":    "46",
	"bgblue":    "44",
	"bgviolet":  "45",
	"lgray":     "37",
	"lgrey":     "37",
	"lred":      "91",
	"lyellow":   "93",
	"lgreen":    "92",
	"lcyan":     "96",
	"lblue":     "94",
	"lviolet":   "95",
}

// Color applies the color code :color to the string :value
func Color(color string, value string) string {
	format := ColorStubStart + Palette[color] + ColorStubEnd
	unformat := ColorStubStart + "0" + ColorStubEnd
	finalValue := value
	if strings.Contains(value, unformat) {
		finalValue = strings.Replace(value, unformat, unformat+format, -1)
	}
	return fmt.Sprintf("%s%s%s", format, finalValue, unformat)
}

// Default sets the string :value to the default color
func Default(value string) string {
	return Color("default", value)
}

// Bold sets the string :value to bold
func Bold(value string) string {
	return Color("bold", value)
}

// Dim dims the string :value
func Dim(value string) string {
	return Color("dim", value)
}

// Italics italicises the string :value
func Italics(value string) string {
	return Color("italics", value)
}

// Underline underlines the string :value
func Underline(value string) string {
	return Color("underline", value)
}

// Black sets the string :value to black
func Black(value string) string {
	return Color("black", value)
}

// Gray sets the string :value to gray
func Gray(value string) string {
	return Color("gray", value)
}

// Grey sets the string :value to grey
func Grey(value string) string {
	return Color("grey", value)
}

// Red sets the string :value to red
func Red(value string) string {
	return Color("red", value)
}

// LightRed sets the string :value to light red
func LightRed(value string) string {
	return Color("lred", value)
}

// Green sets the string :value to green
func Green(value string) string {
	return Color("green", value)
}

// LightGreen sets the string :value to light green
func LightGreen(value string) string {
	return Color("lgreen", value)
}

// Yellow sets the string :value to yellow
func Yellow(value string) string {
	return Color("yellow", value)
}

// LightYellow sets the string :value to light yellow
func LightYellow(value string) string {
	return Color("lyellow", value)
}

// Blue sets the string :value to blue
func Blue(value string) string {
	return Color("blue", value)
}

// LightBlue sets the string :value to light blue
func LightBlue(value string) string {
	return Color("lblue", value)
}

// Violet sets the string :value to violet
func Violet(value string) string {
	return Color("violet", value)
}

// LightViolet sets the string :value to light violet
func LightViolet(value string) string {
	return Color("lviolet", value)
}

// Cyan sets the string :value to cyan
func Cyan(value string) string {
	return Color("cyan", value)
}

// LightCyan sets the string :value to light cyan
func LightCyan(value string) string {
	return Color("lcyan", value)
}

// LightGray sets the string :value to light gray
func LightGray(value string) string {
	return Color("lgray", value)
}

// LightGrey sets the string :value to light grey
func LightGrey(value string) string {
	return Color("lgrey", value)
}

// White sets the string :value to light white
func White(value string) string {
	return Color("white", value)
}

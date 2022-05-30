// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package print

import (
	"fmt"
	"strings"
)

type (
	Print struct {
		Color string
	}
)

const (
	off    = "\x1b[0m"    // Text Reset
	black  = "\x1b[1;30m" // Black
	red    = "\x1b[1;31m" // Red
	green  = "\x1b[1;32m" // Green
	yellow = "\x1b[1;33m" // Yellow
	blue   = "\x1b[1;34m" // Blue
	purple = "\x1b[1;35m" // Purple
	cyan   = "\x1b[1;36m" // Cyan
	White  = "\x1b[1;37m" // White

	redBase    = "\x1b[0;31m" // Red no highlighted
	greenbase  = "\x1b[0;32m" // Green no highlighted
	yellowBase = "\x1b[0;33m" // Yellow no highlighted
	blueBase   = "\x1b[0;34m" // Blue no highlighted
	purpleBase = "\x1b[0;35m" // Purple no highlighted
	cyanBase   = "\x1b[0;36m" // Cyan no highlighted
	WhiteBase  = "\x1b[0;37m" // White no highlighted

	redUnderline = "\x1b[4;31m" // Red underline
	OneLineUP    = "\x1b[A"

	ClearLine   = "\x1b[0G\x1b[2K\x1b[0m\r"
	ClearScreen = "\x1b[H\x1b[2J"
	HEADER      = "---------------"
	LINE        = "_________________________________________________"
)

var (
	dangerZone = fmt.Sprintf("%sDanger Zone%s, be sure you understand the implication!",
		redUnderline, off)
)

func New() *Print {
	return &Print{}
}

// functions to print messsage in the given color
func (p *Print) PrintRed(message string) {
	fmt.Printf("%s%s%s", red, message, off)
}

func (p *Print) PrintGreen(message string) {
	fmt.Printf("%s%s%s", green, message, off)
}

func (p *Print) PrintYellow(message string) {
	fmt.Printf("%s%s%s", yellow, message, off)
}

func (p *Print) PrintBlue(message string) {
	fmt.Printf("%s%s%s", blue, message, off)
}

func (p *Print) PrintPurple(message string) {
	fmt.Printf("%s%s%s", purple, message, off)
}

func (p *Print) PrintCyan(message string) {
	fmt.Printf("%s%s%s", cyan, message, off)
}

func (p *Print) MessageRed(message string) string {
	return fmt.Sprintf("%s%s%s", red, message, off)
}

func (p *Print) MessageGreen(message string) string {
	return fmt.Sprintf("%s%s%s", green, message, off)
}

func (p *Print) MessageYellow(message string) string {
	return fmt.Sprintf("%s%s%s", yellow, message, off)
}

func (p *Print) MessageBlue(message string) string {
	return fmt.Sprintf("%s%s%s", blue, message, off)
}

func (p *Print) MessagePurple(message string) string {
	return fmt.Sprintf("%s%s%s", purple, message, off)
}

func (p *Print) MessageCyan(message string) string {
	return fmt.Sprintf("%s%s%s", cyan, message, off)
}

// function Message line
func (p *Print) PrintLine(lineColor string, count int) string {
	line := strings.Repeat("⎻", count)
	return fmt.Sprintf("%s%s%s", lineColor, line, off)
}

// function to clear the screen
func (p *Print) ClearScreen() string {
	return fmt.Sprintf("%s", ClearScreen)
}

// function to move one line up
func (p *Print) OneLineUP(clearTheLine bool) string {
	out := fmt.Sprintf("%s", OneLineUP)
	if clearTheLine {
		out = out + fmt.Sprintf("%s", ClearLine)
	}
	return out
}

func (p *Print) PrintDangerZone() string {
	return fmt.Sprintf("%s", dangerZone)
}

func (p *Print) PrintColorMessage(messageColor, message string) string {
	return fmt.Sprintf("%s%s%s", messageColor, message, off)
}

// The End Message
func (p *Print) TheEnd() {
	p.PrintGreen("\tEnjoy a cuppa of hot coffee ☕️ / 🥃 \n")
	p.PrintGreen("\tThe End\n")
}

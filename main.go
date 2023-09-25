package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var termWidth int

func init() {
	termWidth = getTerminalWidth()
}

func main() {
	if len(os.Args) < 2 || !isValidFormat(os.Args) {
		printUsage()
		return
	}

	font := "standard.txt"
	align := "left"
	text := ""

	// Parse command-line arguments
	for _, v := range os.Args[1:] {
		if strings.HasPrefix(v, "--align=") {
			align = v[8:]
		} else if v == "standard" || v == "shadow" || v == "thinkertoy" {
			font = v + ".txt"
		} else {
			text += v + " "
		}
	}

	text = strings.TrimSpace(text)

	switch align {
	case "left", "right", "center", "centre":
		renderText(text, font, align)
	case "justify":
		justifyText(text, font)
	default:
		printUsage()
	}
}

func isValidFormat(args []string) bool {
	if len(args) < 2 {
		return false
	}

	// Check if the first argument is "--align=<type>"
	if !strings.HasPrefix(args[1], "--align=") {
		return false
	}

	// Check if there is at least one more argument (STRING or BANNER)
	if len(args) < 3 {
		return false
	}

	return true
}

func renderText(text, font, align string) {
	res := ""
	args := strings.Split(text, "\\n")
	for _, word := range args {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				res += getLine(1+int(letter-' ')*9+i, font)
			}
			printAlignedText(res, align)
			res = ""
		}
	}
}

func justifyText(words, font string) {
	lines := strings.Split(words, "\\n")
	for _, line := range lines {
		words := strings.Fields(line) // Split the line into words
		totalWidth := 0
		for _, word := range words {
			totalWidth += len(word)
		}

		numWords := len(words)
		if numWords <= 1 {
			// Nothing to justify if there's only one word or none
			fmt.Println(line)
			continue
		}

		// Calculate the number of spaces to distribute
		extraSpaces := termWidth - totalWidth
		spacesBetweenWords := extraSpaces / (numWords - 1)
		extraSpaces %= numWords - 1 // Distribute any remaining spaces

		// Build the justified line
		justifiedLine := ""
		for i, word := range words {
			justifiedLine += word
			if i < numWords-1 {
				// Add spaces between words
				justifiedLine += printSpaces(spacesBetweenWords)
				if i < extraSpaces {
					// Add an extra space for remaining spaces
					justifiedLine += " "
				}
			}
		}

		// Print the justified line
		fmt.Println(justifiedLine)
	}
}

func getTerminalWidth() int {
	out, err := exec.Command("tput", "cols").Output()
	if err != nil {
		log.Fatal(err)
	}
	width, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		log.Fatal(err)
	}
	return width
}

func getLine(num int, filename string) string {
	f, e := os.Open("./fonts/" + filename)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()

	content := bufio.NewScanner(f)
	lineNum := 0
	line := ""
	for content.Scan() {
		if lineNum == num {
			line = content.Text()
		}
		lineNum++
	}
	return line
}

func printAlignedText(text, align string) {
	padding := ""
	if align == "right" {
		padding = printSpaces(termWidth - len(text))
	} else if align == "center" || align == "centre" {
		padding = printSpaces((termWidth - len(text)) / 2)
	}
	fmt.Println(padding + text)
}

func printSpaces(num int) string {
	a := ""
	for i := 1; i <= num; i++ {
		a += " "
	}
	return a
}

func printUsage() {
	fmt.Println(`Usage: go run . [OPTIONS] [TEXT]

OPTIONS:
  --align=<type>  Specify the text alignment (left, right, center, justify) (default "left")
  [TEXT]          Text to be rendered as ASCII art

Example: go run . --align=right standard`)
}

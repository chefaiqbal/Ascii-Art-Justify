package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unicode"
	"unsafe"
)

func main() {
	var alignment, bannerType string
	flag.StringVar(&alignment, "align", "left", "text alignment (left, center, right, justify)")
	flag.StringVar(&bannerType, "type", "standard", "banner type (standard, shadow, tinkertoy)")

	err := "\033[47mUsage: go run . [OPTION] [STRING] [BANNER] \n\nExample: go run . --align=right something standard\n\033[0m"

	if len(os.Args) < 2 {
		fmt.Println(err)
		os.Exit(0)
	}

	if string(os.Args[1]) == "--align" || strings.HasPrefix(string(os.Args[1]), "-align") {
		fmt.Println(err)
		os.Exit(0)
	} else {
		flag.Parse()
	}

	args := flag.Args()
	userInput := args[0]
	terminalWidth := terminalWidth()

	if len(args) > 1 && (args[1] == "shadow" || args[1] == "tinkertoy") {
		bannerType = args[1]
	}
	ascii := mapFont(bannerType)

	isASCII := func(s string) bool {
		for _, c := range s {
			if c > unicode.MaxASCII {
				return false
			}
		}
		return true
	}

	if !isASCII(userInput) {
		fmt.Println("Error: input string must be within the range of ASCII characters.")
		os.Exit(0)
	}

	words := []string{userInput}
	if strings.Contains(words[0], "\\n") {
		words = strings.Split(words[0], "\\n")
	}

	if alignment != "left" && alignment != "center" && alignment != "right" && alignment != "justify" {
		fmt.Println("Error: alignment must be left, center, right, or justify.")
		os.Exit(0)
	}

	alignment = strings.ToLower(alignment)
	printOutput(words, ascii, terminalWidth, alignment)
}

func mapFont(fileName string) map[rune][]string {
	file, err := os.Open("banner/" + fileName + ".txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	asciiArr := parseFile(file)

	var asciiStart rune = 32
	ascii := make(map[rune][]string)
	for i, char := range asciiArr {
		ascii[rune(i+int(asciiStart))] = char
	}

	return ascii
}

func parseFile(file *os.File) [][]string {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var asciiChar []string
	var asciiArr [][]string

	counter := 0
	for fileScanner.Scan() {
		if counter == 8 {
			asciiChar = append(asciiChar, fileScanner.Text())
			asciiArr = append(asciiArr, asciiChar)
			asciiChar = []string{}
			counter = 0
			continue
		}
		counter++
		asciiChar = append(asciiChar, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}

	return asciiArr
}

func terminalWidth() int {
	var dimensions [2]uint16

	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions))); err != 0 {
		fmt.Printf("error getting terminal size: %v\n", err)
	}

	return int(dimensions[1])
}

func printOutput(words []string, ascii map[rune][]string, terminalWidth int, align string) {
	var alignment string
	wordsPerLine := 0

	for index, word := range words {
		wordLength := 0

		for _, runes := range word {
			if runes == ' ' && align == "justify" {
				wordsPerLine++
			}
			wordLength = wordLength + len(ascii[runes][4])
		}

		if wordLength > terminalWidth {
			fmt.Println("Words don't fit in terminal.")
			os.Exit(0)
		}

		switch align {
		case "center":
			alignment = strings.Repeat(" ", (terminalWidth-wordLength)/2)
		case "right":
			alignment = strings.Repeat(" ", terminalWidth-wordLength)
		case "justify":
			if wordsPerLine == 0 {
				align = "none"
			} else {
				alignment = strings.Repeat(" ", (terminalWidth-wordLength)/wordsPerLine)
			}
		}

		for i := 0; i <= 8; i++ {
			for j, runes := range word {
				if j == 0 && align != "justify" {
					fmt.Print(alignment)
				}
				if align == "justify" && runes == ' ' {
					fmt.Print(alignment)
				}
				fmt.Print(ascii[runes][i])
			}
			if i == 8 && index != len(words)-1 {
				continue
			}
			fmt.Println()
		}
		wordsPerLine = 0
	}
}

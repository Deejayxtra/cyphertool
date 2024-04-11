package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Underline = "\033[4m"
	Reset     = "\033[0m"
	Green     = "\033[32m"
	Blue      = "\033[1;34m"
	bold      = "\033[1m"
)

func main() {

	if len(os.Args) == 1 || os.Args[1] == "" {
		fmt.Println("Error")
		os.Exit(1)
	}
	if os.Args[1] == "HELP" {
		fmt.Println(bold + "*****************************************************************************************" + Reset)
		fmt.Println("\n" + Blue + Underline + bold + "Four usage options:\n" + Reset)
		fmt.Println(Underline + bold + "1. Decoding a single line:" + Reset + Green + " <$ go run . 'encoded text'>\n" + Reset)
		fmt.Println(Underline + bold + "2. Rendering a .txt resource:" + Reset + Green + " <$ go run . render filetorender.txt>" + Reset)
		fmt.Println("  -Rendering a .txt resource also creates Masterpiece?.txt which contains the output.\n")
		fmt.Println(Underline + bold + "3. Encoding a single line:" + Reset + Green + " <$ go run . encode 'text to encode'>\n" + Reset)
		fmt.Println(Underline + bold + "4. Encoding a .txt resource:" + Reset + Green + " <$ go run . encodeart filetoencode.txt>" + Reset)
		fmt.Println("  -Encoding a .txt resource also creates Masterpiece?.txt which contains the output.\n")
		fmt.Println(bold + "*****************************************************************************************" + Reset)
	}
	if os.Args[1] == "encodeart" {
		artPath := os.Args[2]
		codedArt := inputRead(artPath)
		timed(encode(codedArt))
		file, _ := os.Create("Masterpiece?.txt")
		file.WriteString(encode(codedArt))
		defer file.Close()
	}
	if os.Args[1] == "encode" {
		artPath := os.Args[2]
		codedArt := encode(artPath)
		timed(codedArt)
		fmt.Println()
	}
	if os.Args[1] == "render" {
		artPath := os.Args[2]
		inputArt := inputRead(artPath)
		if !ifValid(inputArt) {
			fmt.Println("Error\n")
		} else {
			timed(artist(inputArt))
			file, _ := os.Create("Masterpiece?.txt")
			file.WriteString(artist(inputArt))
			fmt.Println()
			defer file.Close()

		}
	}
	if len(os.Args) == 2 && os.Args[1] != "HELP" {
		inputArt := os.Args[1]
		if !ifValid(inputArt) {
			fmt.Println("Error\n")
		} else {
			timed(artist(inputArt) + "\n")

		}
	}
}

func timed(text string) { //creates some extra flare for the output

	for _, char := range text {
		fmt.Printf(Blue+"%c"+Reset, char)
		time.Sleep(3 * time.Millisecond)
	}
}

func inputRead(content string) string { //reads the .txt file when rendering .txt resources
	data, err := os.ReadFile(content)
	if os.IsNotExist(err) {
		fmt.Println("Input not found")
		os.Exit(1)
	}
	return string(data)
}

func artist(input string) string { //decodes the encoded text

	index := 0
	result := ""
	code := ""

	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			result += input[index:i]
			index = i
		}
		if input[i] == ']' {
			code = input[index+1 : i]
			parts := strings.SplitN(code, " ", 2)
			mplr, err := strconv.Atoi(parts[0])
			result += multiplier(mplr, parts[1])
			if err != nil {
				return "conversion error"
			}
			if i != len(input)-1 {
				index = i + 1
			}
		}
		if i == len(input)-1 && input[i] != ']' {
			result += input[index : i+1]
		}
	}
	return result
}

func multiplier(n int, s string) string { //artist uses this to duplicate symbols according to the code numbers
	if n <= 0 {
		return ""
	}
	return s + multiplier(n-1, s)
}

func ifValid(input string) bool { //checks the validity of the encoded text, returns error if the code is invalid

	open := false
	expectSpace := false
	firstnumber := false
	expectNumber := false

	for i := 0; i < len(input); i++ {
		if input[i] == ']' {
			if !open {
				return false
			} else {
				open = false
			}
		}
		if input[i] == '[' {
			if open {
				return false
			} else {
				open = true
				expectSpace = true
				firstnumber = true
				expectNumber = true
				i++
			}
		}
		if expectSpace {
			if input[i] == ' ' {
				if expectNumber {
					return false
				} else {
					expectSpace = false
				}
			}
			if input[i] != ' ' {
				if input[i] >= '0' && input[i] <= '9' {
					expectNumber = false
					if input[i] == '0' {
						if firstnumber {
							return false
						}
					} else {
						firstnumber = false
					}
				} else {
					return false
				}

			}
		}
	}
	if !open {
		return true
	} else {
		return false
	}
}

func encode(input string) string { // converts the single and multiline art to encoded text
	if len(input) <= 2 {
		return input
	}
	result := ""
	singlecount := 1
	multicount := 1
	var letter byte
	var subarr string

	for i, j := 0, 1; j < len(input); i, j = i+1, j+1 {
		if input[i] == input[j] {
			singlecount++
			letter = input[i]
			if subarr != "" {
				result += string(subarr)
				subarr = ""
			}
			if j >= len(input)-1 {
				formatedString := ""
				formatedString = fmt.Sprintf("[%d %s]", singlecount, string(letter))
				result += formatedString
			}
		}
		if input[i] != input[j] {
			if singlecount > 1 {
				formatedString := ""
				formatedString = fmt.Sprintf("[%d %s]", singlecount, string(letter))
				result += formatedString
				singlecount = 1
				letter = 0
			} else {
				if len(subarr) >= 2 {
					if subarr == string(input[i])+string(input[j]) {
						multicount++
						if j >= len(input)-1 {
							formatedString := ""
							formatedString = fmt.Sprintf("[%d %s]", multicount, string(subarr))
							result += formatedString
						} else {
							if len(input) >= j+len(subarr) {
								i = i + 1
								j = j + 1
							}
						}
					} else {
						if multicount > 1 {
							formatedString := ""
							formatedString = fmt.Sprintf("[%d %s]", multicount, string(subarr))
							result += formatedString
							multicount = 1
							subarr = string(input[i])
						} else {
							result += string(input[i-2])
							subarr = string(input[i-1]) + string(input[i])
						}
					}
				} else {
					subarr += string(input[i])
				}
			}
		}
		if j >= len(input)-1 && input[i] != input[j] {
			if multicount > 1 {
				formatedString := ""
				formatedString = fmt.Sprintf("[%d %s]%s", multicount, string(subarr), string(input[j]))
				result += formatedString
			} else {
				result += string(subarr) + string(input[j])
			}
		}
	}
	return result
}

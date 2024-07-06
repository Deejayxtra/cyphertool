package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

// main logic, invoking other functions
func main() {
	operation, cypher, message := input()

	var encryptedMessage string
	switch cypher{
	case "ROT13":
		encryptedMessage = encryptShift(message, 13)
	case "reverse":
		encryptedMessage = encryptReverse(message)
	case "custom":
		if operation == "Encrypted" {
		encryptedMessage = encryptCustom(message)
		} else {
			encryptedMessage = decryptCustom(message)
		}
	default:
		fmt.Println("Invalid encoding")
		return
	}

	fmt.Printf("%s message using %s:\n%s\n", operation, cypher, encryptedMessage)
}


// Get the input data required for the operation
func input() (operation, encoding, message string) { 

	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println()
	fmt.Println("Select operation (1/2):")
	fmt.Println("1. Encrypt.")
	fmt.Println("2. Decrypt.")
	fmt.Println()

	// get the desired operation
	// continue asking for input if invalid input is entered
	var op int
	OP:for true {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&op)
		switch op {
		case 1:
			operation = "Encrypted"
			break OP
		case 2:
			operation = "Decrypted"
			break OP
		default:
			fmt.Println("\nInvalid operation")
		}
	}

	fmt.Println()
	fmt.Println("Select cypher (1/3):")
	fmt.Println("1. ROT13.")
	fmt.Println("2. Reverse.")
	fmt.Println("3. Custom.")
	fmt.Println()

	// get the desired encryption method.
	// continue asking for input if invalid input is entered
	var cypher int
	CY:for true {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&cypher)
		switch cypher {
		case 1:
			encoding = "ROT13"
			break CY
		case 2:
			encoding = "reverse"
			break CY
		case 3:
			encoding = "custom"
			break CY
		default:
			fmt.Println("\nInvalid cypher")
		}
	}
		
	fmt.Println()
	fmt.Print("Enter the message: ")
	reader := bufio.NewReader(os.Stdin)
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message)

	return operation, encoding, message
}

// Encrypt the message with shift
func encryptShift(s string, step int) string {
	result := ""
	for _, r := range s {
		shifted := shiftBy(r, step)
		result += string(shifted)
	}
	return result
}

// Encrypt the message with reverse
func encryptReverse(s string) string {
	result := ""
	for _, r := range s {
		reversed := reverseAlphabetValue(r)
		result += string(reversed)
	}
	return result
}

// Encrypt the message with custom encryption
func encryptCustom(s string) string {
	result := ""
	for _, r := range s {
		encrypted := customEncryption(r)
		result += string(encrypted)
	}
	return result
}

func decryptCustom(s string) string {
	result := ""
	for _, r := range s {
		encrypted := customDecryption(r)
		result += string(encrypted)
	}
	return result
}

// Decrypt the message with reverse
func decryptReverse(s string) string {
	result := ""
	for _, r := range s {
		decrypted := reverseAlphabetValue(r)
		result += string(decrypted)
	}
	return result
}

// ShiftBy function
func shiftBy(r rune, step int) rune {
	if r >= 'a' && r <= 'z' {
		shifted := int(r) + step
		if shifted > int('z') {
			shifted -= 26
		}
		return rune(shifted)
	}
	if r >= 'A' && r <= 'Z' {
		shifted := int(r) + step
		if shifted > int('Z') {
			shifted -= 26
		}
		return rune(shifted)
	}
	return r
}

// ReverseAlphabetValue function
func reverseAlphabetValue(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		reversed := 'a' + ('z' - ch)
		return reversed
	}

	if ch >= 'A' && ch <= 'Z' {
		reversed := 'A' + ('Z' - ch)
		return reversed
	}
	return ch
}

// CustomEncryption function
func customEncryption(ch rune) rune {
	key := "qwertyuiopasdfghjklzxcvbnm"

	if ch >= 'a' && ch <= 'z' {
		return rune(key[int(ch)-int('a')])
	}
	if ch >= 'A' && ch <= 'Z' {
		return rune(key[int(ch)-int('A')] - 32) // 32 is the difference between upper and lower case a to A
	}
	return ch
}

// CustomDecryption function

func customDecryption(ch rune) rune {
	key := "qwertyuiopasdfghjklzxcvbnm"
	for i := 0; i < len(key); i++ {
		if ch == rune(key[i]) {
			return rune(i + int('a'))
		} else if ch == rune(key[i] - 32) {
			return rune(i + int('A'))
		}
	}
	return ch
}
package main

import (
	"fmt"
	"strings"
)

// CustomEncryption function
func CustomEncryption(input string) string {
	key := "qwertyuiopasdfghjklzxcvbnm"
	var result strings.Builder

	for _, ch := range input {
		if ch >= 'a' && ch <= 'z' {
			result.WriteRune(rune(key[int(ch)-int('a')]))
		} else if ch >= 'A' && ch <= 'Z' {
			result.WriteRune(rune(key[int(ch)-int('A')] - 32)) // 32 is the difference between upper and lower case a to A
		} else {
			result.WriteRune(ch)
		}
	}

	return result.String()
}

// CustomDecryption function
func CustomDecryption(input string) string {
	key := "qwertyuiopasdfghjklzxcvbnm"
	var result strings.Builder

	for _, ch := range input {
		found := false
		for i := 0; i < len(key); i++ {
			if ch == rune(key[i]) {
				result.WriteRune(rune(i + int('a')))
				found = true
				break
			} else if ch == rune(key[i]-32) {
				result.WriteRune(rune(i + int('A')))
				found = true
				break
			}
		}
		if !found {
			result.WriteRune(ch)
		}
	}

	return result.String()
}

func main() {
	var choice int
	fmt.Println("Select an option:")
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	switch choice {
	case 1:
		encrypted := CustomEncryption(input)
		fmt.Println("Encrypted:", encrypted)
	case 2:
		decrypted := CustomDecryption(input)
		fmt.Println("Decrypted:", decrypted)
	default:
		fmt.Println("Invalid choice.")
	}
}

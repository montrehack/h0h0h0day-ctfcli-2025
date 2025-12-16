package main

import (
	"fmt"
	"unicode"
)

var allowedSpecialMap = map[rune]bool{
	'/':  true,
	'\'': true,
	'(':  true,
	'*':  true,
	')':  true,
	'=':  true,
	'-':  true,
	' ':  true,
}

func validateElfName(s string) error {
	for _, char := range s {
		if unicode.IsLower(char) && char >= 'a' && char <= 'z' {
			continue
		}

		if char >= '0' && char <= '9' {
			continue
		}

		if allowedSpecialMap[char] {
			continue
		}

		return fmt.Errorf("invalid character '%c' found in string: %s", char, s)
	}
	return nil
}

func main() {

	validateElfName("this is a test")
	validateElfName("validelfname")
	validateElfName("name=toto') or '2025-12-01' /*")

	targetRowY := 5
	targetCount := 5
	maxLength := 6

	clearNames := BruteForceNamesForClearRow(targetRowY, targetCount, maxLength)

	fmt.Println("\n--- FINAL RESULTS ---")
	fmt.Printf("Total names found: %d\n", len(clearNames))
	for i, name := range clearNames {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}

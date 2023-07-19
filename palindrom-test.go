package main

import (
	"fmt"
	"math"
	"strings"
)


func main() {
	wordList := []string{
		"SAIPPUAKIVIKAUPPIAS",
		"Aibohphobia",
		"Anna",
		"Civic",
		"My gym",
		"No lemon, no melon",
	}
	for x := 0; x <= len(wordList)-1; x++ {
		fmt.Println(isPalindrome(wordList[x]))
	}

}


func isPalindrome(words string) bool {
	words = strings.ReplaceAll(words, " ", "")
	words = strings.ToLower(words)
	startPoint := int(math.Ceil(float64(len(words)) / 2))
	result := true

	for i := 0; i < startPoint; i++ {
		if !result {
			break
		}

		chars := strings.Split(words, "")
		fromLeft := chars[i]
		fromRight := chars[len(chars)-1-i]

		if strings.Compare(fromLeft, fromRight) == 0 {
			continue
		} else {
			result = false
		}
	}
	return result
}

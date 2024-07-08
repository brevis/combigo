package combigo

import (
	"errors"
	"strings"
)

const Patterns = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateCombinations(pattern string, alphabets []string) []string {
	combinations := make([]string, GetCombinationsCount(pattern, alphabets))
	lastCombination := GetLastPatternCombination(pattern, alphabets)
	var currentCombination string = ""
	var err error
	i := 0
	for currentCombination != lastCombination {
		currentCombination, err = GetNextCombination(currentCombination, pattern, alphabets)
		if err != nil {
			break
		}
		combinations[i] = currentCombination
		i++
	}
	return combinations
}

func GetCombinationsCount(pattern string, alphabets []string) int {
	count := 1
	for _, r := range pattern {
		index := strings.Index(Patterns, string(r))
		if index != -1 && index < len(alphabets) {
			count *= len(alphabets[index])
		}
	}
	return count
}

func GetFirstPatternCombination(pattern string, alphabets []string) string {
	combination := ""
	for _, r := range pattern {
		index := strings.Index(Patterns, string(r))
		if index != -1 && index < len(alphabets) {
			combination += string(alphabets[index][0])
		} else {
			combination += string(r)
		}
	}
	return combination
}

func GetLastPatternCombination(pattern string, alphabets []string) string {
	combination := ""
	for _, r := range pattern {
		index := strings.Index(Patterns, string(r))
		if index != -1 && index < len(alphabets) {
			combination += string(alphabets[index][len(alphabets[index])-1])
		} else {
			combination += string(r)
		}
	}
	return combination
}

func GetNextCombination(combination string, pattern string, alphabets []string) (string, error) {
	if combination == "" {
		return GetFirstPatternCombination(pattern, alphabets), nil
	}
	combination = reverseString(combination)
	pattern = reverseString(pattern)
	var nextCombination string
	gettingNext := true
	for i := 0; i < len(pattern); i++ {
		patternCharacter := pattern[i]
		currentCharacter := string(combination[i])
		if gettingNext {
			index := strings.Index(Patterns, string(patternCharacter))
			if index != -1 && index < len(alphabets) {
				nextCharacter, err := GetNextCharacter(currentCharacter, alphabets[index])
				if err != nil {
					currentCharacter = string(alphabets[index][0])
				} else {
					gettingNext = false
					currentCharacter = string(nextCharacter)
				}
			}
		}
		nextCombination += currentCharacter
	}
	if gettingNext {
		return "", errors.New("unable to get next combination")
	}
	return reverseString(nextCombination), nil
}

func GetNextCharacter(character string, alphabet string) (string, error) {
	position := strings.Index(alphabet, character)
	if position == -1 || position > len(alphabet)-2 {
		return "", errors.New("unable to get next character")
	}
	return string(alphabet[position+1]), nil
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

package combigo

import (
	"slices"
	"testing"
)

func TestGetFirstCombinationMinimalCase(t *testing.T) {
	pattern := "A"
	alphabets := []string{"a", "b"}
	actual := GetFirstPatternCombination(pattern, alphabets)
	expected := "a"
	if actual != expected {
		t.Errorf("Error get first combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "B"
	alphabets = []string{"a", "b"}
	actual = GetFirstPatternCombination(pattern, alphabets)
	expected = "b"
	if actual != expected {
		t.Errorf("Error get first combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "AB"
	alphabets = []string{"a", "bc"}
	actual = GetFirstPatternCombination(pattern, alphabets)
	expected = "ab"
	if actual != expected {
		t.Errorf("Error get first combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "BA"
	alphabets = []string{"a", "bc"}
	actual = GetFirstPatternCombination(pattern, alphabets)
	expected = "ba"
	if actual != expected {
		t.Errorf("Error get first combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "AB"
	alphabets = []string{"ae", "bc"}
	actual = GetFirstPatternCombination(pattern, alphabets)
	expected = "ab"
	if actual != expected {
		t.Errorf("Error get first combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "ABABAB"
	alphabets = []string{"aeiouy", "bcdfghjklmnpqrstvwxz"}
	actual = GetFirstPatternCombination(pattern, alphabets)
	expected = "ababab"
	if actual != expected {
		t.Errorf("Error get first combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "ABABAB.com"
	alphabets = []string{"aeiouy", "bcdfghjklmnpqrstvwxz"}
	actual = GetFirstPatternCombination(pattern, alphabets)
	expected = "ababab.com"
	if actual != expected {
		t.Errorf("Error get first combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}
}

func TestGetLastCombinationMinimalCase(t *testing.T) {
	pattern := "A"
	alphabets := []string{"a", "b"}
	actual := GetLastPatternCombination(pattern, alphabets)
	expected := "a"
	if actual != expected {
		t.Errorf("Error get last combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "B"
	alphabets = []string{"a", "b"}
	actual = GetLastPatternCombination(pattern, alphabets)
	expected = "b"
	if actual != expected {
		t.Errorf("Error get last combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "AB"
	alphabets = []string{"a", "bc"}
	actual = GetLastPatternCombination(pattern, alphabets)
	expected = "ac"
	if actual != expected {
		t.Errorf("Error get last combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "BA"
	alphabets = []string{"a", "bc"}
	actual = GetLastPatternCombination(pattern, alphabets)
	expected = "ca"
	if actual != expected {
		t.Errorf("Error get last combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "AB"
	alphabets = []string{"ae", "bc"}
	actual = GetLastPatternCombination(pattern, alphabets)
	expected = "ec"
	if actual != expected {
		t.Errorf("Error get last combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "ABABAB"
	alphabets = []string{"aeiouy", "bcdfghjklmnpqrstvwxz"}
	actual = GetLastPatternCombination(pattern, alphabets)
	expected = "yzyzyz"
	if actual != expected {
		t.Errorf("Error get last combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	pattern = "ABABAB.com"
	alphabets = []string{"aeiouy", "bcdfghjklmnpqrstvwxz"}
	actual = GetLastPatternCombination(pattern, alphabets)
	expected = "yzyzyz.com"
	if actual != expected {
		t.Errorf("Error get last combination for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}
}

func TestGetNextCharachter(t *testing.T) {
	alphabet := "aeiouy"
	charachter := "a"
	expected := "e"
	actual, _ := GetNextCharachter(charachter, alphabet)
	if actual != expected {
		t.Errorf("Error get next charachter for charachter \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", charachter, alphabet, expected, actual)
	}

	charachter = "u"
	expected = "y"
	actual, _ = GetNextCharachter(charachter, alphabet)
	if actual != expected {
		t.Errorf("Error get next charachter for charachter \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", charachter, alphabet, expected, actual)
	}

	charachter = "y"
	_, err := GetNextCharachter(charachter, alphabet)
	if err == nil {
		t.Errorf("Error get next charachter for charachter \"%s\" and alpabets \"%v\": expected error, got no error", charachter, alphabet)
	}

	charachter = "b"
	_, err = GetNextCharachter(charachter, alphabet)
	if err == nil {
		t.Errorf("Error get next charachter for charachter \"%s\" and alpabets \"%v\": expected error, got no error", charachter, alphabet)
	}
}

func TestGetNextCombination(t *testing.T) {
	alphabets := []string{"aeiouy", "bcdfghjklmnpqrstvwx"}
	combination := ""
	pattern := "A"
	expected := "a"
	actual, _ := GetNextCombination(combination, pattern, alphabets)
	if actual != expected {
		t.Errorf("Error get next combination for combination \"%s\", pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", combination, pattern, alphabets, expected, actual)
	}

	combination = "a"
	pattern = "A"
	expected = "e"
	actual, _ = GetNextCombination(combination, pattern, alphabets)
	if actual != expected {
		t.Errorf("Error get next combination for combination \"%s\", pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", combination, pattern, alphabets, expected, actual)
	}

	combination = "aaa"
	pattern = "AAA"
	expected = "aae"
	actual, _ = GetNextCombination(combination, pattern, alphabets)
	if actual != expected {
		t.Errorf("Error get next combination for combination \"%s\", pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", combination, pattern, alphabets, expected, actual)
	}

	combination = "ab"
	pattern = "AB"
	expected = "ac"
	actual, _ = GetNextCombination(combination, pattern, alphabets)
	if actual != expected {
		t.Errorf("Error get next combination for combination \"%s\", pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", combination, pattern, alphabets, expected, actual)
	}

	combination = "az"
	pattern = "AB"
	expected = "eb"
	actual, _ = GetNextCombination(combination, pattern, alphabets)
	if actual != expected {
		t.Errorf("Error get next combination for combination \"%s\", pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", combination, pattern, alphabets, expected, actual)
	}

	combination = "yz"
	pattern = "AB"
	_, err := GetNextCombination(combination, pattern, alphabets)
	if err == nil {
		t.Errorf("Error get next combination for combination \"%s\", pattern \"%s\" and alpabets \"%v\": expected error, got no error", combination, pattern, alphabets)
	}
}

func TestGenerateMinimalCase(t *testing.T) {
	alphabets := []string{"a", "b"}
	pattern := "AB"
	expected := []string{"ab"}
	actual := GenerateCombinations(pattern, alphabets)
	if !slices.Equal(actual, expected) {
		t.Errorf("Error generate combinations for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}
}

func TestGenerateTwoCharsAlphabets(t *testing.T) {
	alphabets := []string{"ab"}
	pattern := "AA"
	expected := []string{"aa", "ab", "ba", "bb"}
	actual := GenerateCombinations(pattern, alphabets)
	if !slices.Equal(actual, expected) {
		t.Errorf("Error generate two chars aplhabet combinations for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}

	alphabets = []string{"ab", "cd"}
	pattern = "AB"
	expected = []string{"ac", "ad", "bc", "bd"}
	actual = GenerateCombinations(pattern, alphabets)
	if !slices.Equal(actual, expected) {
		t.Errorf("Error generate two chars aplhabet combinations for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}
}

func TestGenerateThreeAlphabets(t *testing.T) {
	alphabets := []string{"abc", "def", "ghi"}
	pattern := "ABC"
	expected := []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}
	actual := GenerateCombinations(pattern, alphabets)
	if !slices.Equal(actual, expected) {
		t.Errorf("Error generate three chars aplhabet combinations for pattern \"%s\" and alpabets \"%v\": expected \"%s\", got \"%s\"", pattern, alphabets, expected, actual)
	}
}

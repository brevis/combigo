package combigo

import (
	"slices"
	"testing"
)

// Helper function to test combinations
func testCombination(t *testing.T, pattern string, alphabets []string, actual, expected, testType string) {
	if actual != expected {
		t.Errorf("Error %s combination for pattern \"%s\" and alphabets \"%v\": expected \"%s\", got \"%s\"", testType, pattern, alphabets, expected, actual)
	}
}

// Helper function to run tests for first combinations
func runFirstCombinationTests(t *testing.T, cases []struct {
	pattern   string
	alphabets []string
	expected  string
}) {
	for _, c := range cases {
		actual := GetFirstPatternCombination(c.pattern, c.alphabets)
		testCombination(t, c.pattern, c.alphabets, actual, c.expected, "get first")
	}
}

func TestGetFirstCombinationMinimalCase(t *testing.T) {
	cases := []struct {
		pattern   string
		alphabets []string
		expected  string
	}{
		{"A", []string{"a", "b"}, "a"},
		{"B", []string{"a", "b"}, "b"},
		{"AB", []string{"a", "bc"}, "ab"},
		{"BA", []string{"a", "bc"}, "ba"},
		{"AB", []string{"ae", "bc"}, "ab"},
		{"ABABAB", []string{"aeiouy", "bcdfghjklmnpqrstvwxz"}, "ababab"},
		{"ABABAB.com", []string{"aeiouy", "bcdfghjklmnpqrstvwxz"}, "ababab.com"},
	}

	runFirstCombinationTests(t, cases)
}

// Helper function to run tests for last combinations
func runLastCombinationTests(t *testing.T, cases []struct {
	pattern   string
	alphabets []string
	expected  string
}) {
	for _, c := range cases {
		actual := GetLastPatternCombination(c.pattern, c.alphabets)
		testCombination(t, c.pattern, c.alphabets, actual, c.expected, "get last")
	}
}

func TestGetLastCombinationMinimalCase(t *testing.T) {
	cases := []struct {
		pattern   string
		alphabets []string
		expected  string
	}{
		{"A", []string{"a", "b"}, "a"},
		{"B", []string{"a", "b"}, "b"},
		{"AB", []string{"a", "bc"}, "ac"},
		{"BA", []string{"a", "bc"}, "ca"},
		{"AB", []string{"ae", "bc"}, "ec"},
		{"ABABAB", []string{"aeiouy", "bcdfghjklmnpqrstvwxz"}, "yzyzyz"},
		{"ABABAB.com", []string{"aeiouy", "bcdfghjklmnpqrstvwxz"}, "yzyzyz.com"},
	}

	runLastCombinationTests(t, cases)
}

func TestGetNextCharacter(t *testing.T) {
	cases := []struct {
		alphabet  string
		character string
		expected  string
		expectErr bool
	}{
		{"aeiouy", "a", "e", false},
		{"aeiouy", "u", "y", false},
		{"aeiouy", "y", "", true},
		{"aeiouy", "b", "", true},
	}

	for _, c := range cases {
		actual, err := GetNextCharacter(c.character, c.alphabet)
		if (err != nil) != c.expectErr {
			t.Errorf("Error get next character for character \"%s\" and alphabet \"%v\": expected error %v, got error %v", c.character, c.alphabet, c.expectErr, err != nil)
		}
		if actual != c.expected && !c.expectErr {
			t.Errorf("Error get next character for character \"%s\" and alphabet \"%v\": expected \"%s\", got \"%s\"", c.character, c.alphabet, c.expected, actual)
		}
	}
}

// Helper function to run tests for next combinations
func runNextCombinationTests(t *testing.T, cases []struct {
	combination string
	pattern     string
	alphabets   []string
	expected    string
	expectErr   bool
}) {
	for _, c := range cases {
		actual, err := GetNextCombination(c.combination, c.pattern, c.alphabets)
		if (err != nil) != c.expectErr {
			t.Errorf("Error get next combination for combination \"%s\", pattern \"%s\" and alphabets \"%v\": expected error %v, got error %v", c.combination, c.pattern, c.alphabets, c.expectErr, err != nil)
		}
		if actual != c.expected && !c.expectErr {
			t.Errorf("Error get next combination for combination \"%s\", pattern \"%s\" and alphabets \"%v\": expected \"%s\", got \"%s\"", c.combination, c.pattern, c.alphabets, c.expected, actual)
		}
	}
}

func TestGetNextCombination(t *testing.T) {
	cases := []struct {
		combination string
		pattern     string
		alphabets   []string
		expected    string
		expectErr   bool
	}{
		{"", "A", []string{"aeiouy", "bcdfghjklmnpqrstvwx"}, "a", false},
		{"a", "A", []string{"aeiouy", "bcdfghjklmnpqrstvwx"}, "e", false},
		{"aaa", "AAA", []string{"aeiouy", "bcdfghjklmnpqrstvwx"}, "aae", false},
		{"ab", "AB", []string{"aeiouy", "bcdfghjklmnpqrstvwx"}, "ac", false},
		{"az", "AB", []string{"aeiouy", "bcdfghjklmnpqrstvwx"}, "eb", false},
		{"yz", "AB", []string{"aeiouy", "bcdfghjklmnpqrstvwx"}, "", true},
	}

	runNextCombinationTests(t, cases)
}

func TestGenerateMinimalCase(t *testing.T) {
	cases := []struct {
		alphabets []string
		pattern   string
		expected  []string
	}{
		{[]string{"a", "b"}, "AB", []string{"ab"}},
	}

	for _, c := range cases {
		actual := GenerateCombinations(c.pattern, c.alphabets)
		if !slices.Equal(actual, c.expected) {
			t.Errorf("Error generate combinations for pattern \"%s\" and alphabets \"%v\": expected \"%v\", got \"%v\"", c.pattern, c.alphabets, c.expected, actual)
		}
	}
}

func TestGenerateTwoCharsAlphabets(t *testing.T) {
	cases := []struct {
		alphabets []string
		pattern   string
		expected  []string
	}{
		{[]string{"ab"}, "AA", []string{"aa", "ab", "ba", "bb"}},
		{[]string{"ab", "cd"}, "AB", []string{"ac", "ad", "bc", "bd"}},
	}

	for _, c := range cases {
		actual := GenerateCombinations(c.pattern, c.alphabets)
		if !slices.Equal(actual, c.expected) {
			t.Errorf("Error generate two chars alphabet combinations for pattern \"%s\" and alphabets \"%v\": expected \"%v\", got \"%v\"", c.pattern, c.alphabets, c.expected, actual)
		}
	}
}

func TestGenerateThreeAlphabets(t *testing.T) {
	cases := []struct {
		alphabets []string
		pattern   string
		expected  []string
	}{
		{[]string{"abc", "def", "ghi"}, "ABC", []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}},
	}

	for _, c := range cases {
		actual := GenerateCombinations(c.pattern, c.alphabets)
		if !slices.Equal(actual, c.expected) {
			t.Errorf("Error generate three chars alphabet combinations for pattern \"%s\" and alphabets \"%v\": expected \"%v\", got \"%v\"", c.pattern, c.alphabets, c.expected, actual)
		}
	}
}

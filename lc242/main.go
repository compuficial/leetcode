package main

import "fmt"

func isAnagram(s string, t string) bool {
	cm := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, c := range s {
		cm[c]++
	}

	for _, c := range t {
		cm[c]--

		if cm[c] == 0 {
			delete(cm, c)
		}
	}

	return len(cm) == 0
}

func main() {
	testCases := []struct {
		s        string
		t        string
		expected bool
		desc     string
	}{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
			desc:     "Example 1: Basic anagram",
		},
		{
			s:        "rat",
			t:        "car",
			expected: false,
			desc:     "Example 2: Not an anagram",
		},
		{
			s:        "",
			t:        "",
			expected: true,
			desc:     "Edge case: Empty strings",
		},
		{
			s:        "a",
			t:        "a",
			expected: true,
			desc:     "Edge case: Single character, same",
		},
		{
			s:        "a",
			t:        "b",
			expected: false,
			desc:     "Edge case: Single character, different",
		},
		{
			s:        "ab",
			t:        "a",
			expected: false,
			desc:     "Edge case: Different lengths",
		},
		{
			s:        "aacc",
			t:        "ccac",
			expected: false,
			desc:     "Edge case: Same characters but different counts",
		},
		{
			s:        "texttwisttime",
			t:        "timetwisttext",
			expected: true,
			desc:     "Longer anagram",
		},
		{
			s:        "ΑΒΓ",
			t:        "ΓΒΑ",
			expected: true,
			desc:     "Non-ASCII characters",
		},
	}

	for i, tc := range testCases {
		result := isAnagram(tc.s, tc.t)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("Test %d: %s\n", i+1, tc.desc)
		fmt.Printf("  s = %q, t = %q\n", tc.s, tc.t)
		fmt.Printf("  Expected: %t, Got: %t %s\n\n", tc.expected, result, status)
	}
}

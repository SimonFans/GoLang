package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		inputString string
		expect      bool
	}{
		{"sas", true},
		{"ass", false},
		// {"aaa", false},
	}
	for _, test := range tests {
		if output := IsPalindrome(test.inputString); output != test.expect {
			t.Errorf("IsPalindrome(%q)=%v, expect to return %v but get %v\n", test.inputString, test.expect, test.expect, output)
		}
	}
}

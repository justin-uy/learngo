package stringutils

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct{ input, expect string }{
		{input: "", expect: ""},
		{input: "a", expect: "a"},
		{input: "ab", expect: "ba"},
		{input: "abc", expect: "cba"},
	}

	for i, testCase := range cases {
		out := Reverse(testCase.input)
		if out != testCase.expect {
			t.Errorf("Test %v - Expected %v; Got %v", i, testCase.expect, out)
		}
	}
}

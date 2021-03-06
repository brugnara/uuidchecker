package uuidchecker

import "testing"

func TestIsValid(t *testing.T) {
	for _, test := range []struct {
		in  string
		out bool
	}{
		{"01051e78-fb56-4776-8647-6dd214b8b7f6", true},
		{"junk01051e78-fb56-4776-8647-6dd214b8b7f6", false},
		{"01051e78-fb56-4776-8647-6dd214b8b7f6junk", false},
		{"junk01051e78-fb56-4776-8647-6dd214b8b7f6junk", false},
		{"01051e78-fb56-8647-6dd214b8b7f6", false},
		{"", false},
	} {
		if out := IsValid(test.in); out != test.out {
			t.Errorf("Expected: %s to be: %v, got: %v", test.in, test.out, out)
		}
	}
}

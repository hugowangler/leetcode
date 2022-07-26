package leetcode

import "testing"

func TestParseBoolExpr(t *testing.T) {
	// table test
	tests := []struct {
		expression string
		expected   bool
	}{
		{"|(&(t,f,t),!(t))", false},
		{"!(f)", true},
		{"|(f,t)", true},
		{"&(t,f)", false},
	}
	for _, test := range tests {
		if actual := parseBoolExpr(test.expression); actual != test.expected {
			t.Errorf("ParseBoolExpr(%s) = %v, expected %v", test.expression, actual, test.expected)
		}
	}
}

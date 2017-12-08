package quiz

import (
	"testing"
)

func TestIsAnswerCorrect(t *testing.T) {
	a1 := Answer{"a1", true}
	a2 := Answer{"a2", false}
	a3 := Answer{"a3", false}
	a4 := Answer{"a4", false}
	q := Question{"", []Answer{a4, a3, a2, a1}, ""}

	var tests = []struct {
		User   UserAnswer
		Expect bool
	}{
		{UserAnswer{"", &q, &a1}, true},
		{UserAnswer{"", &q, &a2}, false},
	}

	for _, test := range tests {
		if test.Expect != test.User.IsAnswerCorrect() {
			t.Errorf("expected %t but go %t", test.Expect, test.User.IsAnswerCorrect)
		}
	}
}

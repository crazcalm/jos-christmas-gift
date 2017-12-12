package quiz

import (
	"testing"
)

func TestTotalScore(t *testing.T) {
	a1 := Answer{"a1", true}
	a2 := Answer{"a2", false}
	a3 := Answer{"a3", false}
	a4 := Answer{"a4", false}
	q := Question{"", []Answer{a4, a3, a2, a1}, ""}

	u1 := UserAnswer{"", &q, &a1}
	u2 := UserAnswer{"", &q, &a2}
	u3 := UserAnswer{"", &q, &a3}
	u4 := UserAnswer{"", &q, &a1}

	var tests = []struct {
		Answers         []UserAnswer
		ExpectedCorrect int
		ExpectedTotal   int
	}{
		{[]UserAnswer{u1, u2, u3, u4}, 2, 4},
	}

	for _, test := range tests {
		right, total := TotalScore(test.Answers)

		if right != test.ExpectedCorrect {
			t.Errorf("Expected %d correct answers, but got %d correct answers", test.ExpectedCorrect, right)
		}

		if total != test.ExpectedTotal {
			t.Errorf("Expected %d, but got %d", test.ExpectedTotal, total)
		}
	}
}

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

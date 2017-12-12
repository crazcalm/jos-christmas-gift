package quiz

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintSolution(t *testing.T) {
	a1 := Answer{"a1", true}
	a2 := Answer{"a2", false}
	a3 := Answer{"a3", false}
	a4 := Answer{"a4", false}

	q := Question{"question", []Answer{a1, a2, a3, a4}, "explanation"}

	correctAnswer := UserAnswer{
		BoxA,
		&q,
		&a1,
	}

	wrongAnswer := UserAnswer{
		BoxC,
		&q,
		&a3,
	}

	var tests = []struct {
		Num      int
		User     UserAnswer
		Expected []string
	}{
		{1, correctAnswer, []string{"1 --", "Question: question", "You correctly selected A", "explanation"}},
		{2, wrongAnswer, []string{"2 --", "Question: question", "You selected C, which is wrong", "explanation"}},
	}

	//Used to check the results
	b := new(bytes.Buffer)

	for _, test := range tests {
		tmpl, err := PrintSolution(test.Num, test.User)
		if err != nil {
			t.Error("Error occured while trying to make the template")
		}
		err = tmpl.Execute(b, test.User)
		if err != nil {
			t.Error("Experienced an unexpected error")
		}
		got := b.String()

		for _, substring := range test.Expected {
			if !strings.Contains(got, substring) {
				t.Errorf("Expected: %s\nReceived: %s", substring, got)
			}
		}

		//Clear buffer
		b.Reset()
	}
}

func TestAnswerQuestionRatio(t *testing.T) {
	var tests = []struct {
		Answers   int
		Questions int
		Expect    string
	}{
		{0, 10, "0/10"},
		{2, 12, "2/12"},
	}

	//buffer used to check results
	b := new(bytes.Buffer)

	for _, test := range tests {
		templ, err := AnswerQuestionRatio(test.Answers, test.Questions)
		if err != nil {
			t.Error("Error occured while trying to create the template")
		}

		err = templ.Execute(b, nil)
		if err != nil {
			t.Error("Experienced and unexpected error")
		}

		result := b.String()
		if !strings.EqualFold(result, test.Expect) {
			t.Errorf("Expected %s, but got %s", test.Expect, result)
		}

		//Clears the buffer
		b.Reset()
	}
}

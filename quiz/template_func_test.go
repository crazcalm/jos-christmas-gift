package quiz

import (
	"bytes"
	"strings"
	"testing"
)

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

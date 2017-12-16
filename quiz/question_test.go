package quiz

import (
	"fmt"
	"strings"
	"testing"
)

func TestCorrectAnswer(t *testing.T) {
	a1 := Answer{"a1", true}
	a2 := Answer{"a2", false}
	a3 := Answer{"a3", false}
	a4 := Answer{"a4", false}

	q := Question{"question", []Answer{a4, a3, a2, a1}, ""}
	result := &q

	answer := result.CorrectAnswer()
	if !strings.EqualFold(a1.Answer, answer.Answer) {
		t.Error("Unable to select the correct answer for this question")
	}
}

func TestShuffleQuestions(t *testing.T) {
	a1 := Answer{"1", true}
	a2 := Answer{"2", true}
	a3 := Answer{"3", true}
	a4 := Answer{"4", true}

	//Empty the Questions pool
	Questions = []Question{}

	//Fill up the Questions pool
	for i := 0; i < 10; i++ {
		q := fmt.Sprintf("question%d", i)
		Questions = append(Questions, Question{q, []Answer{a1, a2, a3, a4}, ""})
	}

	firstQuestion := Questions[0].Question
	shuffled := false

	for i := 0; i < 10; i++ {
		ShuffleQuestions(Questions)

		if !strings.EqualFold(firstQuestion, Questions[i].Question) {
			shuffled = true
		}
	}

	if !shuffled {
		t.Error("ShuffleQuestions did not shuffle the questions...")
	}
}

func TestShuffleAnswers(t *testing.T) {
	a1 := Answer{"1", true}
	a2 := Answer{"2", true}
	a3 := Answer{"3", true}
	a4 := Answer{"4", true}

	q := Question{"question", []Answer{a1, a2, a3, a4}, ""}

	shuffled := false

	for i := 0; i < 10; i++ {
		q.ShuffleAnswers()

		if !strings.EqualFold(a1.Answer, q.Answers[0].Answer) {
			shuffled = true
		}
	}

	if !shuffled {
		t.Error("Question.ShuffleAnswers is not shuffling the Answers")
	}

}

func TestCurrentQuestion(t *testing.T) {
	a := Answer{"answer", true}

	q1 := Question{"q1", []Answer{a, a, a, a}, "none"}
	q2 := Question{"q2", []Answer{a, a, a, a}, "none"}
	q3 := Question{"q3", []Answer{a, a, a, a}, "none"}

	//clear Questions
	Questions = []Question{}

	//fill up the questions
	Questions = append(Questions, q1, q2, q3)

	var tests = []struct {
		Count    int
		Question Question
	}{
		{0, q1},
		{1, q2},
		{2, q3},
	}

	for _, test := range tests {
		QuestionCount = test.Count

		answer := currentQuestion()
		if !strings.EqualFold(answer.Question, test.Question.Question) {
			t.Error("currentQuestion returned the wrong question")
		}
	}
}

func TestNextQuestion(t *testing.T) {
	a := Answer{"answer", true}

	q1 := Question{"q1", []Answer{a, a, a, a}, "none"}
	q2 := Question{"q2", []Answer{a, a, a, a}, "none"}
	q3 := Question{"q3", []Answer{a, a, a, a}, "none"}

	Questions = []Question{}
	QuestionCount = 0
	Questions = append(Questions, q1, q2, q3)

	var tests = []struct {
		Count        int
		ExpectError  bool
		NextQuestion Question
	}{
		{1, false, q2},
		{2, false, q3},
		{3, true, q3},
	}

	for _, test := range tests {
		question, err := nextQuestion()

		if test.ExpectError {
			if err == nil {
				t.Error("Was expecting an error, but err == nil...")
			}
			return
		}

		if !strings.EqualFold(question.Question, test.NextQuestion.Question) {
			t.Errorf("Was expecting this question '%s' but got '%s'", test.NextQuestion.Question, question.Question)
			return
		}

		if QuestionCount != test.Count {
			t.Errorf("QuestionCount was expected to be %d, but was %d", test.Count, QuestionCount)
			return
		}
	}
}

func TestNextQuestionExist(t *testing.T) {
	a := Answer{"answer", true}
	q1 := Question{"q1", []Answer{a, a, a, a}, "none"}
	q2 := Question{"q2", []Answer{a, a, a, a}, "none"}
	q3 := Question{"q3", []Answer{a, a, a, a}, "none"}

	Questions = []Question{}
	Questions = append(Questions, q1, q2, q3)

	var tests = []struct {
		SetCount int
		Result   bool
	}{
		{0, true},
		{1, true},
		{2, false},
		{3, false},
	}

	for _, test := range tests {
		QuestionCount = test.SetCount
		answer := nextQuestionExist()

		if answer != test.Result {
			t.Errorf("Expected %t, but got %t: len(Questions) -- %d", test.Result, answer, len(Questions))
		}
	}
}

func TestCreateQuestions(t *testing.T) {
	question := "question"
	answer1 := "answer1"
	answer2 := "answer2"
	answer3 := "answer3"
	answer4 := "answer4"

	var tests = []struct {
		Records     [][]string
		ExpectError bool
		Question    string
		Answers     []string
	}{
		{[][]string{{question, answer1, answer2, answer3, answer4}}, true, question, []string{answer1, answer2, answer3, answer4}},
		{[][]string{{}}, true, "", []string{}},
		{[][]string{{question, answer1, answer2, answer3, answer4}, {question, answer1, answer2, answer3, answer4}}, false, question, []string{answer1, answer2, answer3, answer4}},
		{[][]string{{question, answer1, answer2, answer3, answer4}, {question, answer1, answer2, answer3, answer4}, {question, answer1, answer2, answer3, answer4}}, false, question, []string{answer1, answer2, answer3, answer4}},
	}

	for _, test := range tests {
		err := CreateQuestions(test.Records)
		if test.ExpectError {
			if err == nil {
				t.Error("Was expecting an error, but err == nil...")
			}
			return
		}

		if err != nil {
			t.Error("err is not equal to nil, but it is epxected to be nil")
		}

		for _, q := range Questions {
			if strings.EqualFold(question, q.Question) {
				t.Errorf("Expected %s, but got %s", question, q.Question)
			}
		}
	}
}

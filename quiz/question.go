package quiz

import (
	"fmt"
)

var (
	//Questions -- A slice of questions
	Questions = []Question{}
	//QuestionCount -- A counter used to keep track of which question was are using
	QuestionCount = 0
)

//Question -- struct to hold a question and its answers
type Question struct {
	Question     string
	Answers      []Answer
	Explaination string
}

func currentQuestion() Question {
	return Questions[QuestionCount]
}

func nextQuestionExist() bool {
	result := false
	if QuestionCount < len(Questions)-1 {
		result = true
	}
	return result
}

func nextQuestion() (q Question, err error) {
	if QuestionCount >= len(Questions)-1 {
		err = fmt.Errorf("No more questions")
		return q, err
	}

	QuestionCount = QuestionCount + 1
	return Questions[QuestionCount], err
}

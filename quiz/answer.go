package quiz

import (
	"strings"
)

var (
	userAnswers = []UserAnswer{}
)

//Answer -- struct to hold an answer
type Answer struct {
	Answer  string
	Correct bool
}

//UserAnswer -- struct to hold individual answers made by the user
type UserAnswer struct {
	Answer     string
	Question   *Question
	UserAnswer *Answer
}

//IsAnswerCorrect -- returns whether or not the user's answer is correct
func (u UserAnswer) IsAnswerCorrect() bool {
	result := false
	answer := u.Question.CorrectAnswer()
	if strings.EqualFold(answer.Answer, u.UserAnswer.Answer) {
		result = true
	}
	return result
}

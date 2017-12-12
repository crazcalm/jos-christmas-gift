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

//TotalScore -- returns the number of correct answers and the total number of questions
func TotalScore(answers []UserAnswer) (right, total int) {
	total = len(answers)

	for _, answer := range answers {
		if answer.IsAnswerCorrect() {
			right++
		}
	}

	return
}

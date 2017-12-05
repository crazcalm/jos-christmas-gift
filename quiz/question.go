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

//CreateQuestions -- Creates the questions that are stored in the Questions variable
func CreateQuestions(records [][]string) {
	for i := 1; i < len(records); i++ {

		a1 := Answer{records[i][1], true}
		a2 := Answer{records[i][2], false}
		a3 := Answer{records[i][3], false}
		a4 := Answer{records[i][4], false}

		question := Question{
			records[i][0],
			[]Answer{a1, a2, a3, a4},
			records[i][5],
		}
		Questions = append(Questions, question)
	}
}

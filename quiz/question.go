package quiz

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//Seeding the random number generator
	rand.Seed(time.Now().UnixNano())
}

var (
	//Questions -- A slice of questions
	Questions = []Question{}
	//QuestionCount -- A counter used to keep track of which question was are using
	QuestionCount = 0
	//QuestionsLimit -- set the number of questions that will be used during the quiz
	QuestionsLimit = 10
)

//Question -- struct to hold a question and its answers
type Question struct {
	Question     string
	Answers      []Answer
	Explaination string
}

//ShuffleQuestions -- Shuffle the list of questions
func ShuffleQuestions(qs []Question) {
	numOfQuestions := len(qs)
	var tempt Question
	var swapIndex int

	for index := range qs {
		swapIndex = rand.Intn(numOfQuestions)
		tempt = qs[index]
		qs[index] = qs[swapIndex]
		qs[swapIndex] = tempt
	}
}

//CorrectAnswer -- returns the correct answer to the question
func (q *Question) CorrectAnswer() Answer {
	var result Answer
	for _, a := range q.Answers {
		if a.Correct {
			result = a
			break
		}
	}
	return result
}

//ShuffleAnswers -- Does and in place shuffle of the answers
func (q Question) ShuffleAnswers() {
	numOfAnswers := len(q.Answers)
	var tempt Answer
	var swapIndex int

	for index := range q.Answers {
		swapIndex = rand.Intn(numOfAnswers)
		tempt = q.Answers[index]
		q.Answers[index] = q.Answers[swapIndex]
		q.Answers[swapIndex] = tempt
	}
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
func CreateQuestions(records [][]string) (err error) {
	if len(records) < 2 {
		err = fmt.Errorf("The length of records is expected to be greater that 1. Current length is %d", len(records))
		return
	}

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
		question.ShuffleAnswers()
		Questions = append(Questions, question)
	}
	//Shuffle Questions
	ShuffleQuestions(Questions)

	return
}

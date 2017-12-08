package quiz

import (
	"fmt"
	"log"
	"text/template"
)

//AnswerQuestionRatio -- A ratio used to let the user know how many questions are left
func AnswerQuestionRatio(a, q int) (tmpl *template.Template, err error) {
	ratio := fmt.Sprintf("%d/%d", a, q)
	tmpl, err = template.New("AnswerQuestionRatio").Parse(ratio)
	if err != nil {
		log.Println(err)
	}
	return
}

//PrintQuestion -- Formats the question for the end screen
func PrintQuestion(q *Question, num int) (tmpl *template.Template, err error) {
	statement := fmt.Sprintf("%d: %s", num, q.Question)
	tmpl, err = template.New("PrintQuestion").Parse(statement)
	if err != nil {
		log.Println(err)
	}
	return
}

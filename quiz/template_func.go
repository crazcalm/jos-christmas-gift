package quiz

import (
	"fmt"
	"log"
	"text/template"
)

//PrintSolution -- Prints out the solution text for a given userAnswer
func PrintSolution(u UserAnswer) (tmpl *template.Template, err error) {
	var result string

	//Figures out which answer the user selected
	selectedAnswer := answersToLetters[u.Answer]

	//Figures out if the answer the user selected is correct or not
	correct := u.IsAnswerCorrect()

	if correct {
		result = fmt.Sprintf("You correctly selected %s", selectedAnswer)
	} else {
		result = fmt.Sprintf("You selected %s, which is wrong", selectedAnswer)
	}

	tmplString := fmt.Sprintf("Question: {{.Question.Question}}\n\n- %s\n\n- Correct Answer: %s", result, u.Question.Explaination)

	tmpl, err = template.New("PrintSolution").Parse(tmplString)
	if err != nil {
		log.Println(err)
	}
	return
}

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

package main

import (
	"encoding/csv"
	"fmt"
	"github.com/jroimartin/gocui"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"github.com/crazcalm/jos-christmas-gift/quiz"
)

var (
	boxesView         = []string{quiz.BoxA, quiz.BoxB, quiz.BoxC, quiz.BoxD}
	activeView        = 0
	questions         = []quiz.Question{}
	questionCount     = 0
	answersToBoxViews = make(map[string]quiz.Answer)
	userAnswers       = []quiz.UserAnswer{}
)

func readCSV() [][]string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
	for _, f := range files {
		fmt.Println(f.Name())
	}
	testFile := "testing.csv"
	_, err = os.Stat(testFile)
	if os.IsNotExist(err) {
		log.Fatalf("file: %s does not exist", testFile)
	}

	file, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	var records [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
		records = append(records, record)
	}

	return records
}

func createQuestions(records [][]string) {
	for i := 1; i < len(records); i++ {

		a1 := quiz.Answer{records[i][1], true}
		a2 := quiz.Answer{records[i][2], false}
		a3 := quiz.Answer{records[i][3], false}
		a4 := quiz.Answer{records[i][4], false}

		question := quiz.Question{
			records[i][0],
			[]quiz.Answer{a1, a2, a3, a4},
			records[i][5],
		}
		questions = append(questions, question)
	}
}

func currentQuestion() quiz.Question {
	return questions[questionCount]
}

func nextQuestionExist() bool {
	result := false
	if questionCount < len(questions) - 1 {
		result = true
	}
	return result
}

func nextQuestion() (q quiz.Question, err error) {
	if questionCount >= len(questions)-1 {
		err = fmt.Errorf("No more questions")
		return q, err
	}

	questionCount = questionCount + 1
	return questions[questionCount], err
}

func writeInfoToLayout(g *gocui.Gui, q quiz.Question) {
	//Write question
	questionBox := getQuestionBoxView(g)
	fmt.Fprintln(questionBox, q.Question)

	//Write answers
	answerBoxViews := getAnswerBoxViews(g)
	for i, answer := range q.Answers {
		answerBoxViews[i].Clear()

		//Adding it to the map
		answersToBoxViews[answerBoxViews[i].Name()] = answer

		//Write the answer to the layout
		fmt.Fprintln(answerBoxViews[i], answer.Answer)
	}

}

func isViewInSlice(viewNames []string, v *gocui.View) bool {
	result := false
	for _, name := range viewNames {
		if strings.EqualFold(name, v.Name()) == true {
			result = true
		}
	}
	return result
}

func selectAnswer(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintln(v, "Selected")

	cQuestion := currentQuestion()
	selectedAnswer := answersToBoxViews[v.Name()]

	a := quiz.UserAnswer{
		v.Name(),
		&cQuestion,
		&selectedAnswer,
	}

	//User answers
	userAnswers = append(userAnswers, a)

	views := g.Views()

	if !nextQuestionExist(){
		g.SetManagerFunc(endScreenLayout)
		err := g.SetKeybinding("", gocui.KeyCtrlC,gocui.ModNone, quit)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
	question, err := nextQuestion()
	if err != nil {
		log.Fatal(err)
	}

	//Write Question and Answers to layout
	writeInfoToLayout(g, question)

	for i, view := range views {

		if strings.EqualFold(quiz.QuestionBox, view.Name()) {
			view.Clear() //clear the question box
			fmt.Fprintf(view, "%d - %s -- question box", i, view.Name())
		}

	}
	return nil
}

func getQuestionBoxView(g *gocui.Gui) *gocui.View {
	var result *gocui.View
	views := g.Views()
	for _, view := range views {
		if strings.EqualFold(quiz.QuestionBox, view.Name()) {
			result = view
		}
	}
	return result
}

func getAnswerBoxViews(g *gocui.Gui) []*gocui.View {
	var questionViews []*gocui.View
	views := g.Views()
	for _, view := range views {
		if isViewInSlice(boxesView, view) {
			questionViews = append(questionViews, view)
		}
	}
	return questionViews
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (activeView + 1) % len(boxesView)
	name := boxesView[nextIndex]

	_, err := setCurrentViewOnTop(g, name)
	if err != nil {
		log.Panicln(err)
	}

	activeView = nextIndex
	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	_, err := g.SetCurrentView(name)
	if err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)

}

func endScreenLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(quiz.EndScreen, -1, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Quiz Over"
		fmt.Fprintln(v, "End of quiz")
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView(quiz.QuestionFrame, -1, -1, maxX, int(0.5*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	if v, err := g.SetView(quiz.QuestionBox, int(0.2*float32(maxX)), int(0.1*float32(maxY)), int(0.8*float32(maxX)), int(0.4*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Question"
		fmt.Fprintln(v, questions[questionCount].Question)
	}

	//Answer Box A
	if v, err := g.SetView(quiz.BoxA, -1, int(0.5*float32(maxY)), int(0.5*float32(maxX)), int(0.73*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "A"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, questions[questionCount].Answers[0].Answer)

		if _, err := setCurrentViewOnTop(g, quiz.BoxA); err != nil {
			return err
		}
	}

	//Answer Box B
	if v, err := g.SetView(quiz.BoxB, int(0.5*float32(maxX)), int(0.5*float32(maxY)), maxX, int(0.73*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "B"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, questions[questionCount].Answers[1].Answer)

	}

	//Answer Box C
	if v, err := g.SetView(quiz.BoxC, -1, int(0.77*float32(maxY)), int(0.5*float32(maxX)), maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "C"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, questions[questionCount].Answers[2].Answer)
	}

	//Answer Box D
	if v, err := g.SetView(quiz.BoxD, int(0.5*float32(maxX)), int(0.77*float32(maxY)), maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "D"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, questions[questionCount].Answers[3].Answer)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	//Get questions
	data := readCSV()
	if len(data) == 0 {
		log.Fatal("unable to read csv file")
	}
	fmt.Println(data)
	createQuestions(data)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Turn highlighting on and set color
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//Pass in the layout I want to use
	g.SetManagerFunc(layout)

	//Quit Keybinding
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		log.Fatal(err)
	}

	//Toggle Answers Keybinding
	err = g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView)
	if err != nil {
		log.Panicln(err)
	}

	//Select an Answer
	for _, view := range boxesView {
		err = g.SetKeybinding(view, gocui.KeyEnter, gocui.ModNone, selectAnswer)
		if err != nil {
			log.Panicln(err)
		}
	}

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

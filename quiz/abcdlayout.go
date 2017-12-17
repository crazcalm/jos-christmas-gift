package quiz

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strings"
)

var (
	//BoxesView -- A slice containing the names of the A, B, C, D answer boxes
	BoxesView         = []string{BoxA, BoxB, BoxC, BoxD}
	activeView        = 0
	answersToBoxViews = make(map[string]Answer)
	answersToLetters  = map[string]string{BoxA: "A", BoxB: "B", BoxC: "C", BoxD: "D"}
)

//SetCurrentViewOnTop -- helper... should be private...
func SetCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	_, err := g.SetCurrentView(name)
	if err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)

}

//ABCDLayout -- the layout responsible for the interactive quiz, select A, B, C or D, screen
func ABCDLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView(QuestionFrame, -1, -1, maxX, int(0.5*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	if v, err := g.SetView(QuestionBox, int(0.2*float32(maxX)), int(0.1*float32(maxY)), int(0.8*float32(maxX)), int(0.4*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Question 1"
		v.Wrap = true
		fmt.Fprintln(v, Questions[QuestionCount].Question)
	}

	//Answer Box A
	if v, err := g.SetView(BoxA, -1, int(0.5*float32(maxY)), int(0.5*float32(maxX)), int(0.73*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "A"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, Questions[QuestionCount].Answers[0].Answer)

		if _, err := SetCurrentViewOnTop(g, BoxA); err != nil {
			return err
		}
	}

	//Answer Box B
	if v, err := g.SetView(BoxB, int(0.5*float32(maxX)), int(0.5*float32(maxY)), maxX, int(0.73*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "B"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, Questions[QuestionCount].Answers[1].Answer)

	}

	//Answer Box C
	if v, err := g.SetView(BoxC, -1, int(0.77*float32(maxY)), int(0.5*float32(maxX)), maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "C"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, Questions[QuestionCount].Answers[2].Answer)
	}

	//Answer Box D
	if v, err := g.SetView(BoxD, int(0.5*float32(maxX)), int(0.77*float32(maxY)), maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "D"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, Questions[QuestionCount].Answers[3].Answer)
	}

	return nil
}

func writeInfoToLayout(g *gocui.Gui, q Question) {
	//Write question
	questionBox := getQuestionBoxView(g)
	questionBox.Clear()
	questionBox.Title = fmt.Sprintf("Question %d", QuestionCount+1)
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

func getQuestionBoxView(g *gocui.Gui) *gocui.View {
	var result *gocui.View
	views := g.Views()
	for _, view := range views {
		if strings.EqualFold(QuestionBox, view.Name()) {
			result = view
		}
	}
	return result
}

func getAnswerBoxViews(g *gocui.Gui) []*gocui.View {
	var questionViews []*gocui.View
	views := g.Views()
	for _, view := range views {
		if isViewInSlice(BoxesView, view) {
			questionViews = append(questionViews, view)
		}
	}
	return questionViews
}

func isViewInSlice(viewNames []string, v *gocui.View) bool {
	result := false
	for _, name := range viewNames {
		if strings.EqualFold(name, v.Name()) {
			result = true
		}
	}
	return result
}

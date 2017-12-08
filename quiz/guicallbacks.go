package quiz

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"strings"
)

//CursorDown -- Callback used to scroll down
func CursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

//CursorUp -- Callback used to scoll up
func CursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

//Quit -- Callback used to quit application
func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

//NextView -- Callback used to interate through the A, B, C, D choices
func NextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (activeView + 1) % len(BoxesView)
	name := BoxesView[nextIndex]

	_, err := SetCurrentViewOnTop(g, name)
	if err != nil {
		log.Panicln(err)
	}

	activeView = nextIndex
	return nil
}

//SelectAnswer -- Callback used to select and answer in the ABCDLayout
func SelectAnswer(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintln(v, "Selected")

	cQuestion := currentQuestion()
	selectedAnswer := answersToBoxViews[v.Name()]

	a := UserAnswer{
		v.Name(),
		&cQuestion,
		&selectedAnswer,
	}

	//User answers
	userAnswers = append(userAnswers, a)

	views := g.Views()

	if !nextQuestionExist() {
		g.SetManagerFunc(endScreenLayout)
		err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit)
		if err != nil {
			log.Panicln(err)
			return err
		}
		err = g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, CursorDown)
		if err != nil {
			log.Panicln(err)
			return err
		}

		err = g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, CursorUp)
		if err != nil {
			log.Panicln(err)
			return err
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

		if strings.EqualFold(QuestionBox, view.Name()) {
			view.Clear() //clear the question box
			fmt.Fprintf(view, "%d - %s -- question box", i, view.Name())
		}

	}
	return nil
}

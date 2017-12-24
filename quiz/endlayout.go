package quiz

import (
	"github.com/jroimartin/gocui"
)

func endScreenLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(EndScreen, 1, 1, maxX-1, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Quiz Over"
		v.Wrap = true
		g.Cursor = true
		g.Highlight = true // frame the quiz answers

		correct, total := TotalScore(userAnswers)
		scoreTmpl, err := AnswerQuestionRatio(correct, total)
		if err != nil {
			panic(err)
		}
		err = scoreTmpl.Execute(v, nil)
		if err != nil {
			panic(err)
		}

		for index, answer := range userAnswers {
			tmpl, err := PrintSolution(index+1, answer)
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(v, answer)
			if err != nil {
				panic(err)
			}
		}

		if _, err := g.SetCurrentView(EndScreen); err != nil {
			return err
		}
	}
	return nil
}

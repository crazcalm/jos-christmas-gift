package quiz

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func endScreenLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(EndScreen, -1, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Quiz Over"
		v.Wrap = true
		g.Cursor = true

		fmt.Fprintf(v, "End of quiz\n\n")

		for index, answer := range userAnswers {
			tmpl, err := PrintSolution(index+1, answer)
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(v, answer)
		}

		if _, err := g.SetCurrentView(EndScreen); err != nil {
			return err
		}
	}
	return nil
}

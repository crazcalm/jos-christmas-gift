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
		fmt.Fprintln(v, "End of quiz")
	}
	return nil
}

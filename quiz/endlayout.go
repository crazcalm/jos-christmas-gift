package quiz

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"io/ioutil"
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
		fmt.Fprintln(v, "End of quiz")
		b, err := ioutil.ReadFile("Mark.Twain-Tom.Sawyer.txt")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(v, "%s", b)
		if _, err := g.SetCurrentView(EndScreen); err != nil {
			return err
		}
	}
	return nil
}

package main

import (
	"github.com/jroimartin/gocui"
	"log"
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	_, err := g.SetCurrentView(name)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}
	return g.SetViewOnTop(name)

}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	_, err := g.SetView("question", -1, -1, maxX, int(0.5*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	//Answer Box A
	answerA, err := g.SetView("answerA", -1, int(0.5*float32(maxY)), int(0.5*float32(maxX)), int(0.73*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	answerA.Title = "A"
	answerA.Editable = false
	answerA.Wrap = true

	//Put answer box A on top
	_, err = setCurrentViewOnTop(g, "answerA")
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	//Answer Box B
	answerB, err := g.SetView("answerB", int(0.5*float32(maxX)), int(0.5*float32(maxY)), maxX, int(0.73*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	answerB.Title = "B"
	answerB.Editable = false
	answerB.Wrap = true

	//Answer Box C
	answerC, err := g.SetView("answerC", -1, int(0.77*float32(maxY)), int(0.5*float32(maxX)), maxY)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	answerC.Title = "C"
	answerC.Editable = false
	answerC.Wrap = true

	//Answer Box D
	answerD, err := g.SetView("answerD", int(0.5*float32(maxX)), int(0.77*float32(maxY)), maxX, maxY)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	answerD.Title = "D"
	answerD.Editable = false
	answerD.Wrap = false

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Turn highlighting on and set color
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//PAss in the layout I want ot use
	g.SetManagerFunc(layout)

	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		log.Fatal(err)
	}

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

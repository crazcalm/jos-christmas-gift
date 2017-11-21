package main

import (
	"github.com/jroimartin/gocui"
	"log"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	_, err := g.SetView("question", -1, -1, maxX, int(0.5*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	//Answer Box A
	answerA, err := g.SetView("answerA", -1, int(0.5*float32(maxY)), int(0.5*float32(maxX)), int(0.75*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	answerA.Title = "A"
	answerA.Editable = false
	answerA.Wrap = true

	//Answer Box B
	answerB, err := g.SetView("answerB", int(0.5*float32(maxX)), int(0.5*float32(maxY)), maxX, int(0.75*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	answerB.Title = "B"
	answerB.Editable = false
	answerB.Wrap = true

	//Answer Box C
	answerC, err := g.SetView("answerC", -1, int(0.75*float32(maxY)), int(0.5*float32(maxX)), maxY)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	answerC.Title = "C"
	answerC.Editable = false
	answerC.Wrap = true

	//Answer Box D
	answerD, err := g.SetView("answerD", int(0.5*float32(maxX)), int(0.75*float32(maxY)), maxX, maxY)
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

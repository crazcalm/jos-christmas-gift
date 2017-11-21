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

	_, err = g.SetView("answerA", -1, int(0.5*float32(maxY)), int(0.5*float32(maxX)), int(0.75*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	_, err = g.SetView("answerB", int(0.5*float32(maxX)), int(0.5*float32(maxY)), maxX, int(0.75*float32(maxY)))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	_, err = g.SetView("answerC", -1, int(0.75*float32(maxY)), int(0.5*float32(maxX)), maxY)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

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

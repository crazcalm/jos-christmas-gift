package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

const (
	boxA          = "answerA"
	boxB          = "answerB"
	boxC          = "answerC"
	boxD          = "answerD"
	questionFrame = "questionFrame"
	questionBox   = "question"
)

var (
	boxesView  = []string{boxA, boxB, boxC, boxD}
	activeView = 0
)

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (activeView + 1) % len(boxesView)
	name := boxesView[nextIndex]

	_, err := setCurrentViewOnTop(g, name)
	if err != nil {
		log.Panicln(err)
	}

	/*
		//Debug
		out, err := g.View(boxB)
		if err != nil {
			return nil
		}
		fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)
	*/

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

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView(questionFrame, -1, -1, maxX, int(0.5*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	if v, err := g.SetView(questionBox, int(0.2*float32(maxX)), int(0.1*float32(maxY)), int(0.8*float32(maxX)), int(0.4*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Question"
		fmt.Fprintln(v, "Test question")
	}

	//Answer Box A
	if v, err := g.SetView(boxA, -1, int(0.5*float32(maxY)), int(0.5*float32(maxX)), int(0.73*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "A"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, "Answer A")

		if _, err := setCurrentViewOnTop(g, boxA); err != nil {
			return err
		}
	}

	//Answer Box B
	if v, err := g.SetView(boxB, int(0.5*float32(maxX)), int(0.5*float32(maxY)), maxX, int(0.73*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "B"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, "Answer B")

	}

	//Answer Box C
	if v, err := g.SetView(boxC, -1, int(0.77*float32(maxY)), int(0.5*float32(maxX)), maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "C"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, "Answer C")
	}

	//Answer Box D
	if v, err := g.SetView(boxD, int(0.5*float32(maxX)), int(0.77*float32(maxY)), maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "D"
		v.Editable = false
		v.Wrap = true

		fmt.Fprintln(v, "Answer D")
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

	//Turn highlighting on and set color
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//PAss in the layout I want ot use
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

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

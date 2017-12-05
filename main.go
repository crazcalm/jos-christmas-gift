package main

import (
	"fmt"
	"github.com/crazcalm/jos-christmas-gift/quiz"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	//Get questions
	data := quiz.ReadCSV()
	if len(data) == 0 {
		log.Fatal("unable to read csv file")
	}
	fmt.Println(data)
	quiz.CreateQuestions(data)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Turn highlighting on and set color
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//Pass in the layout I want to use
	g.SetManagerFunc(quiz.ABCDLayout)

	//Quit Keybinding
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quiz.Quit)
	if err != nil {
		log.Fatal(err)
	}

	//Toggle Answers Keybinding
	err = g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, quiz.NextView)
	if err != nil {
		log.Panicln(err)
	}

	//Select an Answer
	for _, view := range quiz.BoxesView {
		err = g.SetKeybinding(view, gocui.KeyEnter, gocui.ModNone, quiz.SelectAnswer)
		if err != nil {
			log.Panicln(err)
		}
	}

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

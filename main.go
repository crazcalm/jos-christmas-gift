package main

import (
	"encoding/csv"
	"fmt"
	"github.com/crazcalm/jos-christmas-gift/quiz"
	"github.com/jroimartin/gocui"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func readCSV() [][]string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
	for _, f := range files {
		fmt.Println(f.Name())
	}
	testFile := "testing.csv"
	_, err = os.Stat(testFile)
	if os.IsNotExist(err) {
		log.Fatalf("file: %s does not exist", testFile)
	}

	file, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	var records [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
		records = append(records, record)
	}

	return records
}

func createQuestions(records [][]string) {
	for i := 1; i < len(records); i++ {

		a1 := quiz.Answer{records[i][1], true}
		a2 := quiz.Answer{records[i][2], false}
		a3 := quiz.Answer{records[i][3], false}
		a4 := quiz.Answer{records[i][4], false}

		question := quiz.Question{
			records[i][0],
			[]quiz.Answer{a1, a2, a3, a4},
			records[i][5],
		}
		quiz.Questions = append(quiz.Questions, question)
	}
}

func main() {
	//Get questions
	data := readCSV()
	if len(data) == 0 {
		log.Fatal("unable to read csv file")
	}
	fmt.Println(data)
	createQuestions(data)

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

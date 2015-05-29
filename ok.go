package main

import (
	"flag"
	"os"
	"time"

	"github.com/spektroskop/ok/display"
	"github.com/spektroskop/ok/matcher"
	"github.com/spektroskop/ok/reader"
	"github.com/spektroskop/ok/util"

	"github.com/nsf/termbox-go"
)

var (
	choices  []string
	search   string
	matching bool
	matches  matcher.Matches
	selected int
	max      int
	width    int
	height   int
	editor   Editor
	spinner  Spinner
)

func main() {
	flag.Parse()

	if err := display.Init(); err != nil {
		panic(err)
	}

	spinner = NewSpinner(
		[]termbox.Attribute{
			234, 236, 238, 240, 95, 100, 242, 244, 246, 95, 100,
			100, 95, 246, 244, 242, 100, 95, 240, 238, 236, 234,
		}, "--", "~~", "::", "##", "::", "~~", "--")
	editor = NewEditor()
	eventChan := display.Run()
	readerChan := reader.Run()
	matchChan := make(chan matcher.Matches)
	doneChan := make(chan bool)

	var searchTime time.Time
	var output string

Loop:
	for {
		width, height = display.Reset()
		max = util.Limit(len(matches), 0, height-1)
		selected = util.Limit(selected, 0, max-1)

		Draw(editor.Text, width, max, selected)

		termbox.Flush()

		select {
		case <-util.MaybeAfter(time.Millisecond*100, matching):
			util.Debug("Progress")

		case matches = <-matchChan:
			util.Debugf("Matches %d (%v)\n", len(matches), time.Since(searchTime))
			matching = false

		case entries := <-reader.MaybeChan(readerChan, !matching):
			choices = append(choices, entries...)
			util.Debug("Choices", len(choices))

		case event := <-eventChan:
			switch event.Type {
			case termbox.EventKey:
				switch event.Key {
				case termbox.KeyEsc, termbox.KeyCtrlC:
					break Loop
				case termbox.KeyEnter:
					output = editor.Text
					break Loop

				case termbox.KeyArrowDown, termbox.KeyCtrlN:
					selected += 1
				case termbox.KeyArrowUp, termbox.KeyCtrlP:
					selected -= 1

				case termbox.KeyCtrlW:
					if editor.RemoveWord() {
						goto Search
					}
				case termbox.KeyBackspace, termbox.KeyBackspace2:
					if editor.RemoveBackwards() {
						goto Search
					}
				case termbox.KeyCtrlK:
					if editor.RemoveToEnd() {
						goto Search
					}
				case termbox.KeyCtrlU:
					if editor.RemoveToBeginning() {
						goto Search
					}

				case termbox.KeyArrowRight:
					editor.MoveForward()
				case termbox.KeyArrowLeft:
					editor.MoveBackward()
				case termbox.KeyCtrlA:
					editor.MoveStart()
				case termbox.KeyCtrlE:
					editor.MoveEnd()

				case termbox.KeyTab:
					if len(matches) > 0 {
						editor.Update(matches[selected].Text, -1)
						goto Search
					}

				default:
					if event.Ch != 0 {
						editor.Insert(event.Ch)
						goto Search
					}
				}
			}
		}

		continue

	Search:
		util.Debugf("Search `%s'\n", editor.Text)

		close(doneChan)
		doneChan = make(chan bool)
		searchTime = time.Now()
		matching = true

		go matcher.Run(editor.Text, choices, matchChan, doneChan)
	}

	termbox.Close()

	if len(output) != 0 {
		Output(output, os.Args[1:])
	}
}

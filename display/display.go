package display

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func Init() (err error) {
	if err := termbox.Init(); err != nil {
		return err
	}
	termbox.SetOutputMode(termbox.Output256)
	return
}

func makeattrs(attrs ...termbox.Attribute) (a termbox.Attribute, b termbox.Attribute) {
	a, b = termbox.ColorDefault, termbox.ColorDefault
	switch len(attrs) {
	case 2:
		return attrs[0], attrs[1]
	case 1:
		return attrs[0], b
	}
	return
}

func Print(x, y int, s string, attrs ...termbox.Attribute) {
	fg, bg := makeattrs(attrs...)
	var i int
	for _, r := range s {
		termbox.SetCell(x+i, y, r, fg, bg)
		i++
	}
}

func Printf(x, y int, s string, attrs []termbox.Attribute, args ...interface{}) {
	s = fmt.Sprintf(s, args...)
	Print(x, y, s, attrs...)
}

func Reset(attrs ...termbox.Attribute) (int, int) {
	fg, bg := makeattrs(attrs...)
	termbox.Clear(fg, bg)
	return termbox.Size()
}

func Run() chan termbox.Event {
	eventChan := make(chan termbox.Event)
	go func() {
		for {
			eventChan <- termbox.PollEvent()
		}
	}()
	return eventChan
}

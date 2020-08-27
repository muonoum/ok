package main

import (
	"fmt"

	"github.com/spektroskop/ok/display"

	"github.com/dustin/go-humanize"
	"github.com/nsf/termbox-go"
)

func Draw(search string, width, height, selected int) {
	prompt, style := *prompt, 111|termbox.AttrBold

	if !matching {
		display.Print(0, 0, prompt, style)
		display.Print(len(prompt)+1, 0, search)
	} else {
		prompt, style := spinner.Next()
		display.Print(0, 0, prompt, style)
		display.Print(len(prompt)+1, 0, search, 247)
	}

	termbox.SetCursor(len(prompt)+editor.Cursor+1, 0)

	numbers := fmt.Sprintf("%s/%s",
		humanize.Comma(int64(len(matches))),
		humanize.Comma(int64(len(choices))),
	)

	display.Print(width-len(numbers), 0, numbers, 239)

	for y, match := range matches[:max] {
		if y == selected {
			display.Print(0, y+1, "**", 183|termbox.AttrBold)
		} else {
			display.Print(0, y+1, "--", 237)
		}

		var index int

		x := 3
		for _, char := range match.Text {
			if x+index >= width-3 {
				display.Print(x+index+1, y+1, "»", 237)
				break
			}

			if match.Matched[index] {
				display.Print(x+index, y+1, string(char), 81|termbox.AttrUnderline|termbox.AttrBold)
			} else if index > match.Start && index < match.End {
				display.Print(x+index, y+1, string(char), 74)
			} else {
				display.Print(x+index, y+1, string(char), 162)
			}

			index++
		}
	}
}

package reader

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

func Run() chan []string {
	entryChan := make(chan []string)

	go func() {
		var entries []string

		if isatty.IsTerminal(os.Stdin.Fd()) {
			return
		}

		reader := bufio.NewReader(os.Stdin)

		for {
			entry, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}

			entries = append(entries, strings.TrimSpace(entry))

			select {
			case entryChan <- entries:
				entries = []string{}
			default:
			}
		}

		entryChan <- entries
	}()

	return entryChan
}

func MaybeChan(ch <-chan []string, maybe bool) <-chan []string {
	if maybe {
		return ch
	}

	return nil
}

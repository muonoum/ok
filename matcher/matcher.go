package matcher

import (
	"sort"
	"strings"
	"time"

	"github.com/spektroskop/ok/util"
)

type Entry struct {
	Text    string
	Score   float64
	Start   int
	End     int
	Matched map[int]bool
}

type Matches []Entry

func (m Matches) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Matches) Len() int           { return len(m) }
func (m Matches) Less(i, j int) bool { return m[j].Score < m[i].Score }

func Score(search, choice string) (score float64, matched map[int]bool, start, end int) {
	matched = make(map[int]bool)

	for _, r := range search {
		if index := strings.IndexRune(choice[end:], r); index == -1 {
			score = 0

			return
		} else {
			end += index
			matched[end] = true
			score += 1

			if len(matched) == 1 {
				start = end
			}

			end += 1
		}
	}

	score += 1.0 / float64(end-start+1)

	return
}

func Run(search string, choices []string, matchChan chan<- Matches, doneChan <-chan bool) {
	var matches Matches

	if len(search) == 0 {
		for _, choice := range choices {
			matches = append(matches, Entry{choice, 1.0, 0, 0, make(map[int]bool)})
		}

		goto End
	}

	for _, choice := range choices {
		select {
		case <-doneChan:
			util.Debugf("Cancel `%s'\n", search)
			return
		default:
			if score, matched, start, end := Score(strings.ToLower(search), strings.ToLower(choice)); score > 0 {
				matches = append(matches, Entry{choice, score, start, end, matched})
			}
		}
	}

	sort.Sort(matches)

End:
	select {
	case matchChan <- matches:
	case <-time.After(time.Millisecond * 100):
		panic("Timeout `matchChan <- matches'")
	}
}

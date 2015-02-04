package main

import "github.com/nsf/termbox-go"

type Spinner struct {
	Parts      []string
	Attributes []termbox.Attribute
	IndexP     int
	IndexA     int
}

func (s *Spinner) Next() (string, termbox.Attribute) {
	if s.IndexP += 1; s.IndexP >= len(s.Parts) {
		s.IndexP = 0
	}
	if s.IndexA += 1; s.IndexA >= len(s.Attributes) {
		s.IndexA = 0
	}
	return s.Parts[s.IndexP], s.Attributes[s.IndexA]
}

func NewSpinner(a []termbox.Attribute, s ...string) Spinner {
	return Spinner{Parts: s, Attributes: a}
}

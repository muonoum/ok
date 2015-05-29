package main

import "regexp"

type Editor struct {
	Text   string
	Cursor int
	Word   *regexp.Regexp
}

func NewEditor() Editor {
	return Editor{Word: regexp.MustCompile(`\s+|-|\.|,`)}
}

func (editor *Editor) State() (string, int) {
	return editor.Text, editor.Cursor
}

func (editor *Editor) Clear() bool {
	if len(editor.Text) == 0 {
		return false
	}
	editor.Update("", 0)
	return true
}

func (editor *Editor) Update(text string, cursor int) {
	editor.Text, editor.Cursor = text, cursor
	if editor.Cursor < 0 {
		editor.Cursor = len(editor.Text)
	}
}

func (editor *Editor) Insert(r rune) {
	s, c := editor.State()
	editor.Update(s[:c]+string(r)+s[c:], c+1)
}

func (editor *Editor) Remove() bool {
	s, c := editor.State()
	if len(s) == 0 || c < len(s) {
		return false
	}
	editor.Update(s[:c]+s[c+1:len(s)], c)
	return true
}

func (editor *Editor) RemoveBackwards() bool {
	s, c := editor.State()
	if c == 0 || len(s) == 0 {
		return false
	}
	editor.Update(s[:c-1]+s[c:len(s)], c-1)
	return true
}

func (editor *Editor) RemoveToBeginning() bool {
	s, c := editor.State()
	if len(s) == 0 {
		return false
	}
	editor.Update(s[c:], 0)
	return true
}

func (editor *Editor) RemoveToEnd() bool {
	s, c := editor.State()
	if len(s) == 0 || c == len(s) {
		return false
	}
	editor.Update(s[:c], c)
	return true
}

func (editor *Editor) RemoveWord() bool {
	s, c := editor.State()
	if len(s) == 0 || c == 0 {
		return false
	}

	text := s[:c]
	word := editor.Word.FindAllStringIndex(text, -1)
	if word == nil {
		return editor.Clear()
	}
	end := word[len(word)-1]

	if len(word) > 1 && c <= end[1] {
		text = text[:word[len(word)-2][1]]
	} else if c <= end[1] {
		text = ""
	} else {
		text = text[:end[1]]
	}

	editor.Update(text+s[c:], len(text))

	return true
}

func (editor *Editor) MoveStart() {
	editor.Cursor = 0
}

func (editor *Editor) MoveEnd() {
	editor.Cursor = len(editor.Text)
}

func (editor *Editor) MoveForward() {
	if editor.Cursor < len(editor.Text) {
		editor.Cursor += 1
	}
}

func (editor *Editor) MoveBackward() {
	if editor.Cursor > 0 {
		editor.Cursor -= 1
	}
}

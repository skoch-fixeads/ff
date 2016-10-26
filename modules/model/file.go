package model

import (
	"os"
	"bufio"
	"regexp"
	"strings"
)

type Line struct {
	Line string
	Text string
	Num  int
	Confirmed map[int]string
}

type Entity struct {
	Path            string
	Error           error
	NumLines        int
	Output          []*Line
	IsDir           bool
	CaseInsensitive bool
	searchText      string
}

func NewEntity(path string, isDir bool) *Entity {
	return &Entity{
		Path:   path,
		IsDir:  isDir,
		Output: make([]*Line, 0),
	}
}

func (e *Entity) GetType() string {
	if e.IsDir {
		return "dir "
	}

	return "file"
}

func (e *Entity) Write(line *Line) {
	e.Output = append(e.Output, line)
}

func (e *Entity) FindByRegex(regex *regexp.Regexp) error {
	return e.readLines(func(line string) {
		words := regex.FindAllString(line, -1)
		if len(words) > 0 {
			for _, v := range words {
				e.Write(&Line{Line: line, Text: v, Num: e.NumLines})
			}
		}
	})
}

func (e *Entity) FindByText(text string) error {
	e.searchText = text
	return e.readLines(func(line string) {
		l, t := line, text
		if !e.CaseInsensitive {
			l, t = strings.ToLower(line), strings.ToLower(text)
		}

		if strings.Contains(l, t) {
			e.Write(&Line{Line: l, Text: t, Num: e.NumLines})
		}
	})
}

func (e *Entity) readLines(cb func(string)) error {
	file, err := os.Open(e.Path)
	if err != nil {
		e.Error = err
		return err
	}
	defer file.Close()

	e.NumLines = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		e.NumLines++
		cb(line)
	}
	return nil
}
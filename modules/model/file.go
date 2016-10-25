package model

import (
	"os"
	"bufio"
	"strings"
	"regexp"
	"fmt"
)

type Entity struct {
	Path            string
	Error           error
	NumLines        int
	Output          map[int]string
	Buffer          map[int]string

	CaseInsensitive bool
	searchText      string
}

func NewEntity(path string) *Entity {
	return &Entity{
		Path: path,
		Output: make(map[int]string, 0),
		Buffer: make(map[int]string, 0),
	}
}

func (e Entity) Write(i int, text string) {
	e.Output[i] = text
}

func (e Entity) FindByRegex(regex *regexp.Regexp) error {
	var cb = func(line string) bool {
		words := regex.FindAllString(line, -1)
		if len(words) > 0 {
			return true
		}
		return false
	}

	return e.readLines(cb)
}

func (e Entity) FindByText(text string) error {
	var cb = func(line string) bool {
		if (e.CaseInsensitive) {
			return strings.Contains(line, text)
		} else {
			return strings.Contains(strings.ToLower(line), strings.ToLower(text))
		}
	}

	e.searchText = text
	return e.readLines(cb)
}

func (e Entity) readLines(cb func(string) bool) error {
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

		if cb(line) {
			e.Write(e.NumLines, line)
		}

		e.Buffer[e.NumLines] = line
	}

	return nil
}

func (e Entity) ReplaceLine(lineNumber int, text string) error {
	if _, has := e.Buffer[lineNumber]; has {
		fmt.Printf("\t [%v -> %v] \n", e.searchText, text)
		return nil
	}

	return fmt.Errorf("Line %v not found!", lineNumber)
}
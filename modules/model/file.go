package model

import (
	"os"
	"strings"
	"bufio"
)

type Entity struct {
	Path string
	Error error
	NumLines int
	Output map[int]string
}

func NewEntity(path string) *Entity {
	return &Entity{Path: path, Output: make(map[int]string, 0)}
}

func (e Entity) Write(i int, text string) {
	e.Output[i] = text
}

func (e Entity) SearchText(text string) error {
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

		if strings.Contains(line, text) {
			e.Write(e.NumLines, line)
		}
	}

	return nil
}
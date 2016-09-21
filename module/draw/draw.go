package draw

import (
	"fmt"
	"bytes"
	"strings"
)

var buff bytes.Buffer


func PWrite(args ...string)  {
	Write(args...)
	Print()
}

func Write(args ...string)  {
	for _, t := range args {
		buff.WriteString(fmt.Sprintf("%s ", t))
	}
}

func PWriteLn(args ...string)  {
	WriteLn(args...)
	Print()
}

func WriteLn(args ...string)  {
	Write(args...)
	buff.WriteString("\n")
}

func Print() {
	fmt.Printf("%v", buff.String())
}

func Clear() {
	fmt.Printf("\r%s\r", strings.Repeat(" ", 50))
}
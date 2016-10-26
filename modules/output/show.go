package output

import (
	"os"
	"fmt"
	"bufio"
	"strings"

	"github.com/fatih/color"
	"github.com/rodkranz/ff/modules/model"
	"github.com/rodkranz/ff/modules/settings"
)


var (
	w = map[string]func(a ...interface{}) string {}
	r = []*model.Line{}
)

func init() {
	w["white"] = color.New(color.FgWhite).SprintFunc()
	w["yellow"] = color.New(color.FgYellow).SprintFunc()
	w["green"] = color.New(color.FgGreen).SprintFunc()
	w["blue"] = color.New(color.FgBlue).SprintFunc()
	w["hiWrite"] = color.New(color.FgBlack, color.BgGreen).SprintFunc()
}

func PrintLn(e *model.Entity) {
	if len(settings.SearchText) != 0 && len(e.Output) == 0 {
		return
	}

	var extra string = ""
	if len(e.Output) != 0 {
		extra = fmt.Sprintf("(lines: %s)", w["blue"](len(e.Output)))
	}

	fmt.Printf("[%s] %v %v\n",  w["white"](e.GetType()), w["yellow"](e.Path), extra)
	if len(e.Output) > 0 {
		for _, s := range e.Output {
			if len(settings.Replace) != 0 {
				askConfirmation(e, s)
			} else {
				fmt.Printf("[%v]:\t %v\n", w["blue"](s.Num), highlightWord(s.Text, s.Line))
			}
		}
		fmt.Println()
	}
}

func highlightWord(word, line string) string {
	spitted := strings.Split(line, word)
	if len(spitted) == 0 {
		return line
	}
	return strings.Join(spitted, w["hiWrite"](word))
}

func askConfirmation(e *model.Entity, s *model.Line) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("[%v]:\t %v (y/n)?", w["blue"](s.Num), highlightWord(s.Text, s.Line));
		variable, _ := reader.ReadString('\n')

		if len(variable) > 0 {
			variable = strings.ToLower(variable[0:1])
		}

		switch variable {
		case "n" :
			return false
		case "y" :
			return true
		}
	}
}
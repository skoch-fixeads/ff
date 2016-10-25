package output

import (
	"fmt"
	"github.com/rodkranz/ff/modules/model"
	"github.com/rodkranz/ff/modules/settings"
)

func PrintLn(e *model.Entity) {
	if len(settings.SearchText) == 0 {
		fmt.Printf("%v\n", e.Path)
	}

	if len(e.Output) != 0 {
		fmt.Printf("%v (%d)\n", e.Path, len(e.Output))
		for i, s := range e.Output {
			if len(settings.Replace) != 0 {
				fmt.Printf("[%v]:\t %v (y/n)?", i, s)
				e.ReplaceLine(i, settings.Replace)
			} else {
				fmt.Printf("[%v]:\t %v\n", i, s)
			}
		}
	}

}
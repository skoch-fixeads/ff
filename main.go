package main

import (
	"runtime"

	_ "github.com/rodkranz/ff/autoload"

	"github.com/rodkranz/ff/modules/helpers"
	"github.com/rodkranz/ff/modules/model"
	"github.com/rodkranz/ff/modules/search"
	"github.com/rodkranz/ff/modules/settings"
	"github.com/rodkranz/ff/modules/output"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	settings.Directory = settings.Directory
	//settings.SearchText = "rodrigo"

	search.Walk(func(e *model.Entity) error {
		for _, help := range helpers.Helpers {
			if err := help.Run(e); err != nil {
				break;
			}
		}

		output.PrintLn(e)
		return nil
	})

}

package main

import (
	"runtime"

	"github.com/rodkranz/ff/modules/model"
	"github.com/rodkranz/ff/modules/output"
	"github.com/rodkranz/ff/modules/search"
	"github.com/rodkranz/ff/modules/settings"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	settings.Directory = settings.Directory
	settings.SearchText = "rodrigo"
	//settings.CaseInsensitive = false
	settings.Reach = 10
	//settings.Regexp = regexp.MustCompile("(rodkranz|rodrigo)")
	//settings.Replace = "digo"

	search.Walk(func(e *model.Entity) error {
		e.CaseInsensitive = settings.CaseInsensitive
		if len(settings.SearchText) != 0 {
			if err := e.FindByText(settings.SearchText); err != nil {
				return err
			}
		} else if settings.Regexp != nil {
			if err := e.FindByRegex(settings.Regexp); err != nil {
				return err
			}
		}

		output.PrintLn(e)
		return nil
	})
}

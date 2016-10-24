package text

import (
	"github.com/rodkranz/ff/modules/helpers"
	"github.com/rodkranz/ff/modules/model"
	"github.com/rodkranz/ff/modules/settings"
)

type Find struct {}

func init() {
	helpers.RegisterHelp(&Find{})
}

func (t Find) Run(e *model.Entity) error {
	if len(settings.SearchText) == 0 {
		return nil
	}

	return e.SearchText(settings.SearchText)
}



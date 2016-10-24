package replace

import (
	"github.com/rodkranz/ff/modules/helpers"
	"github.com/rodkranz/ff/modules/model"
)

type Replace struct {}

func init() {
	helpers.RegisterHelp(&Replace{})
}

func (t Replace) Run(e *model.Entity) error {
	return nil
}


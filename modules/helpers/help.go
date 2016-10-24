package helpers

import (
	"github.com/rodkranz/ff/modules/model"
)

type Helper interface {
	Run (*model.Entity) error
}

var Helpers = []Helper{}

func RegisterHelp(h Helper) {
	Helpers = append(Helpers, h)
}

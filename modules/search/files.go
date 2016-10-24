package search

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rodkranz/ff/modules/model"
	"github.com/rodkranz/ff/modules/tools"
	"github.com/rodkranz/ff/modules/settings"
	"sort"
)

func IsSkipDir(info os.FileInfo) bool {
	if info.Name() == "." {
		return false
	}

	if info.IsDir() {
		i := sort.SearchStrings(settings.ExcludeExtension, info.Name())
		return (len(settings.ExcludeExtension) != i)
	}

	return false
}

func IsSkipFile(info os.FileInfo) bool {
	if len(settings.FileName) != 0 && !strings.Contains(info.Name(), settings.FileName) {
		return true
	}

	if len(settings.Extension) > 0 && !tools.IsExistsInArray(filepath.Ext(info.Name()), settings.Extension) {
		return true
	}

	return false
}

func Walk(cb func(e *model.Entity) error) error {
	WalkFunc := func(path string, info os.FileInfo, err error) error {
		if IsSkipDir(info) {
			return filepath.SkipDir
		}

		if IsSkipFile(info) {
			return nil
		}

		return cb(model.NewEntity(path))
	}

	return filepath.Walk(settings.Directory, WalkFunc)
}

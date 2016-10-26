package settings

import "regexp"

var (
	Reach            int
	FileName         string
	SearchText       string
	Directory        string
	Extension        []string
	CaseInsensitive  bool
	ExcludeExtension []string
	Regexp           *regexp.Regexp
	Replace          string
)

func init() {
	Reach = 10
	Replace = ""
	FileName = ""
	Directory = "./"
	SearchText = ""
	Extension = []string{}
	CaseInsensitive = true
	ExcludeExtension = []string{".bzr", "CVS", ".git", ".hg", ".svn"}
}

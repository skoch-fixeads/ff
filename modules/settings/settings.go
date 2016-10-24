package settings

var (
	FileName string
	SearchText string
	Directory string
	Extension []string
	ExcludeExtension []string
)

func init() {
	FileName = ""
	SearchText = ""
	Directory = "./"
	Extension = []string{}
	ExcludeExtension = []string{".bzr","CVS",".git",".hg",".svn"}
}

package setting

import (
	"flag"
	"runtime"
	"fmt"
	"strings"
	"regexp"
)

var (
	Name string
	Description string
	Ver float64

	// General
	CPUNum int
	NoColor bool

	// Search Config
	Directory string
	FileFilter string
	TextSearch string
	RegexFilter string
	Regex *regexp.Regexp
	CaseInsensitive bool
	IgnoreFolder []string
	Reach int

	// Replace
	IsReplace bool = false
	ForceReplaceAll bool
	ReplaceWithText string
)

func init() {
	// General
	flag.IntVar(&CPUNum, "cpu", runtime.NumCPU(), fmt.Sprintf("Number of CPU you have %d available", runtime.NumCPU()))
	flag.BoolVar(&NoColor, "-no-color", false, "Disable color output")

	// Search
	flag.StringVar(&TextSearch, "t", "", "Search text")
	flag.StringVar(&FileFilter, "f", "", "Filter by file name")
	flag.StringVar(&Directory, "d", "./", "Directory ffConfig")
	flag.BoolVar(&CaseInsensitive, "i", false, "Search text case insensitive")
	flag.IntVar(&Reach, "r", 10, "Range around of the word")
	flag.StringVar(&RegexFilter, "rg", "", "Search by Regex")
	exclude := *flag.String("-exclude-dir", ".bzr,CVS,.git,.hg,.svn", "Exclude dir from reader")


	// Replace
	flag.StringVar(&ReplaceWithText, "replace", "", "Replace result to text")
	flag.BoolVar(&ForceReplaceAll, "force", false, "Replace all result without ask.")

	// Parse Flags
	flag.Parse()

	// Regex
	if len(RegexFilter) > 0 {
		TextSearch = ""
		Regex = regexp.MustCompile(TextSearch)
	}

	// Ignore
	if len(exclude) > 0 {
		IgnoreFolder = strings.Split(exclude, ",")
	}

	// Replace
	if len(ReplaceWithText) > 0 {
		IsReplace = true
	}

	// CPU
	if CPUNum > runtime.NumCPU() {
		CPUNum = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(CPUNum)

	// Check if has parameters
	if flag.NArg() == 1 {
		if par := flag.Arg(0); par[0] != 45 {
			TextSearch = flag.Arg(0)
		}
	}
}
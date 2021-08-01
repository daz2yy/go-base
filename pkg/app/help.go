package app

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/pflag"
)

const (
	flagHelp          = "Help"
	flagHelpShorthand = "H"
)

func addHelpCommandFlag(usage string, fs *pflag.FlagSet) {
	fs.BoolP(
		flagHelp,
		flagHelpShorthand,
		false,
		fmt.Sprintf("Help for the %s command.", color.GreenString(strings.Split(usage, " ")[0])),
	)
}

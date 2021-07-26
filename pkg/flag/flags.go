package flag

import (
	goflag "flag"
	"strings"

	"github.com/spf13/pflag"

	"github.com/daz2yy/go-base/pkg/log"
)

// WordSepNormalizeFunc changes all flags from "_" to "-" separators.
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}

	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc changes and print warn all flags from "_" to "-" separators.
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		newname := strings.ReplaceAll(name, "_", "-")
		log.Warnf("%s is DEPRECATED and will be removed in a future version. Use %s instead.", name, newname)

		return pflag.NormalizedName(newname)
	}

	return pflag.NormalizedName(name)
}

// InitFlags Normalizes, parse, then logs the command line flags.
func InitFlags() {
	pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

// PrintFlags logs the flags in the flagset.
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Infof("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}

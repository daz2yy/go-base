package cmd

import (
	"flag"
	"io"
	"os"

	cliflag "github.com/daz2yy/go-base/pkg/flag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/daz2yy/go-base/internal/ctl/util/templates"
	genericapiserver "github.com/daz2yy/go-base/internal/pkg/server"
)

// NewDefaultCtlCommand creates the `ctl` command with default arguments.
func NewDefaultCtlCommand() *cobra.Command {
	return NewCtlCommand(os.Stdin, os.Stdout, os.Stderr)
}

// NewCtlCommand returns new initialized instance of 'ctl' root command.
func NewCtlCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "ctl",
		Short: "ctl controls the platform",
		Long: templates.LongDesc(`
			ctl controls the platform, is the client side tool for this platform.
			
			Find more information at:
				https://github.com/daz2yy/go-base`),
		Run: runHelp,
		// Hook before and after run initialize and write profiles to disk, respectively
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return initProfiling()
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			return flushProfiling()
		},
	}
	flags := cmds.PersistentFlags()
	flags.SetNormalizeFunc(cliflag.WarnWordSepNormalizeFunc) // Warn for "_" flags

	// Normalize all flags that are coming from other packages or pre-configurations
	// a.k.a. change all "_" to "-". e.g. glog package
	flags.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	// profiling flag setting.
	addProfilingFlags(flags)

	// TODO: add generic cli options, such as username, password, tls, etc.
	// configFlags := genericclioptions.

	_ = viper.BindPFlags(cmds.PersistentFlags())
	cobra.OnInitialize(func() {
		genericapiserver.LoadConfig(viper.GetString("ctlconfig"), "config")
	})
	cmds.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	// From this point and forward we get warnings on flags that contain "_" separators
	cmds.SetGlobalNormalizationFunc(cliflag.WarnWordSepNormalizeFunc)

	// CommandGroups

	// cmds.AddCommand(version.)

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bilbilak/godini/internal"
)

var setCmd = &cobra.Command{
	Use:   "set [flags] <FILE> [SETTING]... [STDIN]",
	Short: "Set settings in config file",
	Long:  `Add, enable (uncomment), or modify settings in the INI configuration file`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	setCmd.Flags().BoolVarP(&internal.All, "all", "a", false, "Apply to all sections")
	setCmd.Flags().StringVarP(&internal.Section, "section", "s", "", "Apply to the given section")

	rootCmd.AddCommand(setCmd)
}

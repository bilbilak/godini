package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bilbilak/godini/internal"
)

var unsetCmd = &cobra.Command{
	Use:   "unset [flags] <FILE> [SETTING]... [STDIN]",
	Short: "Unset settings in config file",
	Long:  `Disable/Reset (comment-out) settings in the INI configuration file`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		internal.Process(internal.UNSET, args)
	},
}

func init() {
	unsetCmd.Flags().BoolVarP(&internal.All, "all", "a", false, "Apply to all sections")
	unsetCmd.Flags().StringVarP(&internal.Section, "section", "s", "", "Apply to the given section")

	rootCmd.AddCommand(unsetCmd)
}

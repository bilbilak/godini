package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bilbilak/godini/internal"
)

var getCmd = &cobra.Command{
	Use:   "get [flags] <FILE> [SETTING]... [STDIN]",
	Short: "Get settings from config file",
	Long:  `Retrieve settings from the INI configuration file`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	getCmd.Flags().BoolVarP(&internal.All, "all", "a", false, "Apply to all sections")
	getCmd.Flags().StringVarP(&internal.Section, "section", "s", "", "Apply to the given section")
	getCmd.Flags().BoolVarP(&internal.Full, "full", "f", false, "Show fully qualified keys")

	rootCmd.AddCommand(getCmd)
}

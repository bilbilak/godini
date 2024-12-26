package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bilbilak/godini/internal"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [flags] <FILE> [SETTING]... [STDIN]",
	Short: "Delete settings from config file",
	Long:  `Delete settings from the INI configuration file`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&internal.All, "all", "a", false, "Apply to all sections")
	deleteCmd.Flags().StringVarP(&internal.Section, "section", "s", "", "Apply to the given section")

	rootCmd.AddCommand(deleteCmd)
}

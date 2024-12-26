package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	app "github.com/bilbilak/godini/config"
	"github.com/bilbilak/godini/internal"
)

var (
	Help    bool
	Version bool
	License bool
)

var rootCmd = &cobra.Command{
	Use:   strings.ToLower(app.Name),
	Short: "INI Configuration Management Tool",
	Long:  app.Name + ` is a powerful CLI tool for manipulation of configuration files in INI format.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if Help {
			_ = cmd.Help()
			return
		}

		if Version {
			fmt.Println(app.Version)
			return
		}

		if License {
			fmt.Println(app.License)
			return
		}

		internal.Help()
	},
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.PersistentFlags().BoolVar(&Help, "help", false, "Show usage instructions")

	rootCmd.Flags().BoolVar(&Version, "version", false, "Display the installed version number")
	rootCmd.Flags().BoolVar(&License, "license", false, "Display the license name")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		internal.FatalError(err)
	}
}

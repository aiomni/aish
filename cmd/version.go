package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of AISH",
	Long:  `All software has versions. This is AISH's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

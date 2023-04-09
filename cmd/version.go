package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wcli",
	Long:  `All software has versions. This is wcli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wcli v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}

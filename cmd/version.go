package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of API Server",
	Long:  `All software has versions. This is API Server's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Melo CMDB API Server v2.0")
	},
}

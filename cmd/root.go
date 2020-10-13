package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cmcs",
	Short: "CMCS",
	Long: `More of CMCS`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmcs was ran")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

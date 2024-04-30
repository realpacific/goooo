package cmd

import (
	"github.com/spf13/cobra"
	. "goooo/logging"
)

var RootCmd = &cobra.Command{
	Use:   "goooo",
	Short: "Demo application",
	Long:  `Demo application`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		Logger.Fatalln("error executing command %w", err)
	}
}

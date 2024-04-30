package cmd

import (
	"github.com/spf13/cobra"
	"goooo/api"
	. "goooo/config"
	. "goooo/logging"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server",
	Long:  `Starts a http server`,
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("serving at %v", GetPort())
		api.Serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

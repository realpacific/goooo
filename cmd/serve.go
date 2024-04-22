package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	. "goooo/initializer"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server",
	Long:  `Starts a http server`,
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("serving at %v", viper.GetInt("port"))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

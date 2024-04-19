package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server",
	Long:  `Starts a http server`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("serving at %v", viper.GetInt("port"))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

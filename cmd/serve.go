package cmd

import (
	"github.com/spf13/cobra"
	"goooo/api"
	"goooo/bloc"
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

var listEmployeesCmd = &cobra.Command{
	Use:   "employees",
	Short: "lists employees",
	Long:  `Lists  employees`,
	Run: func(cmd *cobra.Command, args []string) {
		bloc.ListEmployees(&bloc.EmployeeParams{})
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(listEmployeesCmd)
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	server2 "github.com/retatu/arq-hexagonal/adapters/web/server"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.NewWebServer()
		server.Service = &productService
		fmt.Println("Webserver has been started")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

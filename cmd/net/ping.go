/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	urlPathString string
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//logic
		ping(urlPathString)

	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPathString, "url", "u", "", "The url top ping")

	if err := pingCmd.MarkFlagRequired(("url")); err != nil {
		fmt.Println(err)
	}
	NetCmd.AddCommand(pingCmd)
}

func ping(domain string) (int, error) {
	response, err := http.Get(domain)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return response.StatusCode, nil
}

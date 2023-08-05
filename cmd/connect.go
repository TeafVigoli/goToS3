/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var endpoint string
var port string

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect and list basic stats for a Kafka endpoint",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connect called")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringVar(&endpoint, "endpoint", "", "kafka connection endpoint")
	_, err := connectCmd.MarkFlagRequired("endpoint")
	
	connectCmd.Flags().StringVar(&port, "port", "", "kafka connection endpoint")
	_, err = connectCmd.MarkFlagRequired("port")

	if err != nil {
        return fmt.Errorf("enpoint or port not set: %w", err)
    }
}

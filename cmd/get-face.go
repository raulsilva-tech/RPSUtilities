/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/raulsilva-tech/RPSUtilities/internal/usecase"
	"github.com/spf13/cobra"
)

// getFaceCmd represents the getFace command
var getFaceCmd = &cobra.Command{
	Use:   "get-face",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		url, _ := cmd.Flags().GetString("url")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		timeout, _ := cmd.Flags().GetInt("timeout")

		if user == "" || password == "" || host == "" || port == 0 || timeout == 0 {
			fmt.Println("Error: Missing required flags (user, password, timeout, host)")
			return
		}

		uc := usecase.NewFBIGetFaceUseCase()
		output, _ := uc.Execute(host, port, user, password, url, timeout)
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// 	return
		// }
		// b, _ := json.MarshalIndent(output, "", "  ")
		b, _ := json.Marshal(output)
		fmt.Println(string(b))

	},
}

func init() {
	fbiCmd.AddCommand(getFaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getFaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getFaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

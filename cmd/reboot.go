/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/raulsilva-tech/RPSUtilities/internal/usecase"
	"github.com/spf13/cobra"
)

// rebootCmd represents the reboot command
var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reinicia a camera",
	Long:  `Envia o comando SystemReboot através do protocolo ONVIF`,
	RunE:  rebootFunc(),
}

func rebootFunc() RunEFun {
	return func(cmd *cobra.Command, args []string) error {

		ip, _ := cmd.Flags().GetString("ip")
		onvifPort, _ := cmd.Flags().GetInt("port")
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")

		//obligatory data were informed?
		if ip != "" && onvifPort != 0 && user != "" { // && password != "" {

			ucReboot := usecase.NewRebootUseCase()
			err := ucReboot.Execute(ip, onvifPort, user, password)
			if err == nil {
				fmt.Println("success")
			}
			return err

		} else {
			return ErrInsufficientData
		}
	}
}

func init() {
	camCmd.AddCommand(rebootCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rebootCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rebootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ErrInsufficientData = errors.New("error: insufficient data")
)

// camCmd represents the cam command
var camCmd = &cobra.Command{
	Use:   "cam",
	Short: "Comandos para comunicação com uma câmera IP.",
	Long: `Comandos disponíveis:
	- Reboot
	- Check
	- Stream
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("ip")
		fmt.Printf("cam called. IP: %s", ip)
	},
}

func init() {
	rootCmd.AddCommand(camCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	camCmd.PersistentFlags().String("ip", "", "IP da câmera")
	camCmd.PersistentFlags().IntP("port", "p", 8899, "Porta ONVIF da câmera")
	camCmd.PersistentFlags().String("user", "", "Usuário da câmera")
	camCmd.PersistentFlags().String("password", "", "Senha da câmera")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// camCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

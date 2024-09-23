/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/raulsilva-tech/RPSUtilities/internal/usecase"
	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Checa se stream da câmera está acessível",
	Long:  ``,
	RunE:  streamFunc(),
}

func streamFunc() RunEFun {
	return func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Flags().GetString("url")

		if url == "" {
			return ErrInsufficientData
		}

		ucStream := usecase.NewStreamUseCase()
		err := ucStream.Execute(url)
		if err == nil {
			fmt.Println("success")
		}
		return err
	}
}

func init() {
	camCmd.AddCommand(streamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// streamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	streamCmd.Flags().String("url", "", "URL stream da câmera")
	streamCmd.MarkFlagRequired("url")

}

package main

import (
	"fmt"
	"os"

	"github.com/m3-chat/infra/internal/deploy"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "infra",
		Short: "Infra CLI for deploying m3-chat services",
	}

	var deployCmd = &cobra.Command{
		Use:   "deploy [service]",
		Short: "Deploy a service (frontend or backend)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			service := args[0]
			switch service {
			case "frontend":
				deploy.Frontend()
			case "backend":
				deploy.Backend()
			default:
				fmt.Printf("Unknown service: %s\n", service)
			}
		},
	}

	rootCmd.AddCommand(deployCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"log"

	"github.com/XuryaX/card-dec-go/server"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "card-dec-go"}

	buildCmd := &cobra.Command{
		Use:   "build",
		Short: "Builds the models and migrates them to the database",
		Run: func(cmd *cobra.Command, args []string) {
			server.Build()
			log.Println("Building models and migrating to the database...")
		},
	}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Runs the card-deck API server",
		Run: func(cmd *cobra.Command, args []string) {
			server.StartServer()
			log.Println("Starting card-deck API server...")
		},
	}

	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

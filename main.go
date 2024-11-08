package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const taskFile = "tasks.json"

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Welcome to the Track-Tracker repo")

	rootCmd := &cobra.Command{
		Use:   "task-cli",
		Short: "Task Tracker CLI for managing your tasks",
	}

	rootCmd.AddCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		log.Info().Msg("Scanning....")

		time.Sleep(3 * time.Second)
	}

}

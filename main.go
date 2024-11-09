package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const taskFile = "tasks.json"

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		addTask(description)
	},
}

func addTask(description string) {
	tasks := loadTasks()

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	saveTasks(tasks)

	log.Info().Msgf("Task added successfully (ID: %d)", newTask.ID)
}

func loadTasks() []Task {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}
		}
		log.Fatal().Err(err).Msg("Could not open tasks file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err != nil && err.Error() != "EOF" {
		log.Fatal().Err(err).Msg("Could not decode tasks")
	}
	return tasks
}

func saveTasks(tasks []Task) {
	file, err := os.Create(taskFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not create tasks file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing tasks file")
		}
	}(file)

	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not format tasks to JSON")
	}

	_, err = file.Write(data)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not write tasks to file")
	}
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	rootCmd := &cobra.Command{
		Use:   "task-cli",
		Short: "Task Tracker CLI for managing your tasks",
	}

	rootCmd.AddCommand(addCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		log.Info().Msg("Hello all Here! Task-Tracker On Air!")
		log.Info().Msg("For use this app run in cmd: .\\go_build_Task_Tracker.exe add <Message> without <>")
		time.Sleep(3 * time.Second)
	}
}

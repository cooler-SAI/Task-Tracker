package main

import (
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func TestAddTask(t *testing.T) {

	log.Info().Msg("Starting TestAddTask")

	description := "Test Task"
	addTask(description)

	tasks := loadTasks()
	lastTask := tasks[len(tasks)-1]

	assert.Equal(t, description, lastTask.Description, "Task description should match")
	assert.Equal(t, "todo", lastTask.Status, "New task should have status 'todo'")
	log.Info().Msg("TestAddTask passed")
}

func TestLoadTasksEmptyFile(t *testing.T) {
	log.Info().Msg("Starting TestLoadTasksEmptyFile")

	err := os.WriteFile(taskFile, []byte(""), 0644)
	assert.NoError(t, err, "Error clearing tasks file")

	tasks := loadTasks()
	assert.Equal(t, 0, len(tasks), "Loaded tasks should be empty for a new file")
	log.Info().Msg("TestLoadTasksEmptyFile passed")
}

func TestSaveTasks(t *testing.T) {
	log.Info().Msg("Starting TestSaveTasks")

	tasks := []Task{
		{
			ID:          1,
			Description: "Saved Task",
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	saveTasks(tasks)
	loadedTasks := loadTasks()

	assert.Equal(t, len(tasks), len(loadedTasks), "Saved and loaded task count should match")
	assert.Equal(t, tasks[0].Description, loadedTasks[0].Description,
		"Loaded task description should match")
	log.Info().Msg("TestSaveTasks passed")
}

func TestAddTaskIDIncrement(t *testing.T) {
	log.Info().Msg("Starting TestAddTaskIDIncrement")

	addTask("First Task")
	tasks := loadTasks()
	initialID := tasks[len(tasks)-1].ID

	addTask("Second Task")
	tasks = loadTasks()
	newID := tasks[len(tasks)-1].ID
	assert.Equal(t, initialID+1, newID, "New task ID should match")
	log.Info().Msg("TestAddTaskIDIncrement passed")

}

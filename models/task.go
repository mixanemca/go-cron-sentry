package models

// Task represent cron task
type Task struct {
	Command string
	Args    []string
}

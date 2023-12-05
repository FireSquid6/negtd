package gtd

import (
	"github.com/firesquid6/negtd/date"
)

// handles the reading of files and parsing them into structs
// Returns a slice of tasks and a slice of errors. The errors shouldn't really be handled. They just need to be printed to a log file somewhere
func ReadInboxFile(file []string, currentDate date.Date) ([]GtdTask, []string) {
	tasks := []GtdTask{}
	errors := []string{}

	for _, line := range file {
		task, err := parseInboxTask(line, currentDate)
		if err != nil {
			errors = append(errors, "Error parsing inbox task: "+line)
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, errors
}

func ReadAgendaFile(file []string, currentDate date.Date) ([]GtdTask, []string) {
	tasks := []GtdTask{}
	errors := []string{}
	readingDate := currentDate

	for _, line := range file {
		if string(line[0]) == "*" {
			newReadingDate, err := date.ParseDate(line)
			readingDate = newReadingDate
			if err != nil {
				errors = append(errors, "Error parsing agenda date: "+line)
				continue
			}
			continue
		}

		task, err := parseAgendaTask(line, readingDate)
		if err != nil {
			errors = append(errors, "Error parsing agenda task: "+line)
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, errors
}

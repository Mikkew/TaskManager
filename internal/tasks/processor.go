package tasks

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	TaskID  int
	Message string
	Error   error
}

func (store *Store) ProcessAll() []Result {
	task := store.List()
	results := make(chan Result, len(task))

	var wg sync.WaitGroup

	for _, task := range task {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			time.Sleep(500 * time.Microsecond)
			results <- Result{
				TaskID:  task.ID,
				Message: fmt.Sprintf("Tarea %d procesada: %s", task.ID, task.Title),
			}
		}(task)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var finalResults []Result
	for result := range results {
		finalResults = append(finalResults, result)
	}

	return finalResults
}

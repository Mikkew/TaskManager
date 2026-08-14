package tasks

type Task struct {
	ID    int
	Title string
	Done  bool
}

// Exportable
func NewTask(id int, title string) Task {
	return Task{
		ID:    id,
		Title: title,
		Done:  false,
	}
}

// No exportable solo para el paquete
func formatTitle(task Task) string {
	return "[" + task.Title + "]"
}

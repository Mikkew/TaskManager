package tasks

type Store struct {
	items  []Task
	nextID int
}

func NewStore() *Store {
	return &Store{nextID: 1}
}

func (store *Store) Add(title string) Task {
	task := NewTask(store.nextID, title)
	store.items = append(store.items, task)
	store.nextID++
	return task
}

func (store *Store) List() []Task {
	return store.items
}

func (store *Store) Complete(id int) bool {
	store.items[id].Done = true
	return true
}

func (store *Store) Delete(id int) bool {
	for index, task := range store.items {
		if task.ID == id {
			store.items = append(store.items[:index], store.items[index+1:]...)
			return true
		}
	}

	return false
}

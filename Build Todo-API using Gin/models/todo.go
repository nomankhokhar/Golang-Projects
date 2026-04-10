package models

// Todo represents a single todo item
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// TodoStore holds the in-memory list and tracks the next ID
type TodoStore struct {
	Todos  []Todo
	nextID int
}

// NewTodoStore creates a TodoStore with seed data
func NewTodoStore() *TodoStore {
	return &TodoStore{
		Todos: []Todo{
			{ID: 1, Title: "Learn Golang", Status: "Incomplete"},
			{ID: 2, Title: "Build API", Status: "Incomplete"},
		},
		nextID: 3,
	}
}

// GetAll returns all todos
func (s *TodoStore) GetAll() []Todo {
	return s.Todos
}

// GetByID returns a todo by its ID, or nil if not found
func (s *TodoStore) GetByID(id int) *Todo {
	for i := range s.Todos {
		if s.Todos[i].ID == id {
			return &s.Todos[i]
		}
	}
	return nil
}

// Create adds a new todo and returns it with a generated ID
func (s *TodoStore) Create(todo Todo) Todo {
	todo.ID = s.nextID
	s.nextID++
	s.Todos = append(s.Todos, todo)
	return todo
}

// Update modifies an existing todo by ID, preserving the original ID.
// Returns the updated todo or nil if not found.
func (s *TodoStore) Update(id int, updated Todo) *Todo {
	for i := range s.Todos {
		if s.Todos[i].ID == id {
			updated.ID = id
			s.Todos[i] = updated
			return &s.Todos[i]
		}
	}
	return nil
}

// Delete removes a todo by ID. Returns true if deleted, false if not found.
func (s *TodoStore) Delete(id int) bool {
	for i := range s.Todos {
		if s.Todos[i].ID == id {
			s.Todos = append(s.Todos[:i], s.Todos[i+1:]...)
			return true
		}
	}
	return false
}

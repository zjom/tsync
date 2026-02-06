package tasks

type Service interface {
	// task lists
	CreateTaskList(*TaskList) (*TaskList, error)
	ListTaskLists() ([]*TaskList, error)
	GetTaskList(listId string) (*TaskList, error)
	UpdateTaskList(*TaskList) (*TaskList, error)
	DeleteTaskList(listId string) error

	// task items
	CreateTask(listId string, t *Task) (*Task, error)
	GetTask(taskId string) (*Task, error)
	UpdateTask(t *Task) (*Task, error)
	DeleteTask(taskId string) error
}

package task

type TaskListService interface {
	CreateTaskList(*TaskList) (*TaskList, error)
	ListTaskLists() ([]*TaskList, error)
	GetTaskList(listId string) (*TaskList, error)
	UpdateTaskList(*TaskList) (*TaskList, error)
	DeleteTaskList(listId string) error
}

type TaskService interface {
	CreateTask(listId string, t *Task) (*Task, error)
	GetTask(taskId string) (*Task, error)
	UpdateTask(t *Task) (*Task, error)
	DeleteTask(taskId string) error
}

package task

type TaskService struct {
	DataStorage
}

type Service interface {
	ServiceAdd(task Task) error
	ServiceTakeAll() []Task
	ServiceById(task string) (*Task, error)
}

type DataService interface {
	Service
}

func (t *TaskService) ServiceAdd(task Task) error {
	task.Status = StatusInProcess
	err := t.Storage.AddTask(task.Id, task.Status)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskService) ServiceTakeAll() []Task {
	return t.Storage.TakeTasks()
}
func (t *TaskService) ServiceById(task string) (*Task, error) {
	resptask, err := t.Storage.TakeTaskByID(task)
	if err != nil {
		return nil, err
	}
	return resptask, nil
}

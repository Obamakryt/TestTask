package task

type Task struct {
	Id          string `json:"id,omitempty"`
	Status      string `json:"-"`
	TitleTask   string `json:"title_task,omitempty"`
	TimeCreated string `json:"-"`
}

const (
	StatusInProcess = "In Process"
	StatusComplete  = "Complete"
)
const somewrong = "something wrong, failed response"

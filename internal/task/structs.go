package task

type Task struct {
	Id          string `json:"id,"`
	Status      string `json:"status"`
	TitleTask   string `json:"title"`
	TimeCreated string `json:"timeCreated"`
}

const (
	StatusInProcess = "In Process"
	StatusComplete  = "Complete" // for future
)
const somewrong = "something wrong, failed response"

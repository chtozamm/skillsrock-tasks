package tasks

const (
	StatusNew        = "new"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

type CreateTaskModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

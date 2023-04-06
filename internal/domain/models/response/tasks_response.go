package response

type TaskResponse struct {
	ID          uint   `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	TaskStatus  string `json:"task_status"`
	CompletedAt string `json:"completed_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

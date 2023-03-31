package request

type TaskRequest struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
}

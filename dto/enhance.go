package dto

type SubmitRequest struct {
	VideoURL         string
	TargetResolution string
}

type SubmitResult struct {
	TaskID string
}

type TaskStatus string

const (
	TaskStatusProcessing TaskStatus = "processing"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

type Task struct {
	TaskID   string
	Status   TaskStatus
	VideoURL string
}

package worker

type Payload struct {
	Crontype string `json:"crontype"`
}

type TaskType string

const (
	UpdateUserGroupMinutely TaskType = "updateUserGroup:minutely"
)

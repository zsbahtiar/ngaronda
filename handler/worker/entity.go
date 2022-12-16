package worker

type Payload struct {
	Crontype string `json:"crontype"`
}

type TaskType string

const (
	updateUserGroupMinutely TaskType = "updateUserGroup:minutely"
	updateUserGroupHourly   TaskType = "updateUserGroup:hourly"
	updateUserGroupDaily    TaskType = "updateUserGroup:daily"
	updateUserGroupWeekly   TaskType = "updateUserGroup:weekly"
	updateUserGroupMonthly  TaskType = "updateUserGroup:monthly"
)

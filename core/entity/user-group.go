package entity

import "github.com/google/uuid"

type CronType string

const (
	CronTypeMinutely CronType = "minutely"
	CronTypeHourly   CronType = "hourly"
	CronTypeDaily    CronType = "daily"
	CronTypeWeekly   CronType = "weekly"
	CronTypeMonthly  CronType = "monthly"
)

type UpsertUserGroupRequest struct {
	Serial    uuid.UUID  `json:"-"`
	GroupID   string     `json:"group_id"`
	Users     []UserRank `json:"users"`
	CronType  CronType   `json:"cron_type"`
	CreatedBy string     `json:"created_by"`
}

type UserRank struct {
	Users []string `json:"users"`
	Rank  int      `json:"rank"`
}

type UserGroup struct {
	Serial      uuid.UUID
	GroupID     string
	Users       []UserRank
	CurrentRank int
	CronType    CronType
	CreatedBy   string
}

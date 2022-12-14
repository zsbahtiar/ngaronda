package entity

import "github.com/google/uuid"

type CronType string

const (
	CronTypeDaily   CronType = "daily"
	CronTypeWeekly  CronType = "weekly"
	CronTypeMonthly CronType = "monthly"
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

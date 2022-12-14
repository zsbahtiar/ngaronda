package user_group

import (
	"encoding/json"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"time"
)

const tableName = "user_groups"

type UserGroup struct {
	ID          int
	Serial      string
	GroupID     string
	Users       json.RawMessage
	CurrentRank int
	CronType    string
	CreatedAt   *time.Time
	CreatedBy   string
}

func parseUpsertUserGroupToDto(u *entity.UpsertUserGroupRequest) *UserGroup {
	user, _ := json.Marshal(u.Users)
	return &UserGroup{
		Serial:    u.Serial.String(),
		GroupID:   u.GroupID,
		Users:     user,
		CronType:  string(u.CronType),
		CreatedBy: u.CreatedBy,
	}
}

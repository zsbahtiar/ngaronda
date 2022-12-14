package user_group

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"log"
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

func parseToUserGroupsEntity(userGroupsDto []*UserGroup) []*entity.UserGroup {
	var userGroups []*entity.UserGroup

	for _, ugDto := range userGroupsDto {
		userGroups = append(userGroups, ugDto.toUserGroupEntity())
	}

	return userGroups
}

func (u *UserGroup) toUserGroupEntity() *entity.UserGroup {
	var users []entity.UserRank
	err := json.Unmarshal(u.Users, &users)
	if err != nil {
		log.Printf("failed unmarshal users to entity: %v", err)
	}
	return &entity.UserGroup{
		Serial:      uuid.MustParse(u.Serial),
		GroupID:     u.GroupID,
		Users:       users,
		CurrentRank: u.CurrentRank,
		CronType:    entity.CronType(u.CronType),
		CreatedBy:   u.CreatedBy,
	}
}

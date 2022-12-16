package module

import (
	"context"
	"github.com/google/uuid"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"github.com/zsbahtiar/ngaronda/core/repository"
	"log"
)

type userGroup struct {
	slackRepo     repository.SlackRepository
	userGroupRepo repository.UserGroupRepository
}

type UserGroupUseCase interface {
	AssignUsersToGroup(ctx context.Context, cronType entity.CronType) error
	UpsertUserGroup(ctx context.Context, request *entity.UpsertUserGroupRequest) error
}

func NewUserGroupUseCase(
	slackRepo repository.SlackRepository,
	userGroupRepo repository.UserGroupRepository,
) UserGroupUseCase {
	return &userGroup{
		slackRepo:     slackRepo,
		userGroupRepo: userGroupRepo,
	}
}

func (u *userGroup) AssignUsersToGroup(ctx context.Context, cronType entity.CronType) error {
	userGroups, err := u.
		userGroupRepo.
		GetUserGroupsByCronType(ctx, cronType)

	if err != nil {
		return err
	}
	for _, ug := range userGroups {
		total := len(ug.Users)
		index := ug.CurrentRank
		if total > 1 {
			lastIndex := total - 1
			if index < lastIndex {
				index++
			} else if index == lastIndex {
				index = 0
			}
			err = u.slackRepo.
				AssignUsersToGroup(ctx,
					&entity.AssignUsersToGroupRequest{
						UserGroupID: ug.GroupID,
						Users:       ug.Users[index].Users,
					})
			if err != nil {
				log.Printf("failed to assign users to group slack: %v", err)
				continue
			}
			ug.CurrentRank = index
		}
	}

	return u.userGroupRepo.
		UpdateUserGroupsRank(ctx, userGroups)
}

func (u *userGroup) UpsertUserGroup(ctx context.Context, request *entity.UpsertUserGroupRequest) error {
	request.Serial = uuid.New()
	return u.userGroupRepo.UpsertUserGroup(ctx, request)
}

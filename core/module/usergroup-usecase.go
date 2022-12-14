package module

import (
	"context"
	"github.com/google/uuid"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"github.com/zsbahtiar/ngaronda/core/repository"
)

type userGroup struct {
	slackRepo     repository.SlackRepository
	userGroupRepo repository.UserGroupRepository
}

type UserGroupUseCase interface {
	AssignUsersToGroup(ctx context.Context, request *entity.AssignUsersToGroupRequest) error
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

func (u *userGroup) AssignUsersToGroup(ctx context.Context, request *entity.AssignUsersToGroupRequest) error {
	return u.slackRepo.AssignUsersToGroup(ctx, request)
}

func (u *userGroup) UpsertUserGroup(ctx context.Context, request *entity.UpsertUserGroupRequest) error {
	request.Serial = uuid.New()
	return u.userGroupRepo.UpsertUserGroup(ctx, request)
}

package repository

import (
	"context"
	"github.com/zsbahtiar/ngaronda/core/entity"
)

type UserGroupRepository interface {
	UpsertUserGroup(ctx context.Context, request *entity.UpsertUserGroupRequest) error
	GetUserGroupsByCronType(ctx context.Context, cronType entity.CronType) ([]*entity.UserGroup, error)
	UpdateUserGroupsRank(ctx context.Context, userGroups []*entity.UserGroup) error
}

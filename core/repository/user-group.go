package repository

import (
	"context"
	"github.com/zsbahtiar/ngaronda/core/entity"
)

type UserGroupRepository interface {
	UpsertUserGroup(ctx context.Context, request *entity.UpsertUserGroupRequest) error
}

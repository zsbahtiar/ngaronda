package repository

import (
	"context"
	"github.com/zsbahtiar/ngaronda/core/entity"
)

type SlackRepository interface {
	AssignUsersToGroup(ctx context.Context, request *entity.AssignUsersToGroupRequest) error
}

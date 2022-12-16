package task

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"github.com/zsbahtiar/ngaronda/core/module"
)

type taskHandle struct {
	userGroupUsecase module.UserGroupUseCase
}

type TaskHandle interface {
	HandleTaskUsersGroup(ctx context.Context, t *asynq.Task) error
}

func NewTaskHandler(userGroupUsecase module.UserGroupUseCase) TaskHandle {
	return &taskHandle{userGroupUsecase: userGroupUsecase}
}

func (th *taskHandle) HandleTaskUsersGroup(ctx context.Context, t *asynq.Task) error {
	p := entity.PayloadTask{}
	err := json.Unmarshal(t.Payload(), &p)
	if err != nil {
		return err
	}
	return th.userGroupUsecase.AssignUsersToGroup(ctx, p.Crontype)
}

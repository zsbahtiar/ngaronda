package worker

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/zsbahtiar/ngaronda/core/entity"
)

func (w *worker) handleUserGroupMinutely(ctx context.Context, t *asynq.Task) error {
	p := Payload{}
	err := json.Unmarshal(t.Payload(), &p)
	if err != nil {
		return err
	}
	return w.userGroupUsecase.AssignUsersToGroup(ctx, entity.CronType(p.Crontype))
}

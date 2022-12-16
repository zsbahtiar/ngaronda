package worker

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"log"
)

func (w *worker) updateUsersGroupMinutely() {
	p := Payload{Crontype: string(entity.CronTypeMinutely)}
	b, _ := json.Marshal(&p)

	task := asynq.NewTask(string(UpdateUserGroupMinutely), b)
	entryID, err := w.scheduler.Register(w.cronspec.UpdateUsersGroup, task)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", UpdateUserGroupMinutely, err)
	}

	log.Printf("success register %s to scheduler with entry id: %s", UpdateUserGroupMinutely, entryID)
}

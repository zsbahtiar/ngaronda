package worker

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"log"
)

func (w *worker) updateUsersGroupMinutely() {
	p := Payload{Crontype: string(entity.CronTypeMinutely)}
	b, err := json.Marshal(&p)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupMinutely, err)
	}
	task := asynq.NewTask(string(updateUserGroupMinutely), b)
	entryID, err := w.scheduler.Register(w.cronspec.UpdateUsersGroupMinutely, task)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupMinutely, err)
	}

	log.Printf("success register %s to scheduler with entry id: %s", updateUserGroupMinutely, entryID)
}

func (w *worker) updateUsersGroupHourly() {
	p := Payload{Crontype: string(entity.CronTypeHourly)}
	b, err := json.Marshal(&p)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupHourly, err)
	}
	task := asynq.NewTask(string(updateUserGroupHourly), b)
	entryID, err := w.scheduler.Register(w.cronspec.UpdateUsersGroupHourly, task)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupHourly, err)
	}

	log.Printf("success register %s to scheduler with entry id: %s", updateUserGroupHourly, entryID)
}

func (w *worker) updateUsersGroupDaily() {
	p := Payload{Crontype: string(entity.CronTypeDaily)}
	b, err := json.Marshal(&p)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupDaily, err)
	}
	task := asynq.NewTask(string(updateUserGroupDaily), b)
	entryID, err := w.scheduler.Register(w.cronspec.UpdateUsersGroupDaily, task)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupDaily, err)
	}

	log.Printf("success register %s to scheduler with entry id: %s", updateUserGroupDaily, entryID)
}

func (w *worker) updateUsersGroupWeekly() {
	p := Payload{Crontype: string(entity.CronTypeWeekly)}
	b, err := json.Marshal(&p)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupWeekly, err)
	}
	task := asynq.NewTask(string(updateUserGroupWeekly), b)
	entryID, err := w.scheduler.Register(w.cronspec.UpdateUsersGroupWeekly, task)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupWeekly, err)
	}

	log.Printf("success register %s to scheduler with entry id: %s", updateUserGroupWeekly, entryID)

}

func (w *worker) updateUsersGroupMonthly() {
	p := Payload{Crontype: string(entity.CronTypeMonthly)}
	b, err := json.Marshal(&p)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupMonthly, err)
	}
	task := asynq.NewTask(string(updateUserGroupMonthly), b)
	entryID, err := w.scheduler.Register(w.cronspec.UpdateUsersGroupMonthly, task)
	if err != nil {
		log.Fatalf("failed register %s to scheduler, err: %v ", updateUserGroupMonthly, err)
	}

	log.Printf("success register %s to scheduler with entry id: %s", updateUserGroupMonthly, entryID)

}

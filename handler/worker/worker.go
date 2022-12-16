package worker

import (
	"github.com/hibiken/asynq"
	"github.com/zsbahtiar/ngaronda/config"
	"github.com/zsbahtiar/ngaronda/handler/worker/task"
)

type worker struct {
	redisConn   asynq.RedisClientOpt
	scheduler   *asynq.Scheduler
	cronspec    config.ScheduleCron
	taskHandler task.TaskHandle
}

type Worker interface {
	Start() error
}

func NewWorker(
	redisConn asynq.RedisClientOpt,
	cronspec config.ScheduleCron,
	taskHandler task.TaskHandle,
) Worker {
	scheduler := asynq.NewScheduler(redisConn, &asynq.SchedulerOpts{})
	return &worker{
		redisConn,
		scheduler,
		cronspec,
		taskHandler,
	}
}

// Start running asynq server and periodic task or scheduler
func (w *worker) Start() error {
	w.applyRegistry()

	go w.startServer(w.redisConn)

	return w.scheduler.Run()

}

// applyRegistry method apply all task on registry to scheduler
func (w *worker) applyRegistry() {
	w.updateUsersGroupMinutely()
	w.updateUsersGroupDaily()
	w.updateUsersGroupWeekly()
	w.updateUsersGroupMonthly()

}

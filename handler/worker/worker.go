package worker

import (
	"github.com/hibiken/asynq"
	"github.com/zsbahtiar/ngaronda/config"
	"github.com/zsbahtiar/ngaronda/core/module"
)

type worker struct {
	redisConn        asynq.RedisClientOpt
	scheduler        *asynq.Scheduler
	cronspec         config.ScheduleCron
	userGroupUsecase module.UserGroupUseCase
}

type Worker interface {
	Start() error
}

func NewWorker(
	redisConn asynq.RedisClientOpt,
	cronspec config.ScheduleCron,
	userGroupUsecase module.UserGroupUseCase,
) Worker {
	scheduler := asynq.NewScheduler(redisConn, &asynq.SchedulerOpts{})
	return &worker{
		redisConn,
		scheduler,
		cronspec,
		userGroupUsecase,
	}
}

func (w *worker) Start() error {
	go w.startServer(w.redisConn)
	w.register()
	return w.scheduler.Run()

}

func (w *worker) register() {
	w.updateUsersGroupMinutely()
}

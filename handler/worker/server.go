package worker

import (
	"github.com/hibiken/asynq"
	"log"
)

func (w *worker) startServer(redis asynq.RedisClientOpt) {
	srv := asynq.NewServer(
		redis,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(string(UpdateUserGroupMinutely), w.handleUserGroupMinutely)
	// ...register other handlers...

	log.Println("start worker server")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

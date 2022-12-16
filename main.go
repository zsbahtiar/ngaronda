package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/zsbahtiar/ngaronda/config"
	"github.com/zsbahtiar/ngaronda/core/module"
	"github.com/zsbahtiar/ngaronda/handler/api"
	"github.com/zsbahtiar/ngaronda/handler/worker"
	"github.com/zsbahtiar/ngaronda/handler/worker/task"
	"github.com/zsbahtiar/ngaronda/pkg/database"
	"github.com/zsbahtiar/ngaronda/repository/slack"
	userGroup "github.com/zsbahtiar/ngaronda/repository/user-group"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.Get()

	redisAsyncConn := asynq.RedisClientOpt{
		Addr:        fmt.Sprintf("%s:%d", cfg.RedisHostName, cfg.RedisPort),
		Username:    cfg.RedisUserName,
		Password:    cfg.RedisPassword,
		DialTimeout: time.Second * 10,
		ReadTimeout: time.Second * 10,
		TLSConfig:   nil,
	}

	db := database.InitMysql(cfg.Database)
	userGroupRepo := userGroup.NewRepository(db)
	slackRepo := slack.NewRepository(cfg.SlackBaseURL, cfg.SlackAPIKey, cfg.SlackBotAPIKey)
	userGroupUsecase := module.NewUserGroupUseCase(slackRepo, userGroupRepo)
	taskHandler := task.NewTaskHandler(userGroupUsecase)

	userGroupApi := api.NewUerGroupApi(userGroupUsecase)

	srv := gin.Default()
	srv.POST("/user-group", userGroupApi.UpsertUsersToGroup)
	srv.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"isSuccess": true,
			"message":   "sehat walafiat",
		})
	})

	go func() {
		err := srv.Run(fmt.Sprintf(":%s", cfg.HTTPPort))
		if err != nil {
			log.Fatal(err)
		}
	}()

	w := worker.NewWorker(redisAsyncConn, cfg.ScheduleCron, taskHandler)

	if err := w.Start(); err != nil {
		log.Fatal(err)
	}

}

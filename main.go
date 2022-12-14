package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zsbahtiar/ngaronda/config"
	"github.com/zsbahtiar/ngaronda/core/module"
	"github.com/zsbahtiar/ngaronda/handler/api"
	"github.com/zsbahtiar/ngaronda/pkg/database"
	"github.com/zsbahtiar/ngaronda/repository/slack"
	userGroup "github.com/zsbahtiar/ngaronda/repository/user-group"
	"log"
)

func main() {
	cfg := config.Get()
	db := database.InitMysql(cfg.Database)
	userGroupRepo := userGroup.NewRepository(db)
	slackRepo := slack.NewRepository(cfg.SlackBaseURL, cfg.SlackAPIKey)
	userGroupUsecase := module.NewUserGroupUseCase(slackRepo, userGroupRepo)

	userGroupApi := api.NewUerGroupApi(userGroupUsecase)

	srv := gin.Default()
	srv.POST("/user-group", userGroupApi.UpsertUsersToGroup)

	if err := srv.Run(fmt.Sprintf(":%s", cfg.HTTPPort)); err != nil {
		log.Fatal(err)
	}

}

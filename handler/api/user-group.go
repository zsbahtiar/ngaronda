package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"github.com/zsbahtiar/ngaronda/core/module"
	"net/http"
)

type userGroupApi struct {
	userGroupUseCase module.UserGroupUseCase
}

type UserGroupApi interface {
	UpsertUsersToGroup(c *gin.Context)
}

func NewUerGroupApi(userGroupUseCase module.UserGroupUseCase) UserGroupApi {
	return &userGroupApi{userGroupUseCase: userGroupUseCase}
}

func (u *userGroupApi) UpsertUsersToGroup(c *gin.Context) {
	var request entity.UpsertUserGroupRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"isSuccess": false,
			"message":   err.Error(),
		})
		return
	}
	err = u.userGroupUseCase.UpsertUserGroup(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"isSuccess": false,
			"message":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"isSuccess": true,
		"message":   "success",
	})
}

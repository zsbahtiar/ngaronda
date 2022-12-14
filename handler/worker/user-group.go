package worker

import (
	"github.com/gin-gonic/gin"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"github.com/zsbahtiar/ngaronda/core/module"
	"net/http"
)

type userGroupWorker struct {
	userGroupUseCase module.UserGroupUseCase
}

type UserGroupWorker interface {
	AssignUsersToGroup(c *gin.Context)
}

func NewUerGroupWorker(userGroupUseCase module.UserGroupUseCase) UserGroupWorker {
	return &userGroupWorker{userGroupUseCase: userGroupUseCase}
}

func (u *userGroupWorker) AssignUsersToGroup(c *gin.Context) {
	var request entity.AssignUsersToGroupRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"isSuccess": false,
			"message":   err.Error(),
		})
		return
	}
	err = u.userGroupUseCase.AssignUsersToGroup(c, &request)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"isSuccess": false,
				"message":   err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isSuccess": true,
		"message":   "success",
	})
}

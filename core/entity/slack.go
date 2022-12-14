package entity

type AssignUsersToGroupRequest struct {
	UserGroupID string   `json:"user_group_id"`
	Users       []string `json:"users"`
}

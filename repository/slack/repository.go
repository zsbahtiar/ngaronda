package slack

import (
	"context"
	"fmt"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"github.com/zsbahtiar/ngaronda/core/repository"
	"github.com/zsbahtiar/ngaronda/pkg/client"
	"net/http"
	"strings"
)

type slack struct {
	baseURL   string
	secretKey string
}

func NewRepository(baseURL string, secretKey string) repository.SlackRepository {
	return &slack{baseURL: baseURL, secretKey: secretKey}
}

func (s *slack) AssignUsersToGroup(ctx context.Context, request *entity.AssignUsersToGroupRequest) error {
	reqMap := map[string]map[string]string{
		"header": {
			"Authorization": fmt.Sprintf("Bearer %s", s.secretKey),
			"Content-Type":  "application/x-www-form-urlencoded; application/json",
		},
		"query": {
			"usergroup": request.UserGroupID,
			"users":     strings.Join(request.Users, ","),
		},
	}

	resp, err := client.Do(ctx, fmt.Sprintf("%susergroups.users.update", s.baseURL), http.MethodPost, reqMap)
	if err != nil {
		return err
	}

	if resp["ok"].(bool) != true {
		return fmt.Errorf("failed to asssign users to group")
	}

	return nil
}

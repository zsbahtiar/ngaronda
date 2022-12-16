package slack

import (
	"context"
	"fmt"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"github.com/zsbahtiar/ngaronda/core/repository"
	"github.com/zsbahtiar/ngaronda/pkg/client"
	"log"
	"net/http"
	"strings"
)

type slack struct {
	baseURL      string
	secretKey    string
	secretBotKey string
}

func NewRepository(baseURL string, secretKey string, secretBotKey string) repository.SlackRepository {
	return &slack{baseURL: baseURL, secretKey: secretKey, secretBotKey: secretBotKey}
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

	// ignored error
	s.tellHasAssigned(ctx, request)

	return nil
}

// tellHasAssigned bot chat to user tell has been assign on user group
func (s *slack) tellHasAssigned(ctx context.Context, req *entity.AssignUsersToGroupRequest) {
	reqMap := map[string]map[string]string{
		"header": {
			"Authorization": fmt.Sprintf("Bearer %s", s.secretBotKey),
			"Content-Type":  "application/x-www-form-urlencoded; application/json",
		},
		"query": {
			"channel": "",
			"text":    fmt.Sprintf("Halo, kamu telah ditambahkan ke <!subteam^%s> sebagai kang ronda. Semangat yaaa!!!", req.UserGroupID),
		},
	}

	for _, user := range req.Users {
		reqMap["query"]["channel"] = user
		resp, err := client.Do(ctx, fmt.Sprintf("%schat.postMessage", s.baseURL), http.MethodPost, reqMap)
		if err != nil {
			log.Printf("failed to tell user: %s has assigned on: %s with err: %v", user, req.UserGroupID, err)
		}

		if resp["ok"].(bool) != true {
			log.Printf("failed to tell user: %s has assigned on: %s", user, req.UserGroupID)
		}
	}

}

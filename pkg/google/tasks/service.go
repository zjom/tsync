package tasks

import (
	"context"
	"net/http"

	"github.com/zjom/tsync/pkg/google/oauth"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

func init() {
	oauth.RegisterScope(tasks.TasksScope)
}

func NewService(client *http.Client, ctx context.Context) (*tasks.Service, error) {
	srv, err := tasks.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return srv, nil
}

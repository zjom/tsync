package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zjom/gts/pkg/google/oauth"
	"github.com/zjom/gts/pkg/google/tasks"
)

func main() {
	ctx := context.Background()
	client, err := oauth.GetClient(
		oauth.WithContext(ctx))
	if err != nil {
		panic(err)
	}
	srv, err := tasks.NewService(client, ctx)
	if err != nil {
		panic(err)
	}

	// List Task Lists
	t, err := srv.Tasklists.List().MaxResults(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists: %v", err)
	}

	fmt.Println("Task Lists:")
	if len(t.Items) == 0 {
		fmt.Print("No task lists found.")
	} else {
		for _, i := range t.Items {
			fmt.Printf("%s (%s)\n", i.Title, i.Id)
		}
	}
}

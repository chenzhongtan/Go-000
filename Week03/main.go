package main

import (
	"Week03/HttpService"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	group ,ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return HttpService.Service(ctx, ":8080")
	})

	group.Go(func() error {
		return HttpService.Service(ctx, "8081")
	})

	group.Go(func() error {
		return HttpService.Service(ctx, "8082")
	})

	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

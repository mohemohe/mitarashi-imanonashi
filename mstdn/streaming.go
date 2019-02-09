package mstdn

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/mohemohe/mitarashi-imanonashi/config"
)

type (
	StreamEvent struct {
		Raw mastodon.Event
	}
)

func OnStream() (chan StreamEvent, error) {
	client := mastodon.NewClient(&mastodon.Config{
		Server:      config.GetEnv().Mastodon.Server,
		AccessToken: config.GetEnv().Mastodon.AccessToken,
	})

	ch, err := client.StreamingUser(context.Background())
	if err != nil {
		return nil, err
	}

	c := make(chan StreamEvent)
	go func() {
		for {
			r := <- ch
			c <- StreamEvent{
				Raw: r,
			}
		}
	}()
	return c, nil
}

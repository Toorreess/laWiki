package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

type Client struct {
	Project string
	Storage *firestore.Client
	Ctx     context.Context
}

func (c *Client) Init(ctx context.Context) error {
	fsClient, err := firestore.NewClient(ctx, c.Project)
	if err != nil {
		return err
	}
	c.Storage = fsClient
	c.Ctx = ctx
	return nil
}

func (c *Client) Close() error {
	if c.Storage == nil {
		return fmt.Errorf("no client found")
	}
	return c.Storage.Close()
}

package firestore

import (
	"context"
	"fmt"

	"encoding/json"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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

func (c Client) Close() error {
	if c.Storage == nil {
		return fmt.Errorf("no client found")
	}
	return c.Storage.Close()
}

func (c Client) Get(index, id string, entity interface{}) (map[string]interface{}, error) {
	if c.Storage == nil {
		return nil, fmt.Errorf("no client found.")
	}

	collection := c.Storage.Collection(index)
	doc := collection.Doc(id)
	docsnap, err := doc.Get(c.Ctx)

	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	if err := docsnap.DataTo(&result); err != nil {
		return nil, err
	}
	result["id"] = id

	if _, ok := result["creation_date"]; !ok {
		result["creation_date"] = docsnap.CreateTime
	}
	if _, ok := result["modification_date"]; !ok {
		result["modification_date"] = docsnap.UpdateTime
	}
	if result["deleted"].(bool) {
		return nil, fmt.Errorf("not found")
	}
	return result, nil
}

func (c Client) Create(index string, entity interface{}) (map[string]interface{}, error) {
	if c.Storage == nil {
		return nil, fmt.Errorf("no client found.")
	}

	collection := c.Storage.Collection(index)
	doc, wr, err := collection.Add(c.Ctx, entity)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})

	inrec, err := json.Marshal(entity)
	json.Unmarshal(inrec, &result)

	result["id"] = doc.ID
	if _, ok := result["creation_date"]; !ok {
		result["creation_date"] = wr.UpdateTime
	}
	if _, ok := result["modification_date"]; !ok {
		result["modification_date"] = wr.UpdateTime
	}

	return result, nil
}

func (c Client) Update(index, id string, entity interface{}, updates map[string]interface{}) (map[string]interface{}, error) {
	if c.Storage == nil {
		return nil, fmt.Errorf("no client found.")
	}

	collection := c.Storage.Collection(index)
	doc := collection.Doc(id)
	var fsUpdates []firestore.Update
	for k, v := range updates {
		fsUpdates = append(fsUpdates, firestore.Update{Path: k, Value: v})
	}

	_, err := doc.Update(c.Ctx, fsUpdates)
	if err != nil {
		return nil, err
	}
	return c.Get(index, id, entity)
}

func (c Client) Delete(index, id string) error {
	if c.Storage == nil {
		return fmt.Errorf("no client found.")
	}

	collection := c.Storage.Collection(index)
	doc := collection.Doc(id)
	if _, err := doc.Update(c.Ctx, []firestore.Update{{Path: "deleted", Value: true}}); err != nil {
		return err
	}
	return nil
}

func (c Client) List(index string, query map[string]string, limit, offset int, orderBy, order string, entity interface{}) ([]map[string]interface{}, error) {
	if c.Storage == nil {
		return nil, fmt.Errorf("no client found.")
	}

	collection := c.Storage.Collection(index)
	q := collection.Query

	for k, v := range query {
		if v != "" {
			q = q.Where(k, ">=", v)
		}
	}

	q = q.Where("deleted", "==", false)

	if orderBy != "" && order != "" {
		var fbDirection firestore.Direction
		if order == "ASC" {
			fbDirection = firestore.Asc
		} else {
			fbDirection = firestore.Desc
		}
		q = q.OrderBy(orderBy, fbDirection)
	}

	q = q.Limit(limit)
	q = q.Offset(offset)

	iter := q.Documents(c.Ctx)
	defer iter.Stop()

	var results []map[string]interface{}
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		if err := docSnap.DataTo(&result); err != nil {
			return nil, err
		}

		result["id"] = docSnap.Ref.ID
		result["modification_date"] = docSnap.UpdateTime
		result["creation_date"] = docSnap.CreateTime

		results = append(results, result)
	}

	return results, nil
}

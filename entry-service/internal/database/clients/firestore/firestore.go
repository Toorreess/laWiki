package firestore

import (
	"context"
	"fmt"
	"io"
	"time"

	"encoding/json"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/Toorreess/laWiki/entry-service/internal/model"
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

	vcsCollection := doc.Collection("VCS")
	iter := vcsCollection.Where("deleted", "==", false).Documents(c.Ctx)
	defer iter.Stop()

	var results []string
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		var version model.Version
		if err := docSnap.DataTo(&version); err != nil {
			return nil, err
		}
		if version.Latest == true {
			result["content_url"] = version.ContentURL
			result["latest_version"] = docSnap.Ref.ID
		}

		version.ID = docSnap.Ref.ID
		results = append(results, version.ID)
	}

	result["version_list"] = results
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
	content := entity.(*model.Entry).Content

	url, err := UploadFileFromMemory(fmt.Sprintf("%s-%d.txt", doc.ID, time.Now().Unix()), []byte(content))
	if err != nil {
		return nil, err
	}

	entryVersion := model.Version{
		Author:     result["author"].(string),
		Latest:     true,
		Deleted:    false,
		ContentURL: url,
	}

	doc, _, err = doc.Collection("VCS").Add(c.Ctx, entryVersion)
	if err != nil {
		return nil, err
	}
	result["version_list"] = []string{doc.ID}
	result["latest_version"] = doc.ID
	result["content_url"] = url

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

	vcs := doc.Collection("VCS").Where("latest", "==", true)
	iter := vcs.Documents(c.Ctx)
	defer iter.Stop()

	var result model.Version
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		if err := docSnap.DataTo(&result); err != nil {
			return nil, err
		}

		doc.Collection("VCS").Doc(result.ID).Update(c.Ctx, []firestore.Update{
			{
				Path:  "latest",
				Value: false,
			},
		})
	}
	content := entity.(model.Entry).Content

	url, err := UploadFileFromMemory(fmt.Sprintf("%s-%d.txt", doc.ID, time.Now().Unix()), []byte(content))
	if err != nil {
		return nil, err
	}

	entryVersion := model.Version{
		Author:     result.Author,
		Latest:     true,
		Deleted:    false,
		ContentURL: url,
	}

	_, _, err = doc.Collection("VCS").Add(c.Ctx, entryVersion)
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

func (c Client) SetLatest(index, entry_id, version_id string) error {
	entry, err := c.Get(index, entry_id, model.Entry{})
	if err != nil {
		return err
	}

	latestVersion := entry["latest_version"].(string)
	doc := c.Storage.Collection(index).Doc(entry_id).Collection("VCS").Doc(latestVersion)
	if _, err := doc.Update(c.Ctx, []firestore.Update{{Path: "latest", Value: false}}); err != nil {
		return err
	}

	doc = c.Storage.Collection(index).Doc(entry_id).Collection("VCS").Doc(version_id)
	if _, err := doc.Update(c.Ctx, []firestore.Update{{Path: "latest", Value: true}}); err != nil {
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

		vcsIter := c.Storage.Collection(index).Doc(docSnap.Ref.ID).Collection("VCS").Documents(c.Ctx)
		var versions []string

		for {
			versionSnap, err := vcsIter.Next()
			if err == iterator.Done {
				break
			}

			if err != nil {
				return nil, err
			}
			var version model.Version
			if err := versionSnap.DataTo(&version); err != nil {
				return nil, err
			}
			if version.Latest == true {
				result["latest_version"] = version.ID
				result["content_url"] = version.ContentURL
			}
			versions = append(versions, version.ID)
		}

		result["version_list"] = versions
		result["id"] = docSnap.Ref.ID
		result["modification_date"] = docSnap.UpdateTime
		result["creation_date"] = docSnap.CreateTime

		results = append(results, result)
	}

	return results, nil
}

func DownloadFileIntoMemory(object string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := client.Bucket("lawiki-89989.appspot.com").Object(object).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %v", object, err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}

	return data, nil
}

func UploadFileFromMemory(object string, data []byte) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := client.Bucket("lawiki-89989.appspot.com").Object(object).NewWriter(ctx)

	if _, err = wc.Write(data); err != nil {
		return "", fmt.Errorf("Writer.Write: %v", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	url := fmt.Sprintf("https://storage.googleapis.com/lawiki-89989.appspot.com/%s", object)
	return url, nil
}

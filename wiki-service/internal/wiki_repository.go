package internal

import (
	"github.com/Toorreess/laWiki/wiki-service/internal/database"
	"github.com/Toorreess/laWiki/wiki-service/internal/model"
)

type IWikiRepository interface {
	Create(wm *model.Wiki) (map[string]interface{}, error)
	Get(id string) (map[string]interface{}, error)
	Update(id string, updates map[string]interface{}) (map[string]interface{}, error)
	Delete(id string) error

	List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error)
}

type wikiRepository struct {
	db *database.Connection
}

func NewWikiRepository(db *database.Connection) IWikiRepository {
	return &wikiRepository{db: db}
}

const WIKI_INDEX_NAME = "Wiki"

func (wr *wikiRepository) Create(wm *model.Wiki) (map[string]interface{}, error) {
	wmr, err := wr.db.Create(WIKI_INDEX_NAME, wm)
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (wr *wikiRepository) Get(id string) (map[string]interface{}, error) {
	wmr, err := wr.db.Get(WIKI_INDEX_NAME, id, model.Wiki{})
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (wr *wikiRepository) Update(id string, updates map[string]interface{}) (map[string]interface{}, error) {
	wmr, err := wr.db.Update(WIKI_INDEX_NAME, id, model.Wiki{}, updates)
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (wr *wikiRepository) Delete(id string) error {
	err := wr.db.Delete(WIKI_INDEX_NAME, id)
	if err != nil {
		return err
	}
	return nil
}

func (wr *wikiRepository) List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error) {
	wmr, err := wr.db.List(WIKI_INDEX_NAME, query, limit, offset, orderBy, order, model.Wiki{})
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

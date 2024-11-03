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

	List(query string, limit, offset int) ([]map[string]interface{}, error)
}

type wikiRepository struct {
	db *database.Connection
}

func NewWikiRepository(db *database.Connection) IWikiRepository {
	return &wikiRepository{db: db}
}

const WIKI_INDEX_GAME = "Wiki"

func (wr *wikiRepository) Create(wm *model.Wiki) (map[string]interface{}, error) {
	wmr, err := wr.db.Create(WIKI_INDEX_GAME, wm)
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (wr *wikiRepository) Get(id string) (map[string]interface{}, error) {
	wmr, err := wr.db.Get(WIKI_INDEX_GAME, id, model.Wiki{})
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (wr *wikiRepository) Update(id string, updates map[string]interface{}) (map[string]interface{}, error) {
	wmr, err := wr.db.Update(WIKI_INDEX_GAME, id, model.Wiki{}, updates)
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (wr *wikiRepository) Delete(id string) error {
	err := wr.db.Delete(WIKI_INDEX_GAME, id)
	if err != nil {
		return err
	}
	return nil
}

func (wr *wikiRepository) List(query string, limit, offset int) ([]map[string]interface{}, error) {
	wmr, err := wr.db.List(WIKI_INDEX_GAME, query, limit, offset, model.Wiki{})
	if err != nil {
		return nil, err
	}
	return wmr, nil
}
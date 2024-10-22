package internal

import (
	"github.com/Toorreess/laWiki/wiki-service/internal/model"
	"github.com/Toorreess/laWiki/wiki-service/pkg/database"
)

type IWikiRepository interface {
	Create(wm *model.Wiki) (map[string]interface{}, error)
	Read(id string) (map[string]interface{}, error)
	Update(id string, w map[string]interface{}) (map[string]interface{}, error)
	Delete(id string) error

	List(query string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error)
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

func (wr *wikiRepository) Read(id string) (map[string]interface{}, error) {
	wmr, err := wr.db.Read(WIKI_INDEX_GAME, id)
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (wr *wikiRepository) Update(id string, wm map[string]interface{}) (map[string]interface{}, error) {
	wmr, err := wr.db.Update(WIKI_INDEX_GAME, id, wm)
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

func (wr *wikiRepository) List(query string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error) {
	wmr, err := wr.db.List(WIKI_INDEX_GAME, query, limit, offset, orderBy, order, model.Wiki{})
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

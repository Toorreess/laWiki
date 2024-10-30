package internal

import (
	"github.com/Toorreess/laWiki/entry-service/internal/database"
	"github.com/Toorreess/laWiki/entry-service/internal/model"
)

type IEntryRepository interface {
	Create(wm *model.Entry) (map[string]interface{}, error)
	Get(id string) (map[string]interface{}, error)
	Update(id string, updates map[string]interface{}) (map[string]interface{}, error)
	Delete(id string) error

	List(query string, limit, offset int) ([]map[string]interface{}, error)
}

type entryRepository struct {
	db *database.Connection
}

func NewEntryRepository(db *database.Connection) IEntryRepository {
	return &entryRepository{db: db}
}

const ENTRY_INDEX_GAME = "Entry"

func (er *entryRepository) Create(wm *model.Entry) (map[string]interface{}, error) {
	wmr, err := er.db.Create(ENTRY_INDEX_GAME, wm)
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (er *entryRepository) Get(id string) (map[string]interface{}, error) {
	wmr, err := er.db.Client.(database.DBClient).Get(ENTRY_INDEX_GAME, id, model.Entry{})
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (er *entryRepository) Update(id string, updates map[string]interface{}) (map[string]interface{}, error) {
	wmr, err := er.db.Update(ENTRY_INDEX_GAME, id, updates)
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

func (er *entryRepository) Delete(id string) error {
	err := er.db.Delete(ENTRY_INDEX_GAME, id)
	if err != nil {
		return err
	}
	return nil
}

func (er *entryRepository) List(query string, limit, offset int) ([]map[string]interface{}, error) {
	wmr, err := er.db.List(ENTRY_INDEX_GAME, query, limit, offset, model.Entry{})
	if err != nil {
		return nil, err
	}
	return wmr, nil
}

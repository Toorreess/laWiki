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

	List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error)
}

type entryRepository struct {
	db *database.Connection
}

func NewEntryRepository(db *database.Connection) IEntryRepository {
	return &entryRepository{db: db}
}

const ENTRY_INDEX_NAME = "Entry_v2"
const VERSION_INDEX_NAME = "VCS"

func (er *entryRepository) Create(wm *model.Entry) (map[string]interface{}, error) {
	emr, err := er.db.Create(ENTRY_INDEX_NAME, wm)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *entryRepository) Get(id string) (map[string]interface{}, error) {
	emr, err := er.db.Client.(database.DBClient).Get(ENTRY_INDEX_NAME, id)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *entryRepository) Update(id string, updates map[string]interface{}) (map[string]interface{}, error) {
	emr, err := er.db.Update(ENTRY_INDEX_NAME, id, updates)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *entryRepository) Delete(id string) error {
	err := er.db.Delete(ENTRY_INDEX_NAME, id)
	if err != nil {
		return err
	}
	return nil
}

func (er *entryRepository) List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error) {
	emr, err := er.db.List(ENTRY_INDEX_NAME, query, limit, offset, orderBy, order)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

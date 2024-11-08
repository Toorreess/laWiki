package internal

import (
	"github.com/Toorreess/laWiki/comment-service/internal/database"
	"github.com/Toorreess/laWiki/comment-service/internal/model"
)

type ICommentRepository interface {
	Create(wm *model.Comment) (map[string]interface{}, error)
	Get(id string) (map[string]interface{}, error)
	Update(id string, update map[string]interface{}) (map[string]interface{}, error)
	Delete(id string) error

	List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error)
}

type commentRepository struct {
	db *database.Connection
}

func NewCommentRepository(db *database.Connection) ICommentRepository {
	return &commentRepository{db: db}
}

const ENTRY_INDEX_NAME = "Comment"

func (er *commentRepository) Create(wm *model.Comment) (map[string]interface{}, error) {
	emr, err := er.db.Create(ENTRY_INDEX_NAME, wm)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *commentRepository) Get(id string) (map[string]interface{}, error) {
	emr, err := er.db.Client.(database.DBClient).Get(ENTRY_INDEX_NAME, id, model.Comment{})
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *commentRepository) Update(id string, updates map[string]interface{}) (map[string]interface{}, error) {
	emr, err := er.db.Update(ENTRY_INDEX_NAME, id, model.Comment{}, updates)
	if err != nil {
		return nil, err
	}
	return emr, nil
}

func (er *commentRepository) Delete(id string) error {
	err := er.db.Delete(ENTRY_INDEX_NAME, id)
	if err != nil {
		return err
	}
	return nil
}

func (er *commentRepository) List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error) {
	emr, err := er.db.List(ENTRY_INDEX_NAME, query, limit, offset, orderBy, order, model.Comment{})
	if err != nil {
		return nil, err
	}
	return emr, nil
}

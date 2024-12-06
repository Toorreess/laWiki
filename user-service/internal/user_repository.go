package internal

import (
	"github.com/Toorreess/laWiki/user-service/internal/database"
	"github.com/Toorreess/laWiki/user-service/internal/model"
)

type IUserRepository interface {
	Create(um *model.User) (map[string]interface{}, error)
	Get(id string) (map[string]interface{}, error)
	Update(id string, updates map[string]interface{}) (map[string]interface{}, error)
	Delete(id string) error
	List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error)
}

type userRepository struct {
	db *database.Connection
}

func NewUserRepository(db *database.Connection) IUserRepository {
	return &userRepository{db: db}
}

const USER_INDEX_NAME = "User"

func (ur *userRepository) Create(um *model.User) (map[string]interface{}, error) {
	umr, err := ur.db.Create(USER_INDEX_NAME, um)
	if err != nil {
		return nil, err
	}
	return umr, nil
}

func (ur *userRepository) Get(id string) (map[string]interface{}, error) {
	umr, err := ur.db.Get(USER_INDEX_NAME, id)
	if err != nil {
		return nil, err
	}
	return umr, nil
}

func (ur *userRepository) Update(id string, updates map[string]interface{}) (map[string]interface{}, error) {
	umr, err := ur.db.Update(USER_INDEX_NAME, id, updates)
	if err != nil {
		return nil, err
	}
	return umr, nil
}

func (ur *userRepository) Delete(id string) error {
	err := ur.db.Delete(USER_INDEX_NAME, id)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) List(query map[string]string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error) {
	umr, err := ur.db.List(USER_INDEX_NAME, query, limit, offset, orderBy, order)
	if err != nil {
		return nil, err
	}
	return umr, nil
}

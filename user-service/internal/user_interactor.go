package internal

import (
	"github.com/Toorreess/laWiki/user-service/internal/model"
)

type IUserInteractor interface {
	Create(um *model.User) (*model.User, error)
	Get(id string) (*model.User, error)
	Update(id string, updates map[string]interface{}) (*model.User, error)
	Delete(id string) error
	List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error)
}

type userInteractor struct {
	UserRepository IUserRepository
	UserPresenter  IUserPresenter
}

func NewUserInteractor(r IUserRepository, p IUserPresenter) IUserInteractor {
	return &userInteractor{
		UserRepository: r,
		UserPresenter:  p,
	}
}

func (ui *userInteractor) Create(wm *model.User) (*model.User, error) {
	result, err := ui.UserRepository.Create(wm)
	if err != nil {
		return nil, err
	}
	return ui.UserPresenter.ResponseUser(result), nil
}

func (ui *userInteractor) Get(id string) (*model.User, error) {
	um, err := ui.UserRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return ui.UserPresenter.ResponseUser(um), nil
}

func (ui *userInteractor) Update(id string, updates map[string]interface{}) (*model.User, error) {
	um, err := ui.UserRepository.Update(id, updates)

	if err != nil {
		return nil, err
	}
	return ui.UserPresenter.ResponseUser(um), nil
}

func (ui *userInteractor) Delete(id string) error {
	return ui.UserRepository.Delete(id)
}

func (ui *userInteractor) List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error) {
	result, err := ui.UserRepository.List(query, limit, offset, orderBy, order)
	if err != nil {
		return nil, err
	}

	return ui.UserPresenter.ResponseUsers(result, limit, offset), nil
}

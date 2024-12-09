package internal

import (
	"fmt"

	"github.com/Toorreess/laWiki/user-service/internal/model"
)

type IUserInteractor interface {
	Create(um *model.User) (*model.User, error)
	Get(id string) (*model.User, error)
	Update(id string, updates map[string]interface{}) (*model.User, error)
	Delete(id string) error
	List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error)

	AddNotification(userID string, notification *model.Notification) error
	ReadNotification(userID, notificationID string) error
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

func (ui *userInteractor) AddNotification(userID string, notification *model.Notification) error {
	userMap, err := ui.UserRepository.Get(userID)
	if err != nil {
		return err
	}
	user := ui.UserPresenter.ResponseUser(userMap)

	user.Notifications = append(user.Notifications, *notification)

	_, err = ui.UserRepository.Update(userID, map[string]interface{}{
		"notifications": user.Notifications,
	})

	return err
}

func (ui *userInteractor) ReadNotification(userID string, notificationID string) error {
	userMap, err := ui.UserRepository.Get(userID)
	if err != nil {
		return err
	}
	user := ui.UserPresenter.ResponseUser(userMap)

	notificationList := user.Notifications
	found := false
	for _, notif := range notificationList {
		if notif.ID == notificationID {
			notif.Read = true
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("notification not found")
	}

	user.Notifications = notificationList

	_, err = ui.UserRepository.Update(userID, map[string]interface{}{
		"notifications": notificationList,
	})

	return err
}

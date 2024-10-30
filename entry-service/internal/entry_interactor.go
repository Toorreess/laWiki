package internal

import "github.com/Toorreess/laWiki/entry-service/internal/model"

type IEntryInteractor interface {
	Create(wm *model.Entry) (*model.Entry, error)
	Get(id string) (*model.Entry, error)
	Update(id string, updates map[string]interface{}) (*model.Entry, error)
	Delete(id string) error
	List(query string, limit, offset int) (map[string]interface{}, error)
}

type entryInteractor struct {
	EntryRepository IEntryRepository
	EntryPresenter  IEntryPresenter
}

func NewEntryInteractor(r IEntryRepository, p IEntryPresenter) IEntryInteractor {
	return &entryInteractor{
		EntryRepository: r,
		EntryPresenter:  p,
	}
}

func (wi *entryInteractor) Create(em *model.Entry) (*model.Entry, error) {
	result, err := wi.EntryRepository.Create(em)
	if err != nil {
		return nil, err
	}
	return wi.EntryPresenter.ResponseEntry(result), nil
}

func (wi *entryInteractor) Get(id string) (*model.Entry, error) {
	cm, err := wi.EntryRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return wi.EntryPresenter.ResponseEntry(cm), nil
}

func (wi *entryInteractor) Update(id string, updates map[string]interface{}) (*model.Entry, error) {
	cm, err := wi.EntryRepository.Update(id, updates)
	if err != nil {
		return nil, err
	}
	return wi.EntryPresenter.ResponseEntry(cm), nil
}

func (wi *entryInteractor) Delete(id string) error {
	return wi.EntryRepository.Delete(id)
}

func (wi *entryInteractor) List(query string, limit int, offset int) (map[string]interface{}, error) {
	result, err := wi.EntryRepository.List(query, limit, offset)
	if err != nil {
		return nil, err
	}
	return wi.EntryPresenter.ResponseEntries(result, limit, offset), nil
}

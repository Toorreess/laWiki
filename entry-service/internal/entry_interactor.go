package internal

import "github.com/Toorreess/laWiki/entry-service/internal/model"

type IEntryInteractor interface {
	Create(m *model.Entry) (*model.Entry, error)
	Get(id string) (*model.Entry, error)
	Update(id string, updates map[string]interface{}) (*model.Entry, error)
	Delete(id string) error
	List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error)
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

func (ei *entryInteractor) Create(em *model.Entry) (*model.Entry, error) {
	result, err := ei.EntryRepository.Create(em)
	if err != nil {
		return nil, err
	}
	return ei.EntryPresenter.ResponseEntry(result), nil
}

func (ei *entryInteractor) Get(id string) (*model.Entry, error) {
	em, err := ei.EntryRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return ei.EntryPresenter.ResponseEntry(em), nil
}

func (ei *entryInteractor) Update(id string, updates map[string]interface{}) (*model.Entry, error) {
	em, err := ei.EntryRepository.Update(id, updates)
	if err != nil {
		return nil, err
	}
	return ei.EntryPresenter.ResponseEntry(em), nil
}

func (ei *entryInteractor) Delete(id string) error {
	return ei.EntryRepository.Delete(id)
}

func (ei *entryInteractor) List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error) {
	result, err := ei.EntryRepository.List(query, limit, offset, orderBy, order)
	if err != nil {
		return nil, err
	}
	return ei.EntryPresenter.ResponseEntries(result, limit, offset), nil
}

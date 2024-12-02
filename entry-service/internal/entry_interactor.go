package internal

import (
	"mime/multipart"

	"firebase.google.com/go/v4/storage"
	"github.com/Toorreess/laWiki/entry-service/internal/model"
)

type IEntryInteractor interface {
	Create(entryData map[string]string, file multipart.File, storageClient *storage.Client) (*model.Entry, error)
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

func (ei *entryInteractor) Create(entryData map[string]string, file multipart.File, storageClient *storage.Client) (*model.Entry, error) {
	var em *model.Entry

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

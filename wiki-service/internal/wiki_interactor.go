package internal

import (
	"github.com/Toorreess/laWiki/wiki-service/internal/model"
)

type IWikiInteractor interface {
	Create(wm *model.Wiki) (*model.Wiki, error)
	Get(id string) (*model.Wiki, error)
	Update(id string, updates map[string]interface{}) (*model.Wiki, error)
	Delete(id string) error
	List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error)
}

type wikiInteractor struct {
	WikiRepository IWikiRepository
	WikiPresenter  IWikiPresenter
}

func NewWikiInteractor(r IWikiRepository, p IWikiPresenter) IWikiInteractor {
	return &wikiInteractor{
		WikiRepository: r,
		WikiPresenter:  p,
	}
}

func (wi *wikiInteractor) Create(wm *model.Wiki) (*model.Wiki, error) {
	result, err := wi.WikiRepository.Create(wm)
	if err != nil {
		return nil, err
	}
	return wi.WikiPresenter.ResponseWiki(result), nil
}

func (wi *wikiInteractor) Get(id string) (*model.Wiki, error) {
	cm, err := wi.WikiRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return wi.WikiPresenter.ResponseWiki(cm), nil
}

func (wi *wikiInteractor) Update(id string, updates map[string]interface{}) (*model.Wiki, error) {
	cm, err := wi.WikiRepository.Update(id, updates)
	if err != nil {
		return nil, err
	}
	return wi.WikiPresenter.ResponseWiki(cm), nil
}

func (wi *wikiInteractor) Delete(id string) error {
	return wi.WikiRepository.Delete(id)
}

func (wi *wikiInteractor) List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error) {
	result, err := wi.WikiRepository.List(query, limit, offset, orderBy, order)
	if err != nil {
		return nil, err
	}

	return wi.WikiPresenter.ResponseWikis(result, limit, offset), nil
}

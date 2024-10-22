package internal

import "github.com/Toorreess/laWiki/wiki-service/internal/model"

type IWikiInteractor interface {
	Create(wm *model.Wiki) (*model.Wiki, error)
	Read(id string) (*model.Wiki, error)
	Update(id string, w map[string]interface{}) (*model.Wiki, error)
	Delete(id string) error

	List(query string, limit, offset int, orderBy, order string) ([]map[string]interface{}, error)
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

func (wi *wikiInteractor) Read(id string) (*model.Wiki, error) {
	panic("unimplemented")
}

func (wi *wikiInteractor) Update(id string, w map[string]interface{}) (*model.Wiki, error) {
	panic("unimplemented")
}

func (wi *wikiInteractor) Delete(id string) error {
	panic("unimplemented")
}

func (wi *wikiInteractor) List(query string, limit int, offset int, orderBy string, order string) ([]map[string]interface{}, error) {
	panic("unimplemented")
}

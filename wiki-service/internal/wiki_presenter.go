package internal

import (
	"encoding/json"

	"github.com/Toorreess/laWiki/wiki-service/internal/model"
)

type IWikiPresenter interface {
	ResponseWiki(wikiMap map[string]interface{}) *model.Wiki
	ResponseWikis(w []map[string]interface{}) map[string]interface{}
}

type wikiPresenter struct{}

func NewWikiPresenter() IWikiPresenter {
	return &wikiPresenter{}
}

func (wp *wikiPresenter) ResponseWiki(wikiMap map[string]interface{}) *model.Wiki {
	jsonBody, _ := json.Marshal(wikiMap)
	wm := model.Wiki{}
	json.Unmarshal(jsonBody, &wm)
	return &wm
}

// ResponseWikis implements IWikiPresenter.
func (wp *wikiPresenter) ResponseWikis(w []map[string]interface{}) map[string]interface{} {
	panic("unimplemented")
}

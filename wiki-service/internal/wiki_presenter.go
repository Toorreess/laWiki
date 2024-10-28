package internal

import (
	"encoding/json"

	"github.com/Toorreess/laWiki/wiki-service/internal/model"
)

type IWikiPresenter interface {
	ResponseWiki(wikiMap map[string]interface{}) *model.Wiki
	ResponseWikis(wikiList []map[string]interface{}, limit, offset int) map[string]interface{}
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

func (wp *wikiPresenter) ResponseWikis(wikiList []map[string]interface{}, limit, offset int) map[string]interface{} {
	resultMap := make(map[string]interface{})
	var results []*model.Wiki

	for _, wikiMap := range wikiList {
		jsonBody, _ := json.Marshal(wikiMap)
		wm := model.Wiki{}
		json.Unmarshal(jsonBody, &wm)
		results = append(results, &wm)
	}

	resultMap["items"] = results
	if len(results) == limit {
		resultMap["next_offset"] = offset + limit
	} else {
		resultMap["next_offset"] = nil
	}
	if (offset - limit) < 0 {
		resultMap["previous_offset"] = nil
	} else {
		resultMap["previous_offset"] = offset - limit
	}

	resultMap["offset"] = offset
	resultMap["limit"] = limit
	resultMap["total"] = len(results)
	return resultMap
}

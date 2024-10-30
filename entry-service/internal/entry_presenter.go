package internal

import (
	"encoding/json"

	"github.com/Toorreess/laWiki/entry-service/internal/model"
)

type IEntryPresenter interface {
	ResponseEntry(entryMap map[string]interface{}) *model.Entry
	ResponseEntries(entryList []map[string]interface{}, limit, offset int) map[string]interface{}
}

type entryPresenter struct{}

func NewEntryPresenter() IEntryPresenter {
	return &entryPresenter{}
}

func (wp *entryPresenter) ResponseEntry(wikiMap map[string]interface{}) *model.Entry {
	jsonBody, _ := json.Marshal(wikiMap)
	wm := model.Entry{}
	json.Unmarshal(jsonBody, &wm)
	return &wm
}

func (wp *entryPresenter) ResponseEntries(wikiList []map[string]interface{}, limit, offset int) map[string]interface{} {
	resultMap := make(map[string]interface{})
	var results []*model.Entry

	for _, wikiMap := range wikiList {
		jsonBody, _ := json.Marshal(wikiMap)
		wm := model.Entry{}
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

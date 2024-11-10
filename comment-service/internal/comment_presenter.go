package internal

import (
	"encoding/json"

	"github.com/Toorreess/laWiki/comment-service/internal/model"
)

type ICommentPresenter interface {
	ResponseComment(commentMap map[string]interface{}) *model.Comment
	ResponseComments(commentList []map[string]interface{}, limit, offset int) map[string]interface{}
}

type commentPresenter struct{}

func NewCommentPresenter() ICommentPresenter {
	return &commentPresenter{}
}

func (wp *commentPresenter) ResponseComment(commentMap map[string]interface{}) *model.Comment {
	jsonBody, _ := json.Marshal(commentMap)
	wm := model.Comment{}
	json.Unmarshal(jsonBody, &wm)
	return &wm
}

func (wp *commentPresenter) ResponseComments(commentList []map[string]interface{}, limit, offset int) map[string]interface{} {
	resultMap := make(map[string]interface{})
	var results []*model.Comment

	for _, commentList := range commentList {
		jsonBody, _ := json.Marshal(commentList)
		wm := model.Comment{}
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

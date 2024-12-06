package internal

import (
	"encoding/json"

	"github.com/Toorreess/laWiki/user-service/internal/model"
)

type IUserPresenter interface {
	ResponseUser(userMap map[string]interface{}) *model.User
	ResponseUsers(userList []map[string]interface{}, limit, offset int) map[string]interface{}
}

type userPresenter struct{}

func NewUserPresenter() IUserPresenter {
	return &userPresenter{}
}

func (wp *userPresenter) ResponseUser(userMap map[string]interface{}) *model.User {
	jsonBody, _ := json.Marshal(userMap)
	um := model.User{}
	json.Unmarshal(jsonBody, &um)
	return &um
}

func (wp *userPresenter) ResponseUsers(userList []map[string]interface{}, limit, offset int) map[string]interface{} {
	resultMap := make(map[string]interface{})
	var results []*model.User

	for _, userMap := range userList {
		jsonBody, _ := json.Marshal(userMap)
		um := model.User{}
		json.Unmarshal(jsonBody, &um)
		results = append(results, &um)
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

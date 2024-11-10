package internal

import "github.com/Toorreess/laWiki/comment-service/internal/model"

type ICommentInteractor interface {
	Create(m *model.Comment) (*model.Comment, error)
	Get(id string) (*model.Comment, error)
	Update(id string, updates map[string]interface{}) (*model.Comment, error)
	Delete(id string) error
	List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error)
}

type commentInteractor struct {
	commentRepository ICommentRepository
	commentPresenter  ICommentPresenter
}

func NewcommentInteractor(r ICommentRepository, p ICommentPresenter) ICommentInteractor {
	return &commentInteractor{
		commentRepository: r,
		commentPresenter:  p,
	}
}

func (ei *commentInteractor) Create(em *model.Comment) (*model.Comment, error) {
	result, err := ei.commentRepository.Create(em)
	if err != nil {
		return nil, err
	}
	return ei.commentPresenter.ResponseComment(result), nil
}

func (ei *commentInteractor) Get(id string) (*model.Comment, error) {
	em, err := ei.commentRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return ei.commentPresenter.ResponseComment(em), nil
}

func (ei *commentInteractor) Update(id string, updates map[string]interface{}) (*model.Comment, error) {
	em, err := ei.commentRepository.Update(id, updates)
	if err != nil {
		return nil, err
	}
	return ei.commentPresenter.ResponseComment(em), nil
}

func (ei *commentInteractor) Delete(id string) error {
	return ei.commentRepository.Delete(id)
}

func (ei *commentInteractor) List(query map[string]string, limit, offset int, orderBy, order string) (map[string]interface{}, error) {
	result, err := ei.commentRepository.List(query, limit, offset, orderBy, order)
	if err != nil {
		return nil, err
	}
	return ei.commentPresenter.ResponseComments(result, limit, offset), nil
}

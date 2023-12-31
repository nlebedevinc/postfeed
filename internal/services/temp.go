package services

import (
	"context"

	"github.com/nlebedevinc/postfeed/internal/models"
)

type PostService struct {
	Db Redis[models.Post]
}

func (s PostService) Get(ctx context.Context, id string) (models.Post, error) {
	return s.Db.Get(ctx, "post:"+id)
}

func (s PostService) Save(ctx context.Context, post models.Post) (models.Post, error) {
	post.GenerateId()
	if err := s.Db.Save(ctx, post); err != nil {
		return post, err
	}

	return post, nil
}

func (s PostService) MGet(ctx context.Context, uid ...string) ([]models.Post, error) {
	ids := make([]string, len(uid))
	for i, s := range uid {
		ids[i] = "post:" + s
	}
	return s.Db.MGet(ctx, ids...)
}

func NewPostService(db Redis[models.Post]) PostService {
	s := PostService{Db: db}
	return s
}

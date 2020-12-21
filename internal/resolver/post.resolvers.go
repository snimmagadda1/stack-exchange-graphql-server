package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/snimmagadda1/graphql-api/generated"
	"github.com/snimmagadda1/graphql-api/model"
)

func (r *postResolver) Answers(ctx context.Context, obj *model.Post) ([]*model.Post, error) {
	answers := []*model.Post{}
	if obj.AnswerCount == nil || *obj.AnswerCount == 0 {
		return answers, nil
	}
	r.DB.Where("ParentId = ?", obj.ID).Order("Score desc").Find(&answers)
	return answers, nil
}

func (r *postResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	if obj.CommentCount == nil || *obj.CommentCount == 0 {
		return comments, nil
	}
	r.DB.Where("PostId = ?", obj.ID).Order("creationDate asc").Find(&comments)
	return comments, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }

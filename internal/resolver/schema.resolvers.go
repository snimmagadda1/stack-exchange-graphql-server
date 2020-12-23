package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/snimmagadda1/graphql-api/generated"
	"github.com/snimmagadda1/graphql-api/model"
)

func (r *queryResolver) GetPostHistory(ctx context.Context, id int) (*model.PostHistory, error) {
	var ph model.PostHistory
	r.DB.First(&ph, id)

	return &ph, nil
}

func (r *queryResolver) GetBadge(ctx context.Context, id int) (*model.Badge, error) {
	var b model.Badge
	r.DB.First(&b, id)

	return &b, nil
}

func (r *queryResolver) GetVote(ctx context.Context, id int) (*model.Vote, error) {
	var v model.Vote
	r.DB.First(&v, id)

	return &v, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

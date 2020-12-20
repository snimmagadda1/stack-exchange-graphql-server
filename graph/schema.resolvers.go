package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/snimmagadda1/graphql-api/graph/generated"
	"github.com/snimmagadda1/graphql-api/graph/model"
)

func (r *queryResolver) GetUser(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	r.DB.First(&user, id)

	return &user, nil
}

func (r *queryResolver) GetPost(ctx context.Context, id int) (*model.Post, error) {
	var post model.Post
	r.DB.First(&post, id)

	return &post, nil
}

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

func (r *queryResolver) GetComment(ctx context.Context, id int) (*model.Comment, error) {
	var com model.Comment
	r.DB.First(&com, id)

	return &com, nil
}

func (r *queryResolver) GetVote(ctx context.Context, id int) (*model.Vote, error) {
	var v model.Vote
	r.DB.First(&v, id)

	return &v, nil
}

func (r *queryResolver) AllPostsCursor(ctx context.Context, first *int, after *string, where *model.PostsWhere) (*model.PostsCursor, error) {
	if first != nil && *first < 0 {
		logrus.Panic(fmt.Errorf("first must be positive"))
	}
	// field to sort by
	field := "Id"
	var order *model.PostsOrderBy
	if where != nil {
		order = where.Order
		switch key := *order.Field; key {
		case model.PostsSortFieldsOpaqueKey:
			field = "Id"
		case model.PostsSortFieldsActivity:
			field = "LastActivityDate"
		case model.PostsSortFieldsCreation:
			field = "CreationDate"
		case model.PostsSortFieldsVotes:
			field = "Score"
		default:
			field = "Id"
		}

	}

	// query start
	start := 0
	if after != nil {
		b, err := strconv.Atoi(*after)
		if err != nil {
			return nil, err
		}
		start = b // might need to be + 1
	}

	var total int64
	r.DB.Model(&model.Post{}).Count(&total)
	logrus.Infof("Total count of posts in db found %d", total)
	limit := 10
	if first != nil {
		limit = *first
	}
	// select * from posts where id = after order by id DESC limit first
	posts := []model.Post{}
	r.DB.Where(field+" > ?", start).Limit(limit).Find(&posts).Order(field + " desc")

	// create edges from results
	postEdges := []*model.PostEdge{}
	for i := range posts {
		toAdd := &model.PostEdge{Cursor: posts[i].ID, Node: &posts[i]}
		postEdges = append(postEdges, toAdd)
	}

	// should limt = first here...?
	pageInfo := model.PageInfo{
		HasNextPage:     int64(start+limit) < total,
		HasPreviousPage: start > 0,
	}

	return &model.PostsCursor{
		Edges:    postEdges,
		PageInfo: &pageInfo,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }

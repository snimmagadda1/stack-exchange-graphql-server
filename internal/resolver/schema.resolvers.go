package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/snimmagadda1/graphql-api/generated"
	"github.com/snimmagadda1/graphql-api/internal/util"
	"github.com/snimmagadda1/graphql-api/model"
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
	// prep query - field to sort by - not currently implemented
	field := "Id"
	var order *model.PostsOrderBy
	if where != nil {
		order = where.Order
		switch key := *order.Field; key {
		case model.PostsSortFieldsActivity:
			field = "LastActivityDate"
		case model.PostsSortFieldsCreation:
			field = "CreationDate"
		case model.PostsSortFieldsVotes:
			field = "Score"
		default:
			field = "Id"
		}
		logrus.Infof("Using field %s with order type %s", field, order.Field)
	}

	// prep query - query start
	start := int64(0)
	if after != nil {
		decoded, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		start, err = strconv.ParseInt(string(decoded), 10, 0)
		if err != nil {
			return nil, err
		}
	}

	// prep query  - limit
	limit := 10
	if first != nil {
		limit = *first
	}
	if limit > 50 {
		logrus.Warn("Limit requested exceeds maximum 50")
		limit = 50
	}

	// result metadata
	var total int64
	r.DB.Model(&model.Post{}).Count(&total)
	logrus.Infof("Total count of posts in db found %d", total)

	// select * from posts where id = after order by id DESC limit first
	posts := []model.Post{}
	r.DB.Where(field+" > ?", start).Limit(limit).Find(&posts).Order(field + " desc")

	// create edges from results
	postEdges := []*model.PostEdge{}
	for i := range posts {
		cursor := util.GetCursor(posts[i], field)
		toAdd := &model.PostEdge{Cursor: cursor, Node: &posts[i]}
		postEdges = append(postEdges, toAdd)
	}

	// should limt = first here...?
	pageInfo := model.PageInfo{
		HasNextPage:     start+int64(limit) < total,
		HasPreviousPage: start > 0,
	}

	return &model.PostsCursor{
		Edges:    postEdges,
		PageInfo: &pageInfo,
	}, nil
}

func (r *queryResolver) AllCommentsCursor(ctx context.Context, first *int, after *string) (*model.CommentsCursor, error) {
	if first != nil && *first < 0 {
		logrus.Panic(fmt.Errorf("first must be positive"))
	}
	// prep query - field to sort by - not currently implemented
	field := "Id"

	// prep query - query start
	start := int64(0)
	if after != nil {
		decoded, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		start, err = strconv.ParseInt(string(decoded), 10, 0)
		if err != nil {
			return nil, err
		}
	}

	// prep query  - limit
	limit := 10
	if first != nil {
		limit = *first
	}
	if limit > 100 {
		logrus.Warn("Limit requested exceeds maximum 100")
		limit = 100
	}

	// result metadata
	var total int64
	r.DB.Model(&model.Comment{}).Count(&total)
	logrus.Infof("Total count of comments in db found %d", total)

	comments := []model.Comment{}
	r.DB.Where(field+" > ?", start).Limit(limit).Find(&comments).Order(field + " desc")

	// create edges from results
	edges := []*model.CommentEdge{}
	for i := range comments {
		cursor := base64.StdEncoding.EncodeToString([]byte(comments[i].ID))
		toAdd := &model.CommentEdge{Cursor: cursor, Node: &comments[i]}
		edges = append(edges, toAdd)
	}

	// should limt = first here...?
	pageInfo := model.PageInfo{
		HasNextPage:     start+int64(limit) < total,
		HasPreviousPage: start > 0,
	}

	return &model.CommentsCursor{
		Edges:    edges,
		PageInfo: &pageInfo,
	}, nil
}

func (r *queryResolver) AllUsersCursor(ctx context.Context, first *int, after *string) (*model.UsersCursor, error) {
	if first != nil && *first < 0 {
		logrus.Panic(fmt.Errorf("first must be positive"))
	}
	// prep query - field to sort by - not currently implemented
	field := "Id"

	// prep query - query start
	start := int64(0)
	if after != nil {
		decoded, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		start, err = strconv.ParseInt(string(decoded), 10, 0)
		if err != nil {
			return nil, err
		}
	}

	// prep query  - limit
	limit := 10
	if first != nil {
		limit = *first
	}
	if limit > 100 {
		logrus.Warn("Limit requested exceeds maximum 100")
		limit = 100
	}

	// result metadata
	var total int64
	r.DB.Model(&model.User{}).Count(&total)
	logrus.Infof("Total count of users in db found %d", total)

	users := []model.User{}
	r.DB.Where(field+" > ?", start).Limit(limit).Find(&users).Order(field + " desc")

	// create edges from results
	edges := []*model.UserEdge{}
	for i := range users {
		cursor := base64.StdEncoding.EncodeToString([]byte(users[i].ID))
		toAdd := &model.UserEdge{Cursor: cursor, Node: &users[i]}
		edges = append(edges, toAdd)
	}

	// should limt = first here...?
	pageInfo := model.PageInfo{
		HasNextPage:     start+int64(limit) < total,
		HasPreviousPage: start > 0,
	}

	return &model.UsersCursor{
		Edges:    edges,
		PageInfo: &pageInfo,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
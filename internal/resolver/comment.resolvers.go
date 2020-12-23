package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/snimmagadda1/graphql-api/internal/dal"
	"github.com/snimmagadda1/graphql-api/model"
)

func (r *queryResolver) GetComment(ctx context.Context, id int) (*model.Comment, error) {
	var com model.Comment
	r.DB.First(&com, id)

	return &com, nil
}

func (r *queryResolver) AllCommentsCursor(ctx context.Context, first *int, after *string) (*model.CommentsCursor, error) {
	if first != nil && *first < 0 {
		logrus.Panic(fmt.Errorf("first must be positive"))
	}
	// prep query sort and bounds
	field := "Id"
	start, limit, err := dal.GetQueryBounds(first, after)
	if err != nil {
		return nil, err
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

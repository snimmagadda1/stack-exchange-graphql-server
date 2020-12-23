package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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

func (r *queryResolver) AllCommentsCursor(ctx context.Context, first *int, after *string, where *model.CommentsWhere) (*model.CommentsCursor, error) {
	if first != nil && *first < 0 {
		logrus.Panic(fmt.Errorf("first must be positive"))
	}
	comments := []model.Comment{}
	// prep query sort and bounds
	field := "CreationDate"
	order := "desc"
	start, limit, err := dal.GetQueryBounds(first, after)
	if err != nil {
		return nil, err
	}

	// result metadata
	var total int64
	r.DB.Model(&model.Comment{}).Count(&total)
	logrus.Infof("Total count of comments in db found %d", total)

	if where != nil {
		sortInfo := where.Order
		if sortInfo.Order != nil {
			order = string(*sortInfo.Order)
		}
		switch key := *sortInfo.Field; key {
		case model.CommentSortFieldsCreation:
			field = "CreationDate"
		case model.CommentSortFieldsVotes:
			field = "Score"
		default:
			field = "CreationDate"
		}
	}

	if after != nil {
		endSQL := ""
		r.DB.Raw("SELECT "+field+" FROM comments where id = ?", start).Scan(&endSQL)
		r.DB.Limit(limit).Where(field+" <= ?", endSQL).Where("id != ? ", start).Order(field + " " + order).Order("id desc").Find(&comments)

		// create edges from results
		var edges []*model.CommentEdge
		for i := range comments {
			edges = append(edges, comments[i].CommentEdge())
		}
		pageInfo := model.PageInfo{
			HasNextPage:     start+int64(limit) < total,
			HasPreviousPage: start > 0,
		}

		return &model.CommentsCursor{
			Edges:    edges,
			PageInfo: &pageInfo,
		}, nil

	}

	r.DB.Limit(limit).Order(field + " " + order).Order("id desc").Find(&comments)

	// create edges from results
	var edges []*model.CommentEdge
	for i := range comments {
		edges = append(edges, comments[i].CommentEdge())
	}
	pageInfo := model.PageInfo{
		HasNextPage:     start+int64(limit) < total,
		HasPreviousPage: start > 0,
	}

	return &model.CommentsCursor{
		Edges:    edges,
		PageInfo: &pageInfo,
	}, nil
}

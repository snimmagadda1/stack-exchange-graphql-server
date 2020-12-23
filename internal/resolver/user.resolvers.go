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

func (r *queryResolver) GetUser(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	r.DB.First(&user, id)

	return &user, nil
}

func (r *queryResolver) AllUsersCursor(ctx context.Context, first *int, after *string, where *model.UsersWhere) (*model.UsersCursor, error) {
	if first != nil && *first < 0 {
		logrus.Panic(fmt.Errorf("first must be positive"))
	}
	users := []model.User{}
	// prep query sort and bounds
	field := "Reputation"
	order := "desc"
	start, limit, err := dal.GetQueryBounds(first, after)
	if err != nil {
		return nil, err
	}
	// result metadata
	var total int64
	r.DB.Model(&model.User{}).Count(&total)
	logrus.Infof("Total count of users in db found %d", total)

	if where != nil {
		sortInfo := where.Order
		if sortInfo.Order != nil {
			order = string(*sortInfo.Order)
		}
		switch key := *sortInfo.Field; key {
		case model.UsersSortFieldsReputation:
			field = "Reputation"
		case model.UsersSortFieldsCreation:
			field = "CreationDate"
		case model.UsersSortFieldsName:
			field = "DisplayName"
		default:
			field = "Score"
		}
	}

	if after != nil {
		endSQL := ""
		r.DB.Raw("SELECT "+field+" FROM users where id = ?", start).Scan(&endSQL)
		r.DB.Limit(limit).Where(field+" <= ?", endSQL).Where("id != ? ", start).Order(field + " " + order).Order("id desc").Find(&users)

		// create edges from results
		var edges []*model.UserEdge
		for i := range users {
			edges = append(edges, users[i].UserEdge())
		}
		pageInfo := model.PageInfo{
			HasNextPage:     start+int64(limit) < total,
			HasPreviousPage: start > 0,
		}

		return &model.UsersCursor{
			Edges:    edges,
			PageInfo: &pageInfo,
		}, nil

	}

	r.DB.Limit(limit).Order(field + " " + order).Order("id desc").Find(&users)

	// create edges from results
	var edges []*model.UserEdge
	for i := range users {
		edges = append(edges, users[i].UserEdge())
	}
	pageInfo := model.PageInfo{
		HasNextPage:     start+int64(limit) < total,
		HasPreviousPage: start > 0,
	}

	return &model.UsersCursor{
		Edges:    edges,
		PageInfo: &pageInfo,
	}, nil
}

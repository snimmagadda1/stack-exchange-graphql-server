package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/snimmagadda1/graphql-api/generated"
	"github.com/snimmagadda1/graphql-api/internal/dal"
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

func (r *queryResolver) GetPost(ctx context.Context, id int) (*model.Post, error) {
	var post model.Post
	r.DB.First(&post, id)

	return &post, nil
}

func (r *queryResolver) AllPostsCursor(ctx context.Context, first *int, after *string, where *model.PostsWhere) (*model.PostsCursor, error) {
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
	r.DB.Model(&model.Post{}).Count(&total)
	logrus.Infof("Total count of posts in db found %d", total)

	// select * from posts where id = after order by id DESC limit first
	posts := []model.Post{}
	if where == nil || where.Order.Field == nil {
		r.DB.Where(field+" > ?", start).Limit(limit).Order(field + " desc").Find(&posts)

		// create edges from results
		var postEdges []*model.PostEdge
		for i := range posts {
			postEdges = append(postEdges, posts[i].PostEdge())
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

	// Case client order... back too our roots SQL
	sortInfo := where.Order
	order := "desc"
	if sortInfo.Order != nil {
		order = string(*sortInfo.Order)
	}

	switch key := *sortInfo.Field; key {
	case model.PostsSortFieldsActivity:
		field = "LastActivityDate"
	case model.PostsSortFieldsCreation:
		field = "CreationDate"
	case model.PostsSortFieldsVotes:
		field = "Score"
	default:
		field = "Id"
	}

	if after != nil {
		endSQL := ""
		r.DB.Raw("SELECT "+field+" FROM posts where id = ?", start).Scan(&endSQL)
		r.DB.Limit(limit).Where(field+" <= ?", endSQL).Where("id != ? ", start).Order(field + " " + order).Order("id desc").Find(&posts)

		// create edges from results
		var postEdges []*model.PostEdge
		for i := range posts {
			postEdges = append(postEdges, posts[i].PostEdge())
		}
		pageInfo := model.PageInfo{
			HasNextPage:     start+int64(limit) < total,
			HasPreviousPage: start > 0,
		}

		return &model.PostsCursor{
			Edges:    postEdges,
			PageInfo: &pageInfo,
		}, nil

	}

	r.DB.Limit(limit).Order(field + " " + order).Order("id desc").Find(&posts)

	// create edges from results
	var postEdges []*model.PostEdge
	for i := range posts {
		postEdges = append(postEdges, posts[i].PostEdge())
	}
	pageInfo := model.PageInfo{
		HasNextPage:     start+int64(limit) < total,
		HasPreviousPage: start > 0,
	}

	return &model.PostsCursor{
		Edges:    postEdges,
		PageInfo: &pageInfo,
	}, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }

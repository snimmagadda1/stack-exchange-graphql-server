package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/snimmagadda1/graphql-api/graph/model"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users []*model.User
	DB    *gorm.DB
}

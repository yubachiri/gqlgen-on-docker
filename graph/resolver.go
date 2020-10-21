package graph

import (
	"go-graphql-api-sample/graph/model"

	"xorm.io/xorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver リゾルバです
type Resolver struct {
	todos []*model.Todo
	users []*model.User
	todo  *model.Todo
	DB    *xorm.Engine
}

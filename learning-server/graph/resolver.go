package graph

import "learning-server/graph/models"

type Resolver struct {
	todos     []*models.Todo
	questions []*models.Question
}

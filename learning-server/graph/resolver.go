//go:generate go run generate.go

package graph

import "learning-server/graph/models"

type Resolver struct {
	questions []*models.Question
}

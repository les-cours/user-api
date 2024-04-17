package resolvers

import (
	"github.com/les-cours/user-api/api/users"
	"github.com/les-cours/user-api/graph"
)

type Resolver struct {
	UserClient users.UserServiceClient
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

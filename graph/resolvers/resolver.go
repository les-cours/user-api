package resolvers

import (
	"github.com/les-cours/user-api/graph"
	bookApi "github.com/les-cours/user-api/protobuf/book"
)

type Resolver struct {
	BookClient bookApi.BookServiceClient
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

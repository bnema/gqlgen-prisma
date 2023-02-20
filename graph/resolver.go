package graph

import (
	"graphql-prisma-api/prisma/db"
)

// Set Prisma client := db.NewClient() in resolver.go
// This is the resolver for the Query type
type Resolver struct {
	Prisma *db.PrismaClient
}

// // This is the resolver for the User type
// func (r *Resolver) User() UserResolver {
// 	return &userResolver{r}
// }

// // This is the resolver for the Post type
// func (r *Resolver) Post() PostResolver {
// 	return &postResolver{r}
// }

// // This is the resolver for the Mutation type
// func (r *Resolver) Mutation() MutationResolver {
// 	return &mutationResolver{r}
// }

// // This is the resolver for the Query type
// func (r *Resolver) Query() QueryResolver {
// 	return &queryResolver{r}
// }

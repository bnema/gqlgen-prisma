package graph

import (
	"context"
	"graphql-prisma-api/prisma/db"

	"github.com/99designs/gqlgen/graphql"
)

// go run github.com/99designs/gqlgen generate
// Set Prisma client := db.NewClient() in resolver.go
// This is the resolver for the Query type
type Resolver struct {
	Prisma *db.PrismaClient
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
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

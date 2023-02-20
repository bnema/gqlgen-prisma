package main

import (
	"context"
	"encoding/json"
	"fmt"

	"graphql-prisma-api/prisma/db"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	createUser, err := client.User.CreateOne(
		db.User.Email.Set("hi@bnema.dev"),
		db.User.Name.Set("Brice"),
		// placeholder image
		db.User.Image.Set("https://cdn.bnema.dev/rocknmeet/placeholder-user-120x120.jpg"),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createUser, "", "  ")
	fmt.Printf("created user: %s\n", result)

	// find a single user
	user, err := client.User.FindUnique(
		db.User.ID.Equals(createUser.ID),
	).Exec(ctx)
	if err != nil {
		return err
	}

	// // create a post and link it to the user
	// createdPost, err := client.Post.CreateOne(
	// 	db.Post.Author.Link(
	// 		db.User.ID.Equals(createUser.ID),
	// 	),
	// 	db.Post.Title.Set("Hi from Prisma!"),
	// 	db.Post.Content.Set("This is my first post"),
	// 	db.Post.Published.Set(true),
	// ).Exec(ctx)
	// if err != nil {
	// 	return err
	// }

	// result, _ = json.MarshalIndent(createdPost, "", "  ")
	// fmt.Printf("created post: %s\n", result)

	// // find a single post
	// post, err := client.Post.FindUnique(
	// 	db.Post.ID.Equals(createdPost.ID),
	// ).Exec(ctx)
	// if err != nil {
	// 	return err
	// }

	// result, _ = json.MarshalIndent(post, "", "  ")
	// fmt.Printf("post: %s\n", result)

	// for optional/nullable values, you need to check the function and create two return values
	// `desc` is a string, and `ok` is a bool whether the record is null or not. If it's null,
	// `ok` is false, and `desc` will default to Go's default values; in this case an empty string (""). Otherwise,
	// `ok` is true and `desc` will be "my description".
	name, ok := user.Name()
	if !ok {
		return fmt.Errorf("user name is null")
	}

	fmt.Printf("user name: %s\n", name)

	return nil
}

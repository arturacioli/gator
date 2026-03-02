package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/arturacioli/gator/internal/database"
)

func HandlerBrowse(s* State, cmd command, user *database.User) error{
	limitParameter := 2
	var err error
	if len(cmd.arguments) > 0{
		limitParameter,err = strconv.Atoi(cmd.arguments[0]) 
		if err != nil{
			return err
		}
	}
	getPostsParam := database.GetPostsForUserParams{
			UserID: user.ID,	
			Limit: int32(limitParameter),	
	}
	posts, err := s.db.GetPostsForUser(context.Background(),getPostsParam)
	if err != nil {
		return err 
	}
	
	fmt.Println("Posts:")
	for _, post := range posts{
		fmt.Printf("%s\n", post.Title.String)
		fmt.Printf("Published: %s\n", post.PublishedAt)
		fmt.Printf("URL: %s\n\n", post.Url)
	}

	return nil
}

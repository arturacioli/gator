package main

import (
	"context"
	"fmt"

	"github.com/arturaciolii/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd command, usr *database.User) error{
	args := cmd.arguments
	if len(args) == 0{
		return fmt.Errorf("Not enough arguments") 
	}
	feed, err := s.db.GetFeed(context.Background(),args[0])
	if err != nil{
		return err
	}
	deleteParams := database.DeleteFollowParams{
		UserID: usr.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFollow(context.Background(),deleteParams)
	if err != nil{
		return err
	}
	fmt.Printf("Sucessfully unfollowed %s\n",feed.Name)	

	return nil

}

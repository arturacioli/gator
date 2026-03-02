package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/arturacioli/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s* State, cmd command, usr *database.User) error{
	args := cmd.arguments
	if len(args) < 1 {
		return errors.New("Not enough arguments")
	}
	url := args[0]

	rUser,err := s.db.GetUser(context.Background(),usr.Name)
	if err != nil{
		return errors.New("current user wasn't found")
	}

	rFeed,err := s.db.GetFeed(context.Background(), url)
	
	if err != nil{
		return errors.New("Feed is not registered")
	}

	follow := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: rUser.ID,
		FeedID: rFeed.ID,
	}
	
	cFollow, err := s.db.CreateFeedFollow(context.Background(), follow)
	if err != nil{
		return err
	}

	fmt.Printf("%s was followed by %s\n", cFollow.FeedName,cFollow.UserName)
	return nil
}

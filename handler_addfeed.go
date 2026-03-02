package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/arturacioli/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s* State, cmd command, usr *database.User) error{
	args := cmd.arguments
	if len(args) < 2{
		return errors.New("Not enough arguments")
	}

	user,err := s.db.GetUser(context.Background(),usr.Name)
	if err != nil{
		return errors.New("current user wasn't found")
	}

	newFeed := database.CreateFeedParams{
		ID: uuid.New(),
		Name: args[0],		
		Url: args[1],
		UserID: user.ID,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	rfeed, err := s.db.CreateFeed(context.Background(),newFeed)
	if err != nil{
		return err
	}

	follow := database.CreateFeedFollowParams{
		ID: uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID: user.ID,
		FeedID: newFeed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), follow)
	if err != nil{
		return err
	}
	fmt.Printf("%s was created and followed by %s\n",rfeed.Name, user.Name)
	return nil
}

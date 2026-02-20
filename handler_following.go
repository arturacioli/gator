package main

import (
	"context"
	"fmt"

	"github.com/arturaciolii/gator/internal/database"
)

func HandlerFollowing(s *State, cmd command, user database.User) error{
	userName := s.cfg.Username

	user, err := s.db.GetUser(context.Background(),userName)
	if err != nil{
		return err
	}

	listFollows,err := s.db.GetFeedFollowsForUser(context.Background(),user.ID)
	if err != nil{
		return err
	}

	fmt.Println("You currently follow:")

	for _,follow := range listFollows{
		fmt.Printf(" %s\n",follow.FeedName)
	}

	return nil

}

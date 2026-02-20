package main

import (
	"context"
	"fmt"
)

func HandlerFeeds(s* State, cmd command) error{

	rfeeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err 
	}
	
	fmt.Println("Current feeds:")
	for _, feed := range rfeeds{
		fmt.Println("======================")
		fmt.Printf(" Name: %s\n", feed.Name)
		fmt.Printf(" Url: %s\n", feed.Url)
		fmt.Printf(" User: %s\n", feed.Name_2)
		fmt.Printf(" Created at: %s\n", feed.CreatedAt)
		fmt.Printf(" Updated at: %s\n", feed.UpdatedAt)
		fmt.Println("======================")
	}

	return nil
}

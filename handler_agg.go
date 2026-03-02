package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/arturaciolii/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAgg(s* State, cmd command) error{
	if len(cmd.arguments) < 1{
		return fmt.Errorf("Not enough arguments")
	}

	timeBetweenRequests,err := time.ParseDuration(cmd.arguments[0])
	if err != nil{
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s* State) error{

	nextFeed,err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil{
		return err
	}

	currentTime := time.Now()
	nullTime := sql.NullTime{
		Time: currentTime,
		Valid: true,
	}
	
	markNextFeedParam := database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		LastFetchedAt: nullTime,
	}

	err = s.db.MarkFeedFetched(context.Background(),markNextFeedParam)

	fetchedFeed,err := fetchFeed(context.Background(),nextFeed.Url)

	for _,item := range fetchedFeed.Channel.Item{
		time.Parse()
		newPost := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: item.Description,
			PublishedAt: item.PubDate,

		}



	}
	
	return nil
}


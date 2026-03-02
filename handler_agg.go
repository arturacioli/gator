package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/lib/pq"
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
	if err != nil{
		return err
	}	

	fetchedFeed,err := fetchFeed(context.Background(),nextFeed.Url)
	if err != nil{
		return err 
	}


	for _,item := range fetchedFeed.Channel.Item{
		titleNull := sql.NullString{
			String: item.Title,
			Valid: true,
		}
		descNull := sql.NullString{
			String: item.Description,
			Valid: true,
		}
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil{
			fmt.Println("Error parsing published date")
			pubAt = nullTime.Time;
		}

		newPost := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: titleNull,
			Url: item.Link,
			Description: descNull,
			PublishedAt: pubAt,
			FeedID: nextFeed.ID,	
		}
		_, err = s.db.CreatePost(context.Background(),newPost)
		if err != nil{
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				continue	
			}
			return err
		}
		fmt.Printf("Created post: %s\n",item.Title)


	}
	
	return nil
}

